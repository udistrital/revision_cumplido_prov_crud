Feature: soporte_cumplido
Test all endpoints exposed by the controller.

   Scenario Outline: /soporte_cumplido
       Given I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response status code is "<codres>"

 Examples:
      |method|route                        |bodyreq |codres           |
      |GET   |/v1/soporte_cumplido   |null    |200 OK           |
      |GET   |/v1/soporte_cumplido/1 |null    |200 OK           |
      |GET   |/v1/soporte_cumplido/5 |null    |404 Not Found    |