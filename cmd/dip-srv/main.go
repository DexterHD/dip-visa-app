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
			HttpResponse(w, err)
			return
		}

		// Replace printer to print report to HTTP
		svc.PrintReport = func(r visa.Report) error {
			HttpResponse(w, r)
			return nil
		}

		// Run Visa confirmation business logic
		err = svc.CheckConfirmation(id)
		HttpResponse(w, "OK")
		return
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HttpResponse(w http.ResponseWriter, msg interface{}) {
	if _, ok := msg.(error); ok {
		w.WriteHeader(500)
	}
	b, err := json.Marshal(msg)
	if err == nil {
		w.Write(b)
	}
}
