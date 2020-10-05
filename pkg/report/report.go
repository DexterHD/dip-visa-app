// Package report implements Visa Application Report persistent storage and
// also implements default Report Printer.
package report

import (
	"fmt"
	"time"

	"github.com/idexter/dip-visa-app/pkg/visa"
)

const defaultDir = "data"

// StoredReport describes visa application report stored in database.
type StoredReport struct {
	ID            int64     `json:"id"`
	ApplicationID int       `json:"application_id"`
	Applicant     string    `json:"applicant"`
	Accepted      bool      `json:"accepted"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// PrintApplicationReport prints information about visa report.
func PrintApplicationReport(vs visa.Report) error {
	fmt.Printf("\n\nID: %d\nApplicant: %s\nAccepted: %v\n\n", vs.ApplicationID, vs.Applicant, vs.Accepted)
	return nil
}
