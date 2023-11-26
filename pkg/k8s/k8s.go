package k8s

import (
	"fmt"
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubernetesClient() (*kubernetes.Clientset, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting user home dir: %v", err)
	}

	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error getting Kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("error getting Kubernetes clientset: %v", err)
	}

	return clientset, nil
}

func LoadPodFromYaml(yamlPath string) (*v1.Pod, error) {
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, fmt.Errorf("error reading yaml file: %v", err)
	}

	var pod v1.Pod
	err = yaml.Unmarshal(yamlFile, &pod)
	if err != nil {
		return nil, fmt.Errorf("error parsing yaml file: %v", err)
	}

	return &pod, nil
}
