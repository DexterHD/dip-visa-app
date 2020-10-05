package visa

import (
	"fmt"
	"time"
)

// MaximumLimit described maximum time in hours from arrival to departure.
const MaximumLimit = 24 * 90

// Application described Visa Application.
type Application struct {
	ID        int
	Name      string
	Arrival   time.Time
	Departure time.Time
	Money     float64
}

// Visa described Visa data.
type Visa struct {
	From      time.Time
	To        time.Time
	Arrival   time.Time
	Departure time.Time
}

// Report described Visa Application Report.
// In other words this structure described Visa Application Result.
type Report struct {
	ID            int64
	ApplicationID int
	Applicant     string
	Accepted      bool
}

// ApplicationStorage described Visa Application storage.
type ApplicationStorage interface {
	GetVisaApplication(id int) (*Application, error)
}

// VisasStorage describes Visas persistent storage.
type VisasStorage interface {
	GetPreviousVisas(name string) ([]Visa, error)
}

// ReportsStorage describes application reports persistent storage.
type ReportsStorage interface {
	SaveApplicationReport(Report) error
	LoadApplicationReport(id int) (*Report, error)
}

// ReportsPrinter implements reports printer.
// With reporter we can redirect service output to different destination.
type ReportsPrinter func(Report) error

// Service implements Visa Confirmation Service.
type Service struct {
	AStore      ApplicationStorage
	VStore      VisasStorage
	RStore      ReportsStorage
	PrintReport ReportsPrinter
}

// CheckConfirmation gets information about Visa confirmation byt applicationID.
func (svc *Service) CheckConfirmation(applicationID int) error {
	// Gather application data.
	a, err := svc.AStore.GetVisaApplication(applicationID)
	if err != nil {
		return fmt.Errorf("can't get application, reason: %w", err)
	}

	// Check if user had VISA's previously.
	visas, err := svc.VStore.GetPreviousVisas(a.Name)
	if err != nil {
		return fmt.Errorf("can't find previous visas, reason: %w", err)
	}

	var visasCount = len(visas)
	var accepted = true

	if a.Departure.Sub(a.Arrival).Hours() > MaximumLimit {
		accepted = false
	}

	if visasCount > 3 {
		accepted = true
	}

	rep := Report{
		ID:            time.Now().Unix(),
		ApplicationID: a.ID,
		Applicant:     a.Name,
		Accepted:      accepted,
	}

	// Save VISA Application report.
	err = svc.RStore.SaveApplicationReport(rep)
	if err != nil {
		return fmt.Errorf("can't save application report, reason: %w", err)
	}

	// Save VISA Application report.
	storedReport, err := svc.RStore.LoadApplicationReport(rep.ApplicationID)
	if err != nil {
		return fmt.Errorf("can't load application report, reason: %w", err)
	}

	err = svc.PrintReport(*storedReport)
	if err != nil {
		return err
	}

	return nil
}
