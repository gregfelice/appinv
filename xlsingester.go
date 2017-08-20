package main

import (
	"github.com/tealeg/xlsx"
	s "strings"
)

func IngestXLS(filename string) []Application {
	excelFileName := filename
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		p("error!")
	}

	apps := make([]Application, 0)

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {

			appName := s.TrimSpace(row.Cells[6].String())
			bizUnit := s.TrimSpace(row.Cells[0].String())

			if len(appName) > 0 && len(bizUnit) > 0 {
				a := Application{ApplicationName: row.Cells[6].String(), BusinessUnit: row.Cells[0].String()}
				apps = append(apps, a)
			} else {
				//p("skipping")
			}
		}
	}
	// fmt.Printf("the final apps list: %s", apps)

	return apps
}
