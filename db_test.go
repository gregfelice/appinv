package main

import (
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"testing"
)

/*
get the database session.

@see creating a global session variable
https://stackoverflow.com/questions/40999637/mgo-query-performance-seems-consistently-slow-500-650ms/41000876#41000876
*/
func getDbSessionForTest() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	sessionCopy := session.Copy()
	return sessionCopy
}

func TestCRUDApplications(t *testing.T) {

	// initialize db connection / session
	session := getDbSessionForTest()
	c := session.DB("test").C("applications")
	defer session.Close()

	appName := "App Foo"
	bizUnit := "BU Foo"
	id := bson.NewObjectId()

	t.Run("create", func(t *testing.T) {
		// create
		app := Application{ID: id, ApplicationName: appName, BusinessUnit: bizUnit}
		err := Create(app)
		checkErr(err)
	})

	t.Run("retrieve one by id", func(t *testing.T) {
		// retrieve one by id
		result := Application{}
		err := c.Find(bson.M{"_id": id}).One(&result)
		checkErr(err)
		//t.Log("Application Name:", result.ApplicationName)
	})

	t.Run("retrieve one by name", func(t *testing.T) {
		// retrieve one by name
		result := Application{}
		err := c.Find(bson.M{"applicationname": appName}).One(&result)
		checkErr(err)
	})

	t.Run("retrieve many by business unit", func(t *testing.T) {
		// retrieve several by business nunit
		var applications []Application
		err := c.Find(bson.M{"businessunit": bizUnit}).All(&applications)
		checkErr(err)
		//t.Logf("RunQuery : Find by BizUnit Count[%d]\n", len(applications))
	})

	t.Run("retrieve all", func(t *testing.T) {
		// retrieve all
		var applications []Application
		err := c.Find(bson.M{}).All(&applications)
		checkErr(err)
		//t.Logf("RunQuery : Find all Count[%d]\n", len(applications))
	})

	// update
	t.Run("update existing app", func(t *testing.T) {
		// t.Logf("changing %s", id)
		q := bson.M{"_id": id}
		s := "Changed App Name"
		change := bson.M{"$set": bson.M{"applicationname": s}}
		err := c.Update(q, change)
		checkErr(err)
		result := Application{}
		err = c.Find(bson.M{"_id": id}).One(&result)
		checkErr(err)
		if result.ApplicationName != s {
			t.Fatal("app name not changed successfully")
		}
		// t.Log(result.ApplicationName)
	})

	// delete - also cleans all test records from the database for this test run.
	t.Run("delete app", func(t *testing.T) {
		err := c.Remove(bson.M{"_id": id})
		checkErr(err)
	})

}
