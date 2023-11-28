package k8s

import (
	"fmt"
	"os"

	coreV1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

func LoadPodFromYaml(yamlPath string) (*coreV1.Pod, error) {
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, fmt.Errorf("error reading yaml file: %v", err)
	}

	var pod coreV1.Pod
	err = yaml.Unmarshal(yamlFile, &pod)
	if err != nil {
		return nil, fmt.Errorf("error parsing yaml file: %v", err)
	}

	return &pod, nil
}

func ApplyPodManifest(pod *coreV1.Pod) (error) {
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
	pod.Spec.Containers[containerIndex].SecurityContext = &coreV1.SecurityContext{ RunAsUser: &rootUser }
}