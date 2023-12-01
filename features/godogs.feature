Feature: jsPolicies
  In order to deploy a pod
  As a developer
  I need to configure the pod to be compliant with all the jsPolicies

  Rules:
  - Pods without memory limit set are not allowed
  - Pods with containers that run as the root user are not allowed
  - Pods cannot be deployed in kube-system and default namespace

  Scenario: Allow deployment of a compliant pod
    Given I have a pod compliant with all the policies enforced
    When I apply its manifest
    Then it should create the pod in the corresponding namespace