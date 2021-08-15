# golang-starter
Golang base project using fiber framework, and implement the Clean code Architecture
<br><br>
<b>Alert</b>
<br>
 - I migrated the structures, like adding infrastructures, move the config folder from internal/config into config/
 - I also don't test after the rewriting, when you clone it and try to implement to your projects, it might need some changes
 - the old version placed on dep-0.0.1 branch
<br><br><br> 
### Folder Structure

here how golang-started manage directory:
```
|-- config
|-- infrastructures
|-- internal
|   |-- middleware
|   |-- routes
|   |-- utils
|       |-- auth
|       |-- encryption
|   |-- web
|-- src
```

all global setting will place in internal folder
all module will place in src

# Config
Config modules for the app

# Infrastructures
infrastructures store the infrastructure layer, such as database, database stored in infrastructures/db folder. If you use mysql, then just create a file called mysql.go and you must define your mysql modules here (it's a dependency).
If you using redis in your project, create a cached folder then create a file called redis.go, and make your modules again. all examples was attached in this repo

# Internal
Internal with their folders are using for the whole apps that will be needed to be a global function that will call in other modules. example:

we have a products modules with their controller, repo, service, routes, in the service we need a jwt for verification the jwt token. the jwt token function will be a callable function on other moduless, hence the jwt must be a global function

# src
This boilerplate based on modules. eg: products modules, user modules, etc. and the modules have their structure based on clean architecture. controllers, models, repositories, services, routers. the flow is router -> controller -> service -> repository.

## entities
entities is a struct to store any objects from database, example you have a products in your tables. then you must create your products.go in the entities folder, it's use to decode the data from database query and store it as a object

## dto
it's similar with entities but the differents are:
<ul>
    <li>don't define dto in repository, the repository is the responsibility of the entities</li>
    <li>dto is usefull to get your request body, request params, or response your data to your rest api</li>
</ul>
example you have products.go as entities and its entities called as Products, and products_response.go as the dto and its dto called as ProductsResponseBody. the repositories will return the data as Products and the services will return it too, but if you want to sent your data as a json response, you should have used the ProductsResponseBody to response it.

## router
All App Router eg: Http Router, RabbitMQ Queue, will defined here

## Handlers
Handlers usefull to encode raw body, query params, make a response to client. note: make the controller based on the protocols eg: HTTP Handlers, RabbiqMQ controller, Grpc controller.
and it also useful to convert entities to dto for responding to the data

## Service
Service contains all of business logic, like creating a products, creating a users, validation products, validation for products quantity, run you transaction. 
note that if you have relations based on the modules example you have a product modules with order modules, when user make an order then product quantity must be reduced based on order quantity, so you must make a transaction in service

## Repository
Repository contains all of database query