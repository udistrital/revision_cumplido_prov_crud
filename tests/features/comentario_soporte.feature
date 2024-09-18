Feature: comentario_soporte
Test all endpoints exposed by the controller.

   Scenario Outline: /comentario_soporte
       Given I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response status code is "<codres>"

 Examples:
      |method|route                        |bodyreq |codres           |
      |GET   |/v1/comentario_soporte   |null    |200 OK           |
      |GET   |/v1/comentario_soporte/18 |null    |200 OK           |
      |GET   |/v1/comentario_soporte/1 |null    |404 Not Found    |