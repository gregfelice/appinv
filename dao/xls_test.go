package dao

import (
	"testing"
)

/*
run the db test.
*/
func TestXLSIngestion(t *testing.T) {

	s := IngestXLS("./applications.xlsx")

	//

	if s == nil {
		t.Error("ingest XLS is returning nil")
	} else {

		// log.Println("GOOD: Ingest XLS returning non-null")
	}
}
