Feature: Tokenomics Namespaces

    # This test
    Scenario: Basic tokenomics validation that Supplier mint equals Application burn
        Given the user has the pocketd binary installed
        And an account exists for "supplier1"
        And an account exists for "app1"
        When the supplier "supplier1" has serviced a session with "20" relays for service "svc1" for application "app1"
        And the user should wait for "10" seconds
        # TODO_IN_THIS_PR: Uncomment these two lines
        Then the account balance of "supplier1" should be "1000" uPOKT "more" than before
        And the account balance of "app1" should be "1000" uPOKT "less" than before
