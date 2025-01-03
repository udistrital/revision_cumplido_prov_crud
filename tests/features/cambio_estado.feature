Feature: cambio_estado
Test all endpoints exposed by the controller.

   Scenario Outline: /cambio_estado_cumplido
       Given I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response status code is "<codres>"

 Examples:
      |method|route                        |bodyreq |codres           |
      |GET   |/v1/cambio_estado_cumplido   |null    |200 OK           |
      |GET   |/v1/cambio_estado_cumplido/5 |null    |200 OK           |
      |GET   |/v1/cambio_estado_cumplido/1 |null    |404 Not Found    |