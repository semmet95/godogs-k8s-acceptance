Feature: jsPolicies

  Scenario: Allow deployment of a compliant pod
    Given I have a pod compliant with all the policies enforced
    When I apply its manifest
    Then it should create the pod in the corresponding namespace