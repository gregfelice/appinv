package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var currentId int

var todos Todos
var applications Applications

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})

	// load data from the database into applications
	loadApplications()
	// fmt.Printf("loaded applications: %s", applications)
}

func loadApplications() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("test").C("applications")

	//log.Printf("RunQuery : %d : Executing\n", query)

	// Retrieve the list of applications.
	//var applications []Application
	e := collection.Find(nil).All(&applications)
	if e != nil {
		log.Printf("RunQuery : ERROR : %s\n", e)
		return
	}

	log.Printf("RunQuery : Count[%d]\n", len(applications))
}

/*
return an application from the mongodb based on application name
*/
func RepoFindApplication(id int) Application {

	// establish session with database
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("applications")

	result := Application{}
	err = c.Find(bson.M{"applicationname": "Toad"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Application Name:", result.ApplicationName)
	return result
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	/*
		for i, t := range todos {
			if t.Id == id {
				todos = append(todos[:i], todos[i+1:]...)
				return nil
			}
		}
		return fmt.Errorf("Could not find Todo with id of %d to delete", id)
	*/
	return nil
}
