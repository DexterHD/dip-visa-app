package visa

import (
	"fmt"
	"time"
)

// Service implements Visa Confirmation Service.
type Service struct {
	AStore      ApplicationStorage
	VStore      VisasStorage
	RStore      ReportsStorage
	PrintReport ReportsPrinter
}

// CheckConfirmation gets information about Visa confirmation byt applicationID.
// This method implement our Business Logic. Don't focus on business-logic itself.
// THE MAIN THING HERE IS THAT OUR METHOD DEPENDS ON INTERFACES, NOT LOW-LEVEL DETAILS.
func (svc *Service) CheckConfirmation(applicationID int) error {
	// Gather application data from some persistent storage.
	// It can be everything which fit ApplicationStorage interface.
	a, err := svc.AStore.GetVisaApplication(applicationID)
	if err != nil {
		return fmt.Errorf("can't get application, reason: %w", err)
	}

	// Next we check is applicant has some previous Visas.
	// Later we will use previous visas to make a decision.
	visas, err := svc.VStore.GetPreviousVisas(a.Name)
	if err != nil {
		return fmt.Errorf("can't find previous visas, reason: %w", err)
	}

	// By default we approve a VISA to the applicant.
	var accepted = true

	// Then we check Time to stay. If it more than maximum time to stay, we reject visa.
	if a.Departure.Sub(a.Arrival).Hours() > MaxTimeToStay {
		accepted = false
	}

	// But if applicant has more than 3 previous visa, we accept application.
	var visasCount = len(visas)
	if visasCount > 3 {
		accepted = true
	}

	// After checks we create application report.
	rep := Report{
		ID:            time.Now().Unix(),
		ApplicationID: a.ID,
		Applicant:     a.Name,
		Accepted:      accepted,
	}

	// Then we save our VISA Application report to reports storage to later usages.
	err = svc.RStore.SaveApplicationReport(rep)
	if err != nil {
		return fmt.Errorf("can't save application report, reason: %w", err)
	}

	// At the end we loads generated report from storage
	storedReport, err := svc.RStore.LoadApplicationReport(rep.ApplicationID)
	if err != nil {
		return fmt.Errorf("can't load application report, reason: %w", err)
	}

	// ... and print it.
	err = svc.PrintReport(*storedReport)
	if err != nil {
		return err
	}

	return nil
}
