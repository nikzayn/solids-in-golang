# Single Responsibility Principle
- To exlain this concept, I will implement one real world example. 


## Overview
- The single responsibility principle is a software design principle that states that a module and a class
should have single or have one reason to change.


## Types of ways in Go
---

### Package-level SRP:
- It means that each package should have a single responsibility, and same the code should be focusing that responsibility.

### Function-level SRP:
- It means that each package should have a single responsibility, and same the code should be focusing that responsibility.

### Type-level SRP:
- It means that each type like struct or interfaces should have a single responsibility, and same the code should be focusing that responsibility.


## Best practices while implementing SRP in Go
- Use interfaces to abstract functionality.
- Readable and reusable packages that have a clear intent.
- Avoid methodologies to different functions and pacakges.
- Small and focused functions and methods should do one thing and do it properly.


## Design approach for a SRP problem
- Let's say, as I do shopping related to groceries, household products etc. I will design a shopping cart system where I will create different packages to set different responsibilites like Cart, Pricing, Item etc.
- While implementing each package let's first start with Cart, we will implement struct and methods like AddItem, RemoveItem, UpdateItem, etc.
- 

