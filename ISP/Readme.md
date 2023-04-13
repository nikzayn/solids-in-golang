# Overview
- The Interface Segregation Principle (ISP) states that clients or users should not be forced to depend on interfaces that they don't
interfere with.
- Note: Interfaces should be generic for specific needs of the clients that uses them, rather than explicit with unnecessary methods.

## Design approach for a ISP principle
- My approach would be to create use two interfaces `Printer` and `Scanner`as we usually have to print and scan our documents for various scenarios.
- To achieve ISP, I will create one more master interface which will acts as a composite interface that combines both `Printer` and `Scanner`.
- So, by this approach will create interfaces that are cohesive, maintainable and easily extensible while avoiding unnecessary dependencies b/w
client and interfaces