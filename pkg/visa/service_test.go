package visa

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_CheckConfirmation(t *testing.T) {

	// Setup
	application := &Application{ID: 1, Name: "Anton", Arrival: time.Time{}, Departure: time.Time{}, Money: 100}
	visas := []Visa{{From: time.Time{}, To: time.Time{}, Arrival: time.Time{}, Departure: time.Time{}}}
	report := &Report{ID: time.Now().Unix(), ApplicationID: 1, Applicant: "Anton", Accepted: false}

	appStore := &MockApplicationStorage{}
	vStore := &MockVisasStorage{}
	rStore := &MockReportsStorage{}

	appStore.On("GetVisaApplication", mock.Anything).Return(application, nil)
	vStore.On("GetPreviousVisas", mock.Anything).Return(visas, nil)

	rStore.On("SaveApplicationReport", mock.Anything).Return(nil)
	rStore.On("LoadApplicationReport", mock.Anything).Return(report, nil)

	svc := &Service{
		AStore:      appStore,
		VStore:      vStore,
		RStore:      rStore,
		PrintReport: func(r Report) error { return nil },
	}

	t.Run("Success", func(t *testing.T) {
		err := svc.CheckConfirmation(application.ID)
		assert.Nil(t, err)

		appStore.AssertCalled(t, "GetVisaApplication", 1)
		vStore.AssertCalled(t, "GetPreviousVisas", "Anton")
		rStore.AssertCalled(t, "SaveApplicationReport", Report{ID: report.ID, ApplicationID: 1, Applicant: "Anton", Accepted: true})
		rStore.AssertCalled(t, "LoadApplicationReport", 1)
	})

	t.Run("Maximum limit", func(t *testing.T) {
		application.Arrival = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
		application.Departure = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

		err := svc.CheckConfirmation(application.ID)
		assert.Nil(t, err)

		appStore.AssertCalled(t, "GetVisaApplication", 1)
		vStore.AssertCalled(t, "GetPreviousVisas", "Anton")
		rStore.AssertCalled(t, "SaveApplicationReport", Report{ID: report.ID, ApplicationID: 1, Applicant: "Anton", Accepted: false})
		rStore.AssertCalled(t, "LoadApplicationReport", 1)
	})

	t.Run("Four visas", func(t *testing.T) {
		vStore := &MockVisasStorage{}
		svc.VStore = vStore

		application.Arrival = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
		application.Departure = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

		visas := []Visa{
			{From: time.Time{}, To: time.Time{}, Arrival: time.Time{}, Departure: time.Time{}},
			{From: time.Time{}, To: time.Time{}, Arrival: time.Time{}, Departure: time.Time{}},
			{From: time.Time{}, To: time.Time{}, Arrival: time.Time{}, Departure: time.Time{}},
			{From: time.Time{}, To: time.Time{}, Arrival: time.Time{}, Departure: time.Time{}},
		}

		vStore.On("GetPreviousVisas", mock.Anything).Return(visas, nil)

		err := svc.CheckConfirmation(application.ID)
		assert.Nil(t, err)

		appStore.AssertCalled(t, "GetVisaApplication", 1)
		vStore.AssertCalled(t, "GetPreviousVisas", "Anton")
		rStore.AssertCalled(t, "SaveApplicationReport", Report{ID: report.ID, ApplicationID: 1, Applicant: "Anton", Accepted: true})
		rStore.AssertCalled(t, "LoadApplicationReport", 1)
	})
}
