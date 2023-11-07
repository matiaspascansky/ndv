package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ndv/infrastructure"
	"ndv/infrastructure/address"
	"ndv/infrastructure/partner"
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func main() {
	fmt.Println("Server is getting started")
	datasource := "root:password@tcp(127.0.0.1:3308)/ndv?parseTime=true"
	db := infrastructure.CreateLocalDataBase(datasource)
	partnerRepo, addressRepo := initializeRepositoriess(db)

	r := routes(partnerRepo, addressRepo)
	fmt.Println("Running at port 8080")
	http.ListenAndServe(":8080", r)

}

func HelloServerr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Running , %s!", r.URL.Path[1:])
}

func initializeRepositoriess(db *sqlx.DB) (partner.PartnerDBRepository, address.AddressDBRepository) {
	partnerRepo, err := partner.NewPartnerRepository(db)
	addressRepo, err := address.NewAddressRepository(db)
	if err != nil {
		fmt.Println("main: cannot initialize handlers repository")
	}
	return partnerRepo, addressRepo
}
