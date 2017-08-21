package main

import (
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"log"
)

// var DBSession *mgo.Session

/*
get the database session.

@see creating a global session variable
https://stackoverflow.com/questions/40999637/mgo-query-performance-seems-consistently-slow-500-650ms/41000876#41000876
*/

func getDbSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	sessionCopy := session.Copy()
	return sessionCopy
}

func Create(application Application) error {
	// initialize db connection / session
	session := getDbSession()
	c := session.DB("test").C("applications")
	defer session.Close()
	err := c.Insert(application)
	checkErr(err)
	return err
}

func RetrieveById() error {
	return nil
}

func Foo() {
	// initialize db connection / session
	session := getDbSession()
	c := session.DB("test").C("applications")
	defer session.Close()

	appName := "App Foo"
	bizUnit := "BU Foo"
	id := bson.NewObjectId()
	id2 := bson.NewObjectId()

	// create
	err := c.Insert(
		&Application{ID: id, ApplicationName: appName, BusinessUnit: bizUnit},
		&Application{ID: id2, ApplicationName: "App Bar", BusinessUnit: bizUnit},
	)
	checkErr(err)

	// retrieve one by id
	result := Application{}
	err = c.Find(bson.M{"_id": id}).One(&result)
	checkErr(err)
	//t.Log("Application Name:", result.ApplicationName)

	// retrieve one by name
	result = Application{}
	err = c.Find(bson.M{"applicationname": appName}).One(&result)
	checkErr(err)

	// retrieve several by business nunit
	var applications []Application
	err = c.Find(bson.M{"businessunit": bizUnit}).All(&applications)
	checkErr(err)
	//t.Logf("RunQuery : Find by BizUnit Count[%d]\n", len(applications))

	// retrieve all
	//var applications []Application
	err = c.Find(bson.M{}).All(&applications)
	checkErr(err)
	//t.Logf("RunQuery : Find all Count[%d]\n", len(applications))

	// update

	// t.Logf("changing %s", id)
	q := bson.M{"_id": id}
	s := "Changed App Name"
	change := bson.M{"$set": bson.M{"applicationname": s}}
	err = c.Update(q, change)
	checkErr(err)
	result = Application{}
	err = c.Find(bson.M{"_id": id}).One(&result)
	checkErr(err)
	if result.ApplicationName != s {
		log.Fatal("app name not changed successfully")
	}
	// t.Log(result.ApplicationName)

	// delete - also cleans all test records from the database for this test run.
	err = c.Remove(bson.M{"_id": id})
	checkErr(err)
	err = c.Remove(bson.M{"_id": id2})
	checkErr(err)

}
