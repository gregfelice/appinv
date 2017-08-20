package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ApplicationDB struct {
	gorm.Model
	ApplicationName          string
	BusinessUnit             string
	ExecutiveLeader          string
	SPOC                     string
	Capability               string
	ApplicationShortName     string
	ApplicationDescription   string
	ApplicationType          string
	ApplicationCategory      string
	ApplicationVersion       string
	EmployeeOrCustomerFacing string
	InstallationCriteria     string
	MultipleInstances        string
	DisasterRecovery         bool
	BusinessCriticality      string
	SOX                      bool
	PCI                      bool
	ContainsPII              bool
	ApplicationDevelopedBy   string
	Vendor                   string
	EstimatedNumberOfUsers   string
	UsageReportingAvailable  bool
	AccessApprover           string
	IsDataSource             bool
	TargetedForRetirement    bool
	Comments                 string
	AnnualCost               string
	PotentialReplacements    string
	Recommendation           string
	ReceivedBy               string
}

func runDbTest() {
	db, err := gorm.Open("sqlite3", "./db/appinv.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Application{})

	// Create
	db.Create(&Application{
		ApplicationName: "TestApplication",
		BusinessUnit:    "Devops",
		ExecutiveLeader: "Greg Felice",
		SPOC:            "Greg Felice"})

	// Read
	var application Application
	db.First(&application, 1)                                        // find application with id 1
	db.First(&application, "ApplicationName = ?", "TestApplication") // find application with code l1212

	// Update - update application's price to 2000
	db.Model(&application).Update("BusinessUnit", "Devops 2")

	// Delete - delete application
	db.Delete(&application)
}
