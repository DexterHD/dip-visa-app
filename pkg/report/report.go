package report

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/DexterHD/dip-visa-app/pkg/visa"
)

const DefaultDir = "data"

type StoredReport struct {
	ID            int64     `json:"id"`
	ApplicationID int       `json:"application_id"`
	Applicant     string    `json:"applicant"`
	Accepted      bool      `json:"accepted"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type FileStorage struct {
	Dir string
}

func NewFileStorage() *FileStorage {
	return &FileStorage{Dir: DefaultDir}
}

func (rs *FileStorage) SaveApplicationReport(vs visa.Report) error {
	data, err := json.Marshal(StoredReport{
		ID:            vs.ID,
		ApplicationID: vs.ApplicationID,
		Applicant:     vs.Applicant,
		Accepted:      vs.Accepted,
		UpdatedAt:     time.Now(),
	})
	if err != nil {
		return fmt.Errorf("could not marshall violations, reason %w", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/violations-%d.json", rs.Dir, vs.ApplicationID), data, os.FileMode(0777))
	if err != nil {
		return fmt.Errorf("could not write violations, reason %w", err)
	}

	return nil
}

func (rs *FileStorage) LoadApplicationReport(applicationID int) (*visa.Report, error) {

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/violations-%d.json", rs.Dir, applicationID))
	if err != nil {
		return nil, fmt.Errorf("could not read violations, reason %w", err)
	}

	sr := &StoredReport{}
	err = json.Unmarshal(b, &sr)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall, reason %w", err)
	}

	return &visa.Report{
		ID:            sr.ID,
		ApplicationID: sr.ApplicationID,
		Applicant:     sr.Applicant,
		Accepted:      sr.Accepted,
	}, nil
}

func PrintApplicationReport(vs visa.Report) error {
	fmt.Printf("\n\nID: %d\nApplicant: %s\nAccepted: %v\n\n", vs.ApplicationID, vs.Applicant, vs.Accepted)
	return nil
}
