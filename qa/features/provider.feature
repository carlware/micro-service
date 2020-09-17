Feature: provider
  In order to use provider API
  As an API user
  I need to able to manage providers

  Scenario: should get a list of providers
    When I send "GET" request to "/"
    Then The response code should be 200
    And The reponse should match json:
    """
    {
      "providers": []
    }
    """
