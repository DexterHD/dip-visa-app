// Package visa implements our application business logic.
// As you can see here, layers are separated in different packages like:
// - visa
// - report
// - storage
// But at the same time current business-logic package is depends on low-level
// packages like "report" and "storage".
package visa

import (
	"fmt"
	"time"

	"github.com/DexterHD/dip-visa-app/pkg/report"
	"github.com/DexterHD/dip-visa-app/pkg/storage"
)

const maxTimeToStay = 24 * 90

// CheckConfirmation implements VISA confirmation business logic.
// Don't focus on business-logic itself, but see how we call low-level functions
// using packages.
func CheckConfirmation(applicationID int) error {
	// Gather application data.
	a, err := storage.GetVisaApplication(applicationID)
	if err != nil {
		return fmt.Errorf("can't get application, reason: %w", err)
	}

	// Check if user had VISA's previously.
	visas, err := storage.GetPreviousVisas(a.Name)
	if err != nil {
		return fmt.Errorf("can't find previous visas, reason: %w", err)
	}

	var visasCount = len(visas)
	var accepted = true

	if a.Departure.Sub(a.Arrival).Hours() > maxTimeToStay {
		accepted = false
	}

	if visasCount > 3 {
		accepted = true
	}

	rep := report.Report{
		ID:            time.Now().Unix(),
		ApplicationID: a.ID,
		Applicant:     a.Name,
		Accepted:      accepted,
	}

	// Save VISA Application report.
	err = report.SaveApplicationReport(rep)
	if err != nil {
		return fmt.Errorf("can't save application report, reason: %w", err)
	}

	// Save VISA Application report.
	storedReport, err := report.LoadApplicationReport(rep.ApplicationID)
	if err != nil {
		return fmt.Errorf("can't load application report, reason: %w", err)
	}

	err = report.PrintApplicationReport(*storedReport)
	if err != nil {
		return err
	}

	return nil
}
