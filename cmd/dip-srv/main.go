package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/DexterHD/dip-visa-app/pkg/report"
	"github.com/DexterHD/dip-visa-app/pkg/storage"
	"github.com/DexterHD/dip-visa-app/pkg/visa"
)

func main() {
	var id int
	flag.IntVar(&id, "id", 0, "Specify application id you want to check")
	flag.Parse()

	svc := &visa.Service{
		AStore:      storage.NewFileApplicationsStorage(),
		VStore:      storage.NewFileVisasStorage(),
		RStore:      report.NewFileStorage(),
		PrintReport: report.PrintApplicationReport,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Parse parameter from the request
		idParam := r.FormValue("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			httpResponse(w, err)
			return
		}

		// Replace printer to print report to HTTP
		svc.PrintReport = func(r visa.Report) error {
			httpResponse(w, r)
			return nil
		}

		// Run Visa confirmation business logic
		if err = svc.CheckConfirmation(id); err != nil {
			httpResponse(w, err)
		}
		httpResponse(w, "OK")
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func httpResponse(w http.ResponseWriter, msg interface{}) {
	if _, ok := msg.(error); ok {
		w.WriteHeader(500)
	}
	b, err := json.Marshal(msg)
	if err != nil {
		log.Printf("can't unmarshal message, reason: %v", err)
	}
	if _, err := w.Write(b); err != nil {
		log.Printf("can't write response, reason: %v", err)
	}
}
