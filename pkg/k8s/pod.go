package k8s

import (
	"context"
	"fmt"
	"os"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func LoadPodFromYaml(yamlPath, podName, podNamespace string) (*coreV1.Pod, error) {
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, fmt.Errorf("error reading yaml file: %v", err)
	}

	var pod coreV1.Pod
	err = yaml.Unmarshal(yamlFile, &pod)
	if err != nil {
		return nil, fmt.Errorf("error parsing yaml file: %v", err)
	}

	pod.SetName(podName)
	pod.SetNamespace(podNamespace)

	return &pod, nil
}

func ApplyPodManifest(pod *coreV1.Pod) (error) {
	_, err := k8sClient.CoreV1().Pods(pod.GetNamespace()).Create(context.Background(), pod, metaV1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func RemoveLimitFromContainer(pod *coreV1.Pod, resourceType coreV1.ResourceName, containerIndex int) {
	delete(pod.Spec.Containers[containerIndex].Resources.Limits, resourceType)
}

func SetPodNamespace(pod *coreV1.Pod, namespace string) {
	pod.SetNamespace(namespace)
}

func SetContainerUserRoot(pod *coreV1.Pod, containerIndex int) {
	rootUser := int64(0)
	pod.Spec.Containers[containerIndex].SecurityContext = &coreV1.SecurityContext{RunAsUser: &rootUser}
}

func GetPodsInNamespace(namespace string) ([]string, error) {
	pods, err := k8sClient.CoreV1().Pods(namespace).List(context.Background(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var podNames []string
	for _, pod := range pods.Items {
		podNames = append(podNames, pod.Name)
	}

	return podNames, nil
}
