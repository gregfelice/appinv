package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"testing"
)

/*
run the db test.
*/
func TestXLSIngestion(t *testing.T) {

	s := IngestXLS("./applications.xlsx")

	if s == nil {
		t.Error("ingest XLS is returning nil")
	} else {
		fmt.Println("GOOD: Ingest XLS returning non-null")
		// fmt.Println("the final apps list: ", s)
	}
}

func TestInsertIntoDatabase(t *testing.T) {

	// get the list of applications
	apps := IngestXLS("./applications.xlsx")

	// establish session with database
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("applications")

	// loop through the apps, insert into database

	for _, app := range apps {

		err = c.Insert(app)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func TestLoadOneFromDatabase(t *testing.T) {

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
}

func TestLoadManyFromDatabase(t *testing.T) {

	// establish session with database
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
	var applications []Application
	e := collection.Find(nil).All(&applications)
	if e != nil {
		log.Printf("RunQuery : ERROR : %s\n", e)
		return
	}

	log.Printf("RunQuery : Count[%d]\n", len(applications))

}
