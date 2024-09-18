Feature: informacion_pago
Test all endpoints exposed by the controller.

   Scenario Outline: /informacion_pago
       Given I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response status code is "<codres>"

 Examples:
      |method|route                        |bodyreq |codres           |
      |GET   |/v1/informacion_pago   |null    |200 OK           |
      |GET   |/v1/informacion_pago/1 |null    |200 OK           |
      |GET   |/v1/informacion_pago/33 |null    |200 OK    |