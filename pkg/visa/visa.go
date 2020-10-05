// Package visa stores business layer of our application.
// As you can see, this package is not depend on any low-level details.
// All details are hidden behind four interfaces:
// - ApplicationStorage
// - VisasStorage
// - ReportsStorage
// - ReportsPrinter
// If you check out main `cmd/...` packages, you can see, how we inject low-level
// dependencies inside our Service implementation.
package visa

// ApplicationStorage described VISA applications persistent storage.
type ApplicationStorage interface {
	GetVisaApplication(id int) (*Application, error)
}

// VisasStorage describes VISA persistent storage.
type VisasStorage interface {
	GetPreviousVisas(applicantName string) ([]Visa, error)
}

// ReportsStorage describes VISA application reports persistent storage.
type ReportsStorage interface {
	SaveApplicationReport(Report) error
	LoadApplicationReport(id int) (*Report, error)
}

// ReportsPrinter implements reports printer.
// We use it as output port for printing application result.
type ReportsPrinter func(Report) error
