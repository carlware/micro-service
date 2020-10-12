Feature: account
  In order to use account API
  As an API user
  I need to able to manage accounts

  Scenario: should get a list of accounts
    When I send "GET" request to "/"
    Then The response code should be 200
    And The reponse should match json:
    """
    {
      "accounts": []
    }
    """
