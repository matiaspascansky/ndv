package main

import (
	"github.com/gorilla/mux"
	address3 "ndv/infrastructure/address"
	partner2 "ndv/infrastructure/partner"
	"ndv/usecases/address"
	"ndv/usecases/partner"
	"ndv/web/handlers"
	address2 "ndv/web/handlers/address"
	partnerWeb "ndv/web/handlers/partner"
)

func routes(partnerRepository partner2.PartnerDBRepository, addressRepository address3.AddressDBRepository) *mux.Router {

	//PARTNER
	getPartnersByTypeUseCase := partner.NewGetPartnerByType(partnerRepository)
	getAllPartnersUseCase := partner.NewGetAllPartners(partnerRepository)
	getAllPartnersHandler := partnerWeb.NewGetAllPartnersHandler(getAllPartnersUseCase, getPartnersByTypeUseCase)

	savePartnerUseCase := partner.NewSavePartner(partnerRepository)
	savePartnerHandler := partnerWeb.NewSavePartnerHandler(savePartnerUseCase)

	getPartnerByIDUseCase := partner.NewGetPartnerByID(partnerRepository)
	getPartnerByIDHandler := partnerWeb.NewGetPartnerByIDHandler(getPartnerByIDUseCase)

	//ADDRESS
	getAddressByID := address.NewGetAddressByID(addressRepository)
	getAddressByIDHandler := address2.NewGetAddressByIDHandler(getAddressByID)

	saveAddress := address.NewSaveAddress(addressRepository)
	saveAddressHandler := address2.NewSaveAddressHandler(saveAddress)

	//HEALTH
	heathCheckHandler := handlers.NewHealthCheckHandler()

	//BUNDLE

	//PARTNER
	router := mux.NewRouter()
	router.HandleFunc("/api/partners", getAllPartnersHandler.Handle).Methods("GET")
	router.HandleFunc("/api/partners", savePartnerHandler.Handle).Methods("POST")
	router.HandleFunc("/api/partners/{id}", getPartnerByIDHandler.Handle).Methods("GET")
	//ADDRESS
	router.HandleFunc("/api/addresses/{id}", getAddressByIDHandler.Handle).Methods("GET")
	router.HandleFunc("/api/users/{id}/address", getAddressByIDHandler.Handle).Methods("GET")

	router.HandleFunc("/api/addresses", saveAddressHandler.Handle).Methods("POST")

	router.HandleFunc("/", heathCheckHandler.Handle).Methods("GET")

	//Partner por ID
	//Partner por Query

	return router
}
