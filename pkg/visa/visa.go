package visa

import (
	"fmt"
	"time"
)

const MaximumLimit = 24 * 90

type Application struct {
	ID        int
	Name      string
	Arrival   time.Time
	Departure time.Time
	Money     float64
}

type Visa struct {
	From      time.Time
	To        time.Time
	Arrival   time.Time
	Departure time.Time
}

type Report struct {
	ID            int64
	ApplicationID int
	Applicant     string
	Accepted      bool
}

type ApplicationStorage interface {
	GetVisaApplication(id int) (*Application, error)
}

type VisasStorage interface {
	GetPreviousVisas(name string) ([]Visa, error)
}

type ReportsStorage interface {
	SaveApplicationReport(Report) error
	LoadApplicationReport(id int) (*Report, error)
}

type ReportsPrinter func(Report) error

type Service struct {
	AStore      ApplicationStorage
	VStore      VisasStorage
	RStore      ReportsStorage
	PrintReport ReportsPrinter
}

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
