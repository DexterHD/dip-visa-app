package visa

import (
	"fmt"
	"time"
)

const ThreeMonthsInHours = 24 * 90

type Report struct {
	ID            int64
	ApplicationID int
	Applicant     string
	Accepted      bool
}

func CheckVisaConfirmation(applicationID int) error {
	// Gather application data.
	a, err := GetVisaApplication(applicationID, "data/applications.json")
	if err != nil {
		return fmt.Errorf("can't get application, reason: %w", err)
	}

	// Check if user had VISA's previously.
	_, err = GetPreviousVisas(a.Name, "data/visas.json")
	if err != nil {
		return fmt.Errorf("can't find previous visas, reason: %w", err)
	}

	var durationLimitExceeded bool
	var durationViolated bool
	var arrivalViolated bool
	var departureViolated bool
	var accepted = true

	if a.Departure.Sub(a.Arrival).Hours() > ThreeMonthsInHours {
		durationLimitExceeded = true
	}

	if durationLimitExceeded || durationViolated || arrivalViolated || departureViolated {
		accepted = false
	}

	report := Report{
		ID:            time.Now().Unix(),
		ApplicationID: a.ID,
		Applicant:     a.Name,
		Accepted:      accepted,
	}

	// Save VISA Application report.
	err = SaveApplicationReport(report, fmt.Sprintf("reports/violations-%d.json", report.ApplicationID))
	if err != nil {
		return fmt.Errorf("can't save application report, reason: %w", err)
	}

	err = PrintApplicationReport(report.ApplicationID)
	if err != nil {
		return err
	}

	return nil
}
