package main

import (
	"flag"
	"log"

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

	err := svc.CheckConfirmation(id)
	if err != nil {
		log.Fatalln(err)
	}
}
