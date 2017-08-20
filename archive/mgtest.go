package main

/*

this code seems to work.

*/

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func runMgTest() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("applications")
	err = c.Insert(
		&Application{"App X", "IT"},
		&Application{"App Y", "IT"})
	if err != nil {
		log.Fatal(err)
	}

	result := Application{}
	err = c.Find(bson.M{"applicationname": "App X"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Application Name:", result.ApplicationName)
}
