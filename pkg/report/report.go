package report

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Report describes visa application report.
type Report struct {
	ID            int64
	ApplicationID int
	Applicant     string
	Accepted      bool
}

var reportsDir = "data"

// SaveApplicationReport save application report to file.
func SaveApplicationReport(vs Report) error {
	data, err := json.Marshal(vs)
	if err != nil {
		return fmt.Errorf("could not marshall violations, reason %w", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/violations-%d.json", reportsDir, vs.ApplicationID), data, os.FileMode(0777))
	if err != nil {
		return fmt.Errorf("could not write violations, reason %w", err)
	}

	return nil
}

// LoadApplicationReport load application report by application id.
func LoadApplicationReport(applicationID int) (*Report, error) {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/violations-%d.json", reportsDir, applicationID))
	if err != nil {
		return nil, fmt.Errorf("could not read violations, reason %w", err)
	}

	vs := &Report{}
	err = json.Unmarshal(b, &vs)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall, reason %w", err)
	}

	return vs, nil
}

// PrintApplicationReport prints application report.
func PrintApplicationReport(vs Report) error {
	fmt.Printf("\n\nID: %d\nApplicant: %s\nAccepted: %v\n\n", vs.ApplicationID, vs.Applicant, vs.Accepted)
	return nil
}
