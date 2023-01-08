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
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kubestackv1alpha1 "github.com/nicklasfrahm/kubestack/api/v1alpha1"
	"github.com/nicklasfrahm/kubestack/pkg/util"
)

// ConnectionReconciler reconciles a Connection object
type ConnectionReconciler struct {
	client.Client
	recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

//+kubebuilder:rbac:groups=kubestack.nicklasfrahm.dev,resources=connections,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kubestack.nicklasfrahm.dev,resources=connections/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kubestack.nicklasfrahm.dev,resources=connections/finalizers,verbs=update
//+kubebuilder:rbac:groups=kubestack.nicklasfrahm.dev,resources=events,verbs=create;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Connection object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
func (r *ConnectionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	conn := new(kubestackv1alpha1.Connection)
	if err := r.Client.Get(ctx, req.NamespacedName, conn); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	secret := new(corev1.Secret)
	if err := r.Client.Get(ctx, client.ObjectKey{
		Namespace: conn.Spec.SecretRef.Namespace,
		Name:      conn.Spec.SecretRef.Name,
	}, secret); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Update the connection status with the OS information.
	osInfo, err := util.ProbeOS(conn, secret)
	if err != nil {
		if strings.HasPrefix(err.Error(), "ssh:") {
			r.recorder.Event(conn, corev1.EventTypeWarning, "ConnectionFailed", err.Error())
		}
		return ctrl.Result{
			RequeueAfter: 15 * time.Second,
		}, err
	}
	conn.Status.OS = *osInfo

	if err := r.Status().Update(ctx, conn); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	r.recorder.Event(conn, corev1.EventTypeNormal, "OSProbed", "OS information probed successfully.")

	// TODO: Implement other logic.

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ConnectionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("connection-controller")

	return ctrl.NewControllerManagedBy(mgr).
		For(&kubestackv1alpha1.Connection{}).
		Complete(r)
}
