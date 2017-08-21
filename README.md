# Appinv - A system for managing applications and their underlying configuration items.

We're using go on the backend, and vue on the front. Using mongo for the database.

Just need to bang out the basic CRUD functions first, don't bother with spreadsheet view.

## Reference: Things Chosen
* [The NewStack: Building an App with Go](https://thenewstack.io/make-a-restful-json-api-go/)
* [How to Build a Microservice in MongoDB with Go](http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/)
* [Vue practical examples](https://tutorialzine.com/2016/03/5-practical-examples-for-learning-vue-js)
* [REST, Go, Gorilla](https://medium.com/@maumribeiro/a-fullstack-epic-part-i-a-rest-api-in-go-accessing-mongo-db-608b46e969cd)

### Database Related
* [MongoDB tutorial](https://www.youtube.com/watch?v=pWbMrx5rVBE)
* [mgo simple example](http://www.jancarloviray.com/blog/go-mongodb-simple-example/)
* [Mongo / mgo samples](https://gist.github.com/ardan-bkennedy/9198289)

### Testing Related
* [Golang unit testing tips](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742}

## Under Evaluation
* [Vue Admin Panel](https://github.com/vue-bulma/vue-admin)

## Reference: Alternatives Reviewed
* [Alternative to Gorilla: REST services made easy in go with go-restful](https://github.com/emicklei/go-restful)
* [go-restful example](http://ernestmicklei.com/2012/11/go-restful-first-working-example/)

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
