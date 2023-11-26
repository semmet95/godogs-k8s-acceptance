package main

import (
	"context"
	"fmt"
	"log"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"godogs-k8s-acceptance/pkg/k8s"
)

func main() {
	k8sClient, err := k8s.GetKubernetesClient()
	if err != nil {
		log.Fatal(err)
	}

	compliantPod, err := k8s.LoadPodFromYaml("./k8s/pods/compliant.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(compliantPod.Spec.Containers[0].Resources.Limits)
	compliantPod.Spec.Containers[0].Resources.Limits[coreV1.ResourceMemory] = resource.Quantity{}
	fmt.Println(compliantPod.Spec.Containers[0].Resources.Limits)

	applyOpResult, err := k8sClient.CoreV1().Pods(compliantPod.Namespace).Create(context.Background(), compliantPod, metaV1.CreateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("apply result is: %v", applyOpResult)
}
