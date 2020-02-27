package main

import (
	"flag"
	"log"
)

func main() {
	var id int
	flag.IntVar(&id, "id", 0, "Specify application id you want to check")
	flag.Parse()

	rID := CheckVisaConfirmation(id)
	PrintApplicationReport(rID)
}

func CheckVisaConfirmation(applicationID int) int {
	// Gather application data.
	a, err := GetVisaApplication(applicationID, "data/applications.json")
	if err != nil {
		log.Fatalf("Can't get application, reason: %v", err)
	}

	// Check if user had VISA's previously.
	v, err := GetPreviousVisas(a.Name, "data/visas.json")
	if err != nil {
		log.Fatalf("Can't find previous visas, reason: %v", err)
	}

	// If a traveler have previous VISA's, check if it was violations.
	vs, err := CheckVisasViolations(a, v)
	if err != nil {
		log.Fatalf("Can't check visas violations: %v", err)
	}

	// Save VISA Application report.
	reportId, err := SaveApplicationReport(vs)
	if err != nil {
		log.Fatalf("Can't save application report, reason: %v", err)
	}

	return reportId
}
