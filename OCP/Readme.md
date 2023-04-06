# Overview
- The open/closed principle is a software design principle that states should be open for extension but closed for modification.
We can use interfaces for this scenario.

## Best practices while implementing SRP in Go
- Use interfaces to abstract functionality.
- Readable and reusable packages that have a clear intent.
- Avoid methodologies to different functions and pacakges.
- Small and focused functions and methods should do one thing and do it properly.


## Design approach for a OCP problem
- Let's say, we want to create greet functionality where we can greet in different languages.
- I will create one interface where it will have all the greet methods behaviour.

