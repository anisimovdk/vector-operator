package k8sutils

import (
	"bytes"
	"context"
	"io"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func CreateOrUpdateService(svc *corev1.Service, c client.Client) (*reconcile.Result, error) {
	return reconcileService(svc, c)
}

func CreateOrUpdateSecret(secret *corev1.Secret, c client.Client) (*reconcile.Result, error) {
	return reconcileSecret(secret, c)
}

func CreateOrUpdateDaemonSet(daemonSet *appsv1.DaemonSet, c client.Client) (*reconcile.Result, error) {
	return reconcileDaemonSet(daemonSet, c)
}

func CreateOrUpdateStatefulSet(statefulSet *appsv1.StatefulSet, c client.Client) (*reconcile.Result, error) {
	return reconcileStatefulSet(statefulSet, c)
}

func CreateOrUpdateServiceAccount(secret *corev1.ServiceAccount, c client.Client) (*reconcile.Result, error) {
	return reconcileServiceAccount(secret, c)
}

func CreateOrUpdateClusterRole(secret *rbacv1.ClusterRole, c client.Client) (*reconcile.Result, error) {
	return reconcileClusterRole(secret, c)
}

func CreateOrUpdateClusterRoleBinding(secret *rbacv1.ClusterRoleBinding, c client.Client) (*reconcile.Result, error) {
	return reconcileClusterRoleBinding(secret, c)
}

func CreatePod(pod *corev1.Pod, c client.Client) error {
	err := c.Create(context.TODO(), pod)
	if err != nil {
		return err
	}
	return nil
}

func GetPod(pod *corev1.Pod, c client.Client) (*corev1.Pod, error) {
	result := &corev1.Pod{}
	err := c.Get(context.TODO(), client.ObjectKeyFromObject(pod), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetPodLogs(pod *corev1.Pod, cs *kubernetes.Clientset) (string, error) {
	count := int64(100)
	podLogOptions := corev1.PodLogOptions{
		TailLines: &count,
	}

	req := cs.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOptions)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", err
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", err
	}
	str := buf.String()

	return str, nil
}

func UpdateStatus(ctx context.Context, obj client.Object, c client.Client) error {
	return c.Status().Update(ctx, obj)
}

func reconcileService(obj runtime.Object, c client.Client) (*reconcile.Result, error) {

	existing := &corev1.Service{}
	desired := obj.(*corev1.Service)

	err := c.Create(context.TODO(), desired)
	if err != nil && errors.IsAlreadyExists(err) {
		err := c.Get(context.TODO(), client.ObjectKeyFromObject(desired), existing)
		if err != nil {
			return nil, err
		}
		if !equality.Semantic.DeepEqual(existing, desired) {
			existing.Spec = desired.Spec
			existing.Labels = desired.Labels
			err := c.Update(context.TODO(), existing)
			return nil, err
		}
	}
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	return nil, nil
}

func reconcileSecret(obj runtime.Object, c client.Client) (*reconcile.Result, error) {

	existing := &corev1.Secret{}
	desired := obj.(*corev1.Secret)

	err := c.Create(context.TODO(), desired)
	if err != nil && errors.IsAlreadyExists(err) {
		err := c.Get(context.TODO(), client.ObjectKeyFromObject(desired), existing)
		if err != nil {
			return nil, err
		}
		if !equality.Semantic.DeepEqual(existing, desired) {
			existing.Data = desired.Data
			existing.Labels = desired.Labels
			err := c.Update(context.TODO(), existing)
			return nil, err
		}
	}
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	return nil, nil
}

func reconcileDaemonSet(obj runtime.Object, c client.Client) (*reconcile.Result, error) {

	existing := &appsv1.DaemonSet{}
	desired := obj.(*appsv1.DaemonSet)

	err := c.Create(context.TODO(), desired)
	if err != nil && errors.IsAlreadyExists(err) {
		err := c.Get(context.TODO(), client.ObjectKeyFromObject(desired), existing)
		if err != nil {
			return nil, err
		}
		if !equality.Semantic.DeepEqual(existing, desired) {
			existing.Spec = desired.Spec
			existing.Labels = desired.Labels
			err := c.Update(context.TODO(), existing)
			return nil, err
		}
	}
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	return nil, nil
}

func reconcileStatefulSet(obj runtime.Object, c client.Client) (*reconcile.Result, error) {

	existing := &appsv1.StatefulSet{}
	desired := obj.(*appsv1.StatefulSet)

	err := c.Create(context.TODO(), desired)
	if err != nil && errors.IsAlreadyExists(err) {
		err := c.Get(context.TODO(), client.ObjectKeyFromObject(desired), existing)
		if err != nil {
			return nil, err
		}
		if !equality.Semantic.DeepEqual(existing, desired) {
			existing.Spec = desired.Spec
			existing.Labels = desired.Labels
			err := c.Update(context.TODO(), existing)
			return nil, err
		}
	}
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	return nil, nil
}

func reconcileServiceAccount(obj runtime.Object, c client.Client) (*reconcile.Result, error) {

	existing := &corev1.ServiceAccount{}
	desired := obj.(*corev1.ServiceAccount)

	err := c.Create(context.TODO(), desired)
	if err != nil && errors.IsAlreadyExists(err) {
		err := c.Get(context.TODO(), client.ObjectKeyFromObject(desired), existing)
		if err != nil {
			return nil, err
		}
	}
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	return nil, nil
}

func reconcileClusterRole(obj runtime.Object, c client.Client) (*reconcile.Result, error) {

	existing := &rbacv1.ClusterRole{}
	desired := obj.(*rbacv1.ClusterRole)

	err := c.Create(context.TODO(), desired)
	if err != nil && errors.IsAlreadyExists(err) {
		err := c.Get(context.TODO(), client.ObjectKeyFromObject(desired), existing)
		if err != nil {
			return nil, err
		}
		if !equality.Semantic.DeepEqual(existing, desired) {
			existing.Rules = desired.Rules
			err := c.Update(context.TODO(), existing)
			return nil, err
		}
	}
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	return nil, nil
}

func reconcileClusterRoleBinding(obj runtime.Object, c client.Client) (*reconcile.Result, error) {

	existing := &rbacv1.ClusterRoleBinding{}
	desired := obj.(*rbacv1.ClusterRoleBinding)

	err := c.Create(context.TODO(), desired)
	if err != nil && errors.IsAlreadyExists(err) {
		err := c.Get(context.TODO(), client.ObjectKeyFromObject(desired), existing)
		if err != nil {
			return nil, err
		}
		if !equality.Semantic.DeepEqual(existing, desired) {
			existing.RoleRef = desired.RoleRef
			existing.Subjects = desired.Subjects
			err := c.Update(context.TODO(), existing)
			return nil, err
		}
	}
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	return nil, nil
}
