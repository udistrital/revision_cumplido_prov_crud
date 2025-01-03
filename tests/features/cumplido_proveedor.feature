Feature: cumplido_proveedor
Test all endpoints exposed by the controller.

   Scenario Outline: /cumplido_proveedor
       Given I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response status code is "<codres>"

 Examples:
      |method|route                        |bodyreq |codres           |
      |GET   |/v1/cumplido_proveedor   |null    |200 OK           |
      |GET   |/v1/cumplido_proveedor/5 |null    |200 OK           |
      |GET   |/v1/cumplido_proveedor/99|null    |404 Not Found    |