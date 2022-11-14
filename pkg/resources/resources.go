package resources

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	installerv1alpha1 "github.com/mrsimonemms/kubebuilder/api/v1alpha1"
)

func getLabels(clientResource *installerv1alpha1.Config) map[string]string {
	return map[string]string{
		"installer": clientResource.Spec.InstallerImage,
	}
}

func CreatePod(clientResource *installerv1alpha1.Config) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clientResource.Spec.InstallerImage,
			Namespace: clientResource.Namespace,
			Labels:    getLabels(clientResource),
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "gitpod-installer", // @todo(sje): do we need some additional things in here?
					Image: clientResource.Spec.InstallerImage,
					Ports: []corev1.ContainerPort{
						{
							Name:          "http",
							ContainerPort: 8080,
							Protocol:      corev1.ProtocolTCP,
						},
					},
				},
			},
			RestartPolicy: corev1.RestartPolicyOnFailure,
		},
	}
}
