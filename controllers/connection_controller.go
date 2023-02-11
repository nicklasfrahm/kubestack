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
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/nicklasfrahm/kubestack/api/v1alpha1"
	"github.com/nicklasfrahm/kubestack/pkg/management"
	"github.com/nicklasfrahm/kubestack/pkg/management/common"
)

const (
	secretField = ".spec.secretRef.name"
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
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch

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
	logger := log.FromContext(ctx)

	conn := new(v1alpha1.Connection)
	err := r.Get(ctx, req.NamespacedName, conn)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	mgmt, err := management.NewClient(req.NamespacedName, common.WithKubernetesClient(r.Client))
	if err != nil {
		r.recorder.Event(conn, corev1.EventTypeWarning, "ConnectionFailed", err.Error())
		logger.Error(err, "failed to create management client")
		return ctrl.Result{}, nil
	}
	defer mgmt.Disconnect()

	// TODO: Although this is idempotent, we may put excessive load on the API server,
	// because we trigger a reconciliation for the secret change and the connection.
	conn.Status.OS = *mgmt.OS()
	if err := r.Status().Update(ctx, conn); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	r.recorder.Event(conn, corev1.EventTypeNormal, "OSProbed", "OS information probed successfully.")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ConnectionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("connection-controller")

	// We need to add an index for the secret name so that we can
	// trigger a reconciliation if a referenced secret changes.
	mgr.GetFieldIndexer().IndexField(context.TODO(), &v1alpha1.Connection{}, secretField, func(rawObj client.Object) []string {
		conn := rawObj.(*v1alpha1.Connection)
		return []string{conn.Spec.SecretRef.Name}
	})

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Connection{}).
		// Watch for changes to referenced secrets.
		Watches(&source.Kind{Type: &corev1.Secret{}}, handler.EnqueueRequestsFromMapFunc(r.findObjectsForSecret)).
		Complete(r)
}

// findObjectsForSecret allows us to trigger a reconciliation if a referenced secret changes.
func (r *ConnectionReconciler) findObjectsForSecret(secret client.Object) []reconcile.Request {
	connections := &v1alpha1.ConnectionList{}
	listOpt := &client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector(secretField, secret.GetName()),
	}
	if err := r.List(context.TODO(), connections, listOpt); err != nil {
		return []reconcile.Request{}
	}

	requests := make([]reconcile.Request, len(connections.Items))
	for i, item := range connections.Items {
		requests[i] = reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      item.GetName(),
				Namespace: item.GetNamespace(),
			},
		}
	}
	return requests
}
