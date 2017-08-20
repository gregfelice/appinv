# Appinv - A system for managing applications and their underlying configuration items.

We're using go on the backend, and vue on the front.
Using mongo for the database.

Reference
---
* [MongoDB tutorial](https://www.youtube.com/watch?v=pWbMrx5rVBE)
* [The NewStack: Building an App with Go](https://thenewstack.io/make-a-restful-json-api-go/)
* [How to Build a Microservice in MongoDB with Go](http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/)
* [mgo samples](https://gist.github.com/ardan-bkennedy/9198289)

Stories
---
### Load applications from excel sheets
ok - what we're dealing with here - getting shit loaded into mongodb from excel.
ok - the basics are working.

### GET applications via REST service
ok - want to pull the data from the database, test with a curl call.
ok - we can do it now. returns a list from the database.

POST applications
DELETE applications
PUT applications

### View applications
ok - we want to do this with vue. get from REST call, make pretty

Search applications

GET ci
POST ci
DELETE ci
PUT ci
