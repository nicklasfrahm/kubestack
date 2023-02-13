/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/nicklasfrahm/kubestack/api/v1alpha1"
	"github.com/nicklasfrahm/kubestack/pkg/management"
	"github.com/nicklasfrahm/kubestack/pkg/management/common"
)

// InterfaceReconciler reconciles a Interface object
type InterfaceReconciler struct {
	client.Client
	recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

//+kubebuilder:rbac:groups=kubestack.nicklasfrahm.dev,resources=interfaces,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kubestack.nicklasfrahm.dev,resources=interfaces/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kubestack.nicklasfrahm.dev,resources=interfaces/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Interface object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
func (r *InterfaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	iface := new(v1alpha1.Interface)
	err := r.Client.Get(ctx, req.NamespacedName, iface)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	connRef := types.NamespacedName{
		Namespace: iface.Spec.ConnectionRef.Namespace,
		Name:      iface.Spec.ConnectionRef.Name,
	}
	if connRef.Namespace == "" {
		connRef.Namespace = req.Namespace
	}

	mgmt, err := management.NewClient(connRef, common.WithKubernetesClient(r.Client))
	if err != nil {
		r.recorder.Event(iface, corev1.EventTypeWarning, "ConnectionFailed", err.Error())
		return ctrl.Result{}, nil
	}
	defer mgmt.Disconnect()

	ifaceService, err := mgmt.Interface()
	if err != nil {
		r.recorder.Event(iface, corev1.EventTypeWarning, "MissingDriver", err.Error())
		return ctrl.Result{}, nil
	}

	if iface, err = ifaceService.UpdateInterface(iface); err != nil {
		r.recorder.Event(iface, corev1.EventTypeWarning, "UpdateFailed", err.Error())
		return ctrl.Result{}, nil
	}

	// Enable two-way sync.
	if err := r.Update(ctx, iface); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// TODO: Set up scaffolding for finalizers.

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *InterfaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("interface-controller")

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Interface{}).
		Complete(r)
}
