package dao

import (
	u "appinv/util"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"testing"
)

/*
get the database session.

@see creating a global session variable
https://stackoverflow.com/questions/40999637/mgo-query-performance-seems-consistently-slow-500-650ms/41000876#41000876

tests against the new dao structure.

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
		u.CheckErr(err)
	})

	t.Run("retrieve one by id", func(t *testing.T) {
		application := Application{}
		err := RetrieveById(&application, id)
		u.CheckErr(err)
		//t.Log(id)
		//t.Log(application)
		if application.ApplicationName != appName {
			t.Fatal("app names dont match")
		}
		//t.Log("Application Name:", result.ApplicationName)
	})

	t.Run("retrieve one by name", func(t *testing.T) {
		application := Application{}
		err := RetrieveByApplicationName(&application, appName)
		u.CheckErr(err)
		//t.Log(id)
		//t.Log(application)
		if application.ApplicationName != appName {
			t.Fatal("app names dont match")
		}
		//t.Log("Application Name:", result.ApplicationName)
	})

	t.Run("retrieve many by business unit", func(t *testing.T) {
		applications := []Application{}
		err := RetrieveByBusinessUnit(&applications, bizUnit)
		//for _, app := range applications {
		//t.Logf("returned app: %s", app)
		//}
		u.CheckErr(err)
	})

	t.Run("retrieve all", func(t *testing.T) {
		applications := []Application{}
		err := RetrieveAll(&applications)
		//for _, app := range applications {
		//	t.Logf("returned app: %s", app)
		//}
		u.CheckErr(err)
	})

	// update
	t.Run("update existing app", func(t *testing.T) {
		// get a copy of existing app
		application := Application{}
		err := RetrieveById(&application, id)
		u.CheckErr(err)
		//t.Logf("retrieved existing app name before: %s", application.ApplicationName)
		application.ApplicationName = "changed app name"
		//t.Logf("application name changed, before persist: %s", application.ApplicationName)

		err = Update(&application)
		u.CheckErr(err)

		changedAndRetrievedApplication := Application{}
		err = c.Find(bson.M{"_id": id}).One(&changedAndRetrievedApplication)
		//t.Logf("retrieved, changed application name after persist: %s", changedAndRetrievedApplication.ApplicationName)
		u.CheckErr(err)

		if application.ApplicationName != changedAndRetrievedApplication.ApplicationName {
			t.Fatal("app names dont match")
		}

		// t.Log(result.ApplicationName)
	})

	// delete - also cleans all test records from the database for this test run.
	t.Run("delete app", func(t *testing.T) {
		err := Remove(id)
		u.CheckErr(err)
		var app Application
		t.Logf("uninitialized app: %s", app)
		//err = RetrieveById(&app, id)
		//u.CheckErr(err)
		//t.Logf("uninitialized app after query: %s", app)
	})

}
