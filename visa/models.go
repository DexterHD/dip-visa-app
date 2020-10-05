package visa

import "time"

type Application struct {
	ID        int
	Name      string
	Arrival   time.Time
	Departure time.Time
	Money     float64
}

type Report struct {
	ID            int64
	ApplicationID int
	Applicant     string
	Accepted      bool
}

type Visa struct {
	From      time.Time
	To        time.Time
	Arrival   time.Time
	Departure time.Time
}
