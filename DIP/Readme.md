# Overview
- The Dependency Injection principle states that high-level modules should not depend on low-level modules, but both should depend on abstraction.
- Dependencies should be abstracted via intefaces to achieve flexibility and extensibility in codebase

## Design approach to implement DIP principle
- So. I will implement one use case when we have to implement multiple datasource with API as a medium. Some datasources like some popular databases like MySQL, PostgreSQL, MongoDB etc. or flat files and some webs services too.
- We will data source interface to get the data from datasources using databases
- Think like this there is an APi, which depends on some data source rather than depending on database which is a concrete implementaion, API depends on the abstraction of data source to achieve dependency, it doesn't necessary to depend on database.
- So basically, it work like this, I will create one database say `db` in main function, which use concrete implementaion of `Database`, later when we want to pass the information of db to `API`, we don't need to send it directly to API as a value, we can rather abstract the `Datasource` which is a mediator b/w `API` and `Database`. Well, that will fulfill the implementation using abstraction. Hence, we don't need to change any pre-written logic for this in case of `API`.
- Hence, this implemenation will DIP and makes our codebase more flexible.