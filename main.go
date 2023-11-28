package main

import (
	"log"

	coreV1 "k8s.io/api/core/v1"

	"godogs-k8s-acceptance/pkg/k8s"
)

func main() {
	err := k8s.InitKubernetesClient()
	if err != nil {
		log.Fatal(err)
	}

	compliantPod, err := k8s.LoadPodFromYaml("./k8s/pods/compliant.yaml")
	if err != nil {
		log.Fatal(err)
	}

	k8s.RemoveLimitFromContainer(compliantPod, coreV1.ResourceMemory, 0)
	k8s.SetContainerUserRoot(compliantPod, 0)
	k8s.SetPodNamespace(compliantPod, "kube-system")

	err = k8s.ApplyPodManifest(compliantPod)
	if err != nil {
		log.Fatal(err)
	}
}
