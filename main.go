package main

import (
	"log"
	"strings"

	coreV1 "k8s.io/api/core/v1"

	"godogs-k8s-acceptance/pkg/k8s"
)

func main() {
	err := k8s.InitKubernetesClient()
	if err != nil {
		log.Fatal(err)
	}

	compliantPod, err := k8s.LoadPodFromYaml("./k8s/pods/compliant.yaml", "compliant-pod", "acceptance-tests")
	if err != nil {
		log.Fatal(err)
	}

	err = k8s.ApplyPodManifest(compliantPod)
	if err != nil {
		log.Fatal(err)
	}

	podNames, err := k8s.GetPodsInNamespace(compliantPod.Namespace)
	if err != nil {
		log.Fatal(err)
	}

	for _, podName := range podNames {
		if(strings.Compare(compliantPod.Name, podName) == 0) {
			log.Println("found the pod")
			break
		}
	}

	k8s.RemoveLimitFromContainer(compliantPod, coreV1.ResourceMemory, 0)
	k8s.SetContainerUser(compliantPod, 0, 0)
	k8s.SetPodNamespace(compliantPod, "kube-system")

	err = k8s.ApplyPodManifest(compliantPod)
	if err != nil {
		log.Fatal(err)
	}
}
