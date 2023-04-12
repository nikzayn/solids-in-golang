# Overview
- The Liskov Substitution Principle states that objects of a superclass should be able to be replaced with objects
of a subclass without affecting the correctness of the program.

## Design approach for LSP principle
- I will take simple examples of geomtry shapes where I will be using two shapes as an example which would be `Rectangle` and `Circle`
- Hence, I will also create method to calculate `area` for both shapes mentioned above.
- Whereas, we would be using `interfaces` to add as many shapes as possible without changing the correctness of the program.
- Hence, we will be achieving to produce generic code with `Shape` as an interface.