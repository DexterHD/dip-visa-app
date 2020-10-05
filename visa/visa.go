// Package visa implements our whole Visa application service.
// It consist several functions which is responsible for different
// parts of our application.
// There is no DI and this version don't use DIP principle.
package visa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const maxTimeToStay = 24 * 90

// CheckVisaConfirmation checks VISA application and prints results.
// It runs whole application business logic.
// As you can see in this implementation, high-level logic depends on
// low level functions which are using files to store data.
func CheckVisaConfirmation(applicationID int) error {
	// Gather application data.
	a, err := getVisaApplication(applicationID, "data/applications.json")
	if err != nil {
		return fmt.Errorf("can't get application, reason: %w", err)
	}

	// Check if user had VISA's previously.
	visas, err := getPreviousVisas(a.Name, "data/visas.json")
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

	report := Report{
		ID:            time.Now().Unix(),
		ApplicationID: a.ID,
		Applicant:     a.Name,
		Accepted:      accepted,
	}

	// Save VISA Application report.
	err = saveApplicationReport(report, "data")
	if err != nil {
		return fmt.Errorf("can't save application report, reason: %w", err)
	}

	// Save VISA Application report.
	storedReport, err := loadApplicationReport(report.ApplicationID, "data")
	if err != nil {
		return fmt.Errorf("can't load application report, reason: %w", err)
	}

	err = printApplicationReport(*storedReport)
	if err != nil {
		return err
	}

	return nil
}

// getVisaApplication gets visa application from file using ID.
func getVisaApplication(id int, filename string) (*Application, error) {

	var apps []Application
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("couldn't read applications database %w", err)
	}

	if err := json.Unmarshal(b, &apps); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal applications database %w", err)
	}

	for _, v := range apps {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("application was not found")
}

// getPreviousVisas gets previous VISAs from file using applicant name.
func getPreviousVisas(name string, filename string) ([]Visa, error) {

	var visas map[string][]Visa
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("couldn't read visas database %w", err)
	}

	if err := json.Unmarshal(b, &visas); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal visas database %w", err)
	}

	if v, ok := visas[name]; ok {
		return v, nil
	}

	return []Visa{}, nil
}

// saveApplicationReport saves visa application report to provided directory.
func saveApplicationReport(vs Report, dir string) error {

	data, err := json.Marshal(vs)
	if err != nil {
		return fmt.Errorf("could not marshall violations, reason %w", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/violations-%d.json", dir, vs.ApplicationID), data, os.FileMode(0777))
	if err != nil {
		return fmt.Errorf("could not write violations, reason %w", err)
	}

	return nil
}

// loadApplicationReport loads application report by id from directory.
func loadApplicationReport(applicationID int, dir string) (*Report, error) {

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/violations-%d.json", dir, applicationID))
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

// printApplicationReport prints visa application report.
func printApplicationReport(vs Report) error {
	fmt.Printf("\n\nID: %d\nApplicant: %s\nAccepted: %v\n\n", vs.ApplicationID, vs.Applicant, vs.Accepted)
	return nil
}
