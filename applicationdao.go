package main

import (
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"log"
)

/*
@todo creating a global session variable
https://stackoverflow.com/questions/40999637/mgo-query-performance-seems-consistently-slow-500-650ms/41000876#41000876
*/

/*
bypassess any errors associated with record not found
*/
func checkDbErr(err error) {
	if err != nil {
		switch err {
		default:
			log.Fatal("Failed update application: ", err)
		case mgo.ErrNotFound:
			log.Println("Record not found")
		}
	}
}

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
	session := getDbSession()
	c := session.DB("test").C("applications")
	defer session.Close()
	err := c.Insert(application)
	checkErr(err)
	return err
}

func RetrieveById(application *Application, id bson.ObjectId) error {
	session := getDbSession()
	c := session.DB("test").C("applications")
	err := c.Find(bson.M{"_id": id}).One(&application)
	checkErr(err)
	return err
	//t.Log("Application Name:", result.ApplicationName)
}

func RetrieveByApplicationName(application *Application, name string) error {
	session := getDbSession()
	c := session.DB("test").C("applications")
	err := c.Find(bson.M{"applicationname": name}).One(&application)
	checkErr(err)
	return err
	//t.Log("Application Name:", result.ApplicationName)
}

/*
accept a pointer to a slice that will contain any return values from the query.
*/
func RetrieveByBusinessUnit(applications *[]Application, name string) error {
	session := getDbSession()
	c := session.DB("test").C("applications")
	err := c.Find(bson.M{"businessunit": name}).All(applications)
	checkErr(err)
	return err
	//t.Log("Application Name:", result.ApplicationName)
}

/*
accept a pointer to a slice that will contain any return values from the query.
*/
func RetrieveAll(applications *[]Application) error {
	session := getDbSession()
	c := session.DB("test").C("applications")
	err := c.Find(bson.M{}).All(applications)
	checkErr(err)
	return err
	//t.Log("Application Name:", result.ApplicationName)
}

/*
take values from the struct and store to the database.
*/
func Update(application *Application) error {
	// t.Logf("changing %s", id)
	session := getDbSession()
	c := session.DB("test").C("applications")
	q := bson.M{"_id": application.ID}
	change := bson.M{"$set": bson.M{
		"applicationname": application.ApplicationName,
		"businessunit":    application.BusinessUnit}}
	err := c.Update(q, change)
	checkDbErr(err)
	return err
	// t.Log(result.ApplicationName)
}

func Remove(id bson.ObjectId) error {
	session := getDbSession()
	c := session.DB("test").C("applications")
	err := c.Remove(bson.M{"_id": id})
	checkErr(err)
	return err
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
