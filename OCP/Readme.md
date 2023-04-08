# Overview
- The open/closed principle is a software design principle that states should be open for extension but closed for modification.
We can use interfaces for this scenario.

## Examples
- In first example, I will be explaining how we can use same functionality for multiple structs
- In second example, I will be explaining how to achieve struct embedding via composition using interfaces to fulfill the ocp concept.


## Design approach for a OCP problem
- Let's say, we want to create greet functionality where we can greet in different languages.
- I will create one interface where it will have all the greet methods behaviour.

