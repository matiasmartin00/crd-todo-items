/*
Copyright 2025.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	matiasmartin00v1 "github.com/matiasmartin00/crd-todo-items/api/v1"
)

// TodoItemReconciler reconciles a TodoItem object
type TodoItemReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=matiasmartin00.matiasmartin00.com,resources=todoitems,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=matiasmartin00.matiasmartin00.com,resources=todoitems/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=matiasmartin00.matiasmartin00.com,resources=todoitems/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TodoItem object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *TodoItemReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx)

	logf.Log.Info("Reconcile called for TodoItem", "NamespacedName", req.NamespacedName)
	logf.Log.Info("Name of the TodoItem", "Name", req.Name)
	logf.Log.Info("Namespace of the TodoItem", "Namespace", req.Namespace)

	var todo matiasmartin00v1.TodoItem
	if err := r.Get(ctx, req.NamespacedName, &todo); err != nil {
		// The TodoItem resource may have been deleted after the reconcile request.
		// Return and don't requeue
		logf.Log.Error(err, "unable to fetch TodoItem")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	todo.Spec.Completed = false
	logf.Log.Info("Reconciled TodoItem", "Spec", todo.Spec)
	logf.Log.Info("Reconciled TodoItem", "Status", todo.Status)
	logf.Log.Info("Generation", "Generation", todo.Generation)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TodoItemReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&matiasmartin00v1.TodoItem{}).
		Named("todoitem").
		Complete(r)
}
