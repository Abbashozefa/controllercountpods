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
	countv1 "github.com/Abbashozefa/controllercountpods/api/v1"
	// corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// CountPodReconciler reconciles a CountPod object
type CountPodReconciler struct {
	client.Client
	// Log logr.logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=count.my.domain,resources=countpods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=count.my.domain,resources=countpods/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=count.my.domain,resources=countpods/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CountPod object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.4/pkg/reconcile
func (r *CountPodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	logger := logf.FromContext(ctx)
	logger.Info("Running")
	logger.Info("Running")
	logger.Info("Running")

	// TODO(user): your logic here
	var mypod countv1.CountPod
	if err := r.Get(ctx, req.NamespacedName, &mypod); err != nil {
		logger.Error(err, "unable to fetch list")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if len(mypod) == 0 {
		logger.V(1).Info("no resources configured")
		return ctrl.Result{}, nil
	} else {
		var podObject corev1.Pod
		err := r.Get(context.Background(), req.NamespacedName, &podObject)
		if err != nil {
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
		logger.V(1).Info("found repoter configured. sending report")
		logger.V(1).Info(podObject.Items)

	}
	logger.Info("Running")
	logger.Info("Running")
	logger.Info("Running")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CountPodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&countv1.CountPod{}).
		Named("countpod").
		Complete(r)
}
