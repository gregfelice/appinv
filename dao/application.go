package dao

import (
	"gopkg.in/mgo.v2/bson"
)

type Application struct {
	ID              bson.ObjectId `json:"id"              bson:"_id,omitempty"`
	ApplicationName string        `json:"applicationname" bson:"applicationname"`
	BusinessUnit    string        `json:"businessunit"    bson:"businessunit"`
}

type Applications []Application

////////////////////////////////////////////////////////////////////

//type Application struct {
//ApplicationName string
//BusinessUnit    string
//Comments        string

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

	AnnualCost string
	PotentialReplacements string
	Recommendation string
*/
//}
