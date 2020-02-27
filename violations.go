package main

import (
	"time"
)

const ThreeMonthsInHours = 24 * 90

type Violations struct {
	ViolationID int64

	ApplicationID int
	Applicant     string

	DepartureViolated bool
	ArrivalViolated   bool
	DurationViolated  bool
}

func CheckVisasViolations(a *Application, v []Visa) (Violations, error) {
	vs := Violations{
		ViolationID:   time.Now().Unix(),
		ApplicationID: a.ID,
		Applicant:     a.Name,

		DepartureViolated: false,
		ArrivalViolated:   false,
		DurationViolated:  false,
	}

	if a.Departure.Sub(a.Arrival).Hours() > ThreeMonthsInHours {
		vs.DurationViolated = true
	}

	return vs, nil
}
