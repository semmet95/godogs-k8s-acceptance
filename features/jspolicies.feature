Feature: jsPolicies
  In order to deploy a pod
  As a developer
  I need to configure the pod to be compliant with all the jsPolicies

  Rules:
  - Pods without memory limit set are not allowed
  - Pods with containers that run as the root user are not allowed
  - Pods cannot be deployed in kube-system and default namespace

  Scenario: Allow deployment of a compliant pod
    Given I create a pod manifest with name compliant-pod-1 in namespace acceptance-tests that is compliant with all policies enforced
    When I apply the pod manifest
    Then the pod should be created in the namespace