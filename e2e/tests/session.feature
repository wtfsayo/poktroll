Feature: Session Namespace

  # TODO_TECHDEBT(@Olshansk, #180): This test requires you to run `make supplier1_stake && make app1_stake` first
  # As a shorter workaround, we can also add steps that stake the application and supplier as part of the scenario.
  Scenario: Supplier completes claim/proof lifecycle for a valid session
    Given the user has the pocketd binary installed
    When the supplier "supplier1" has serviced a session with "5" relays for service "svc1" for application "app1"
    And after the supplier creates a claim for the session for service "svc1" for application "app1"
    Then the claim created by supplier "supplier1" for service "svc1" for application "app1" should be persisted on-chain
#  TODO_IMPROVE: ...
#    And an event should be emitted...
    And after the supplier submits a proof for the session for service "svc1" for application "app1"
    Then the proof submitted by supplier "supplier1" for service "svc1" for application "app1" should be persisted on-chain
#  TODO_IMPROVE: ...
#    And an event should be emitted...
