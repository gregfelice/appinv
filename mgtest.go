package main

/*

this code seems to work.

*/

import (
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Application struct {
	ApplicationName string
	BusinessUnit string
	/*
	ExecutiveLeader string
	SPOC string
	Capability string
	ApplicationShortName string
	ApplicationDescription string
	ApplicationType string
	ApplicationCategory string
	ApplicationVersion string
	EmployeeOrCustomerFacing string
	InstallationCriteria string
	MultipleInstances string
	DisasterRecovery bool
	BusinessCriticality string
	SOX bool
	PCI bool
	ContainsPII bool
	ApplicationDevelopedBy string
	Vendor string
	EstimatedNumberOfUsers string
	UsageReportingAvailable bool
	AccessApprover string
	IsDataSource bool
	TargetedForRetirement bool
	Comments string
	AnnualCost string
	PotentialReplacements string
	Recommendation string
        */
}

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
