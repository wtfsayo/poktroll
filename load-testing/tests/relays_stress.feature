Feature: Loading gateway server with relays

  Scenario: Incrementing the number of relays and actors
    Given localnet is running
    And a rate of "1" relay requests per second is sent per application
    And the following initial actors are staked:
      | actor       | count |
      | application | 4     |
      | gateway     | 1     |
      | supplier    | 1     |
    And more actors are staked as follows:
      | actor       | actor inc amount | blocks per inc | max actors |
      | application | 4                | 4              | 12         |
      | gateway     | 1                | 4              | 3          |
      | supplier    | 1                | 4              | 3          |
    When a load of concurrent relay requests are sent from the applications
    # TODO_IMPROVE: Re-implement this so it's either "wait for N blocks" or "wait for N sessions"
    # to make it more deterministic.
    And the user should wait for "15" seconds
    Then the correct pairs count of claim and proof messages should be committed on-chain