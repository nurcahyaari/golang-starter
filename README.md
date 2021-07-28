# golang-starter
Golang base project using fiber framework
 
### Folder Structure

here how golang-started manage directory:
```
|-- internal
|   |-- app
|   |-- config
|   |-- db
|   |   |-- localdb
|   |   |-- onlinedb
|   |-- middleware
|   |-- routes
|   |-- utils
|       |-- auth
|       |-- encryption
|       |-- response
|-- src
```

all global setting will place in internal folder
all module will place in src


# Internal
Internal with their folders are using for the whole apps that will be needed to be a global function that will call in other domain. example:

we have a products domain with their controller, repo, service, routes, in the service we need a jwt for verification the jwt token. the jwt token function will be a callable function on other domains, hence the jwt must be a global function

# Src
This boilerplate based on domain. eg: products domain, user domain, etc. and the domain have their structure based on clean architecture. controllers, models, repositories, services, routers. the flow is router -> controller -> service -> repository.

## router
All App Router eg: HTTP Router, RabbitMQ Queue, will defined here

## Controller
Controller usefull to encode raw body, query params, make a response to client. note: make the controller based on the protocols eg: HTTP Controller, RabbiqMQ controller, Grpc controller.

## Service
Service contains all of business logic, like creating a products, creating a users, validation products, validation for products quantity. 
note that if you have relations based on the domain example you have a product domain with order domain, when user make an order then product quantity must be reduced based on order quantity, so you must make a transaction in service

## Repository
Repository contains all of database query