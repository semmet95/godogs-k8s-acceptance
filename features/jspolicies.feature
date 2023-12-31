Feature: jsPolicies
  In order to deploy a pod
  As a developer
  I need to configure the pod to be compliant with all the jsPolicies

  Rules:
  - Pods without memory limit set are not allowed
  - Pods with containers that run as the root user are not allowed
  - Pods cannot be deployed in kube-system and default namespace

  Scenario: Allow deployment of a compliant pod
    Given I create a pod manifest with name compliant-pod in namespace acceptance-tests that is compliant with all policies enforced
    When I apply the pod manifest
    Then the pod should be created in the namespace

  Scenario: Block deployment of a pod with a container running as root
    Given I create a pod manifest with name bad-pod-1 in namespace acceptance-tests that is compliant with all policies enforced
    And I set the user of container indexed 0 as 0 i.e., root
    When I apply the pod manifest
    Then the pod should be blocked with error:
      """
      - Field spec.containers[0].securityContext is not allowed.
      """

  Scenario: Block deployment of a pod with a container without memory limit set
    Given I create a pod manifest with name bad-pod-2 in namespace acceptance-tests that is compliant with all policies enforced
    And I remove the memory limit of container indexed 0
    When I apply the pod manifest
    Then the pod should be blocked with error:
      """
      - Memory limit not defined for spec.containers[0]
      """

  Scenario: Block deployment of a pod in the namespace kube-system
    Given I create a pod manifest with name bad-pod-2 in namespace acceptance-tests that is compliant with all policies enforced
    And I set the pod namespace as kube-system
    When I apply the pod manifest
    Then the pod should be blocked with error:
      """
      - Field metadata.namespace is not allowed to be: default | kube-system
      """

  Scenario: Block deployment of a pod with a container with user set to root and with memory limit removed in the namespace kube-system
    Given I create a pod manifest with name bad-pod-2 in namespace acceptance-tests that is compliant with all policies enforced
    And I set the user of container indexed 0 as 0 i.e., root
    And I remove the memory limit of container indexed 0
    And I set the pod namespace as kube-system
    When I apply the pod manifest
    Then the pod should be blocked with error:
      """
      - Field metadata.namespace is not allowed to be: default | kube-system
      - Field spec.containers[0].securityContext is not allowed.
      - Memory limit not defined for spec.containers[0]
      """