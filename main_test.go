package main

import (
	"context"
	"errors"
	"log"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	coreV1 "k8s.io/api/core/v1"

	"godogs-k8s-acceptance/pkg/k8s"
)

type podName struct{}
type podNamespace struct{}
type pod struct{}

func createPodCompliantWithAllPolicies(ctx context.Context, k8sPodName, k8sPodNamespace string) (context.Context, error) {
	k8sPod, err := k8s.LoadPodFromYaml("./k8s/pods/compliant.yaml", k8sPodName, k8sPodNamespace)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, pod{}, k8sPod), nil
}

func applyPodManifest(ctx context.Context) (context.Context, error) {
	k8sPod, ok := ctx.Value(pod{}).(*coreV1.Pod)
	if !ok {
		return ctx, errors.New("there is no pod set to apply")
	}

	err := k8s.ApplyPodManifest(k8sPod)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(
		context.WithValue(ctx, podName{}, k8sPod.GetName()),
		podNamespace{},
		k8sPod.GetNamespace(),
	), nil
}

func podShouldBeInNamespace(ctx context.Context) (context.Context, error) {
	k8sPodName, ok := ctx.Value(podName{}).(string)
	if !ok {
		return ctx, errors.New("pod name is not set")
	}

	k8sPodNamespace, ok := ctx.Value(podNamespace{}).(string)
	if !ok {
		return ctx, errors.New("pod namespace is not set")
	}

	namespacePodNames, err := k8s.GetPodsInNamespace(k8sPodNamespace)
	if err != nil {
		return ctx, err
	}

	for _, namespacePodName := range namespacePodNames {
		if strings.Compare(k8sPodName, namespacePodName) == 0 {
			return ctx, nil
		}
	}

	return ctx, errors.New("pod not found in the namespace")
}

func setUserForContainer(ctx context.Context, user, containerIdx int64) (context.Context, error) {
	k8sPod, ok := ctx.Value(pod{}).(*coreV1.Pod)
	if !ok {
		return ctx, errors.New("there is no pod set to apply")
	}

	k8s.SetContainerUser(k8sPod, 0, 0)

	return ctx, nil
}

func TestJsPolicies(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	err := k8s.InitKubernetesClient()
	if err != nil {
		log.Fatal(err)
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	

	sc.Given(`^I create a pod manifest with name ([a-z0-9][-a-z0-9]*[a-z0-9]?) in namespace ([a-z0-9][-a-z0-9]*[a-z0-9]?) that is compliant with all policies enforced$`, createPodCompliantWithAllPolicies)
	sc.Step(`^I set the user of container indexed (\d+) as (\d+) i.e., root$`, setUserForContainer)
	sc.When(`^I apply the pod manifest$`, applyPodManifest)
	sc.Then(`^the pod should be created in the namespace$`, podShouldBeInNamespace)
}
