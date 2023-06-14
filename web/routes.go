package main

import (
	"github.com/gorilla/mux"
	address3 "ndv/infrastructure/address"
	partner2 "ndv/infrastructure/partner"
	"ndv/usecases/address"
	"ndv/usecases/partner"
	address2 "ndv/web/handlers/address"
	partnerWeb "ndv/web/handlers/partner"
)

func routes(partnerRepository partner2.PartnerDBRepository, addressRepository address3.AddressDBRepository) *mux.Router {

	//PARTNER
	getAllPartnersUseCase := partner.NewGetAllPartners(partnerRepository)
	getAllPartnersHandler := partnerWeb.NewGetAllPartnersHandler(getAllPartnersUseCase)

	savePartnerUseCase := partner.NewSavePartner(partnerRepository)
	savePartnerHandler := partnerWeb.NewSavePartnerHandler(savePartnerUseCase)

	//ADDRESS
	getAddressByID := address.NewGetAddressByID(addressRepository)
	getAddressByIDHandler := address2.NewGetAddressByIDHandler(getAddressByID)

	saveAddress := address.NewSaveAddress(addressRepository)
	saveAddressHandler := address2.NewSaveAddressHandler(saveAddress)

	//BUNDLE

	router := mux.NewRouter()
	router.HandleFunc("/api/partners", getAllPartnersHandler.Handle).Methods("GET")
	router.HandleFunc("/api/partners", savePartnerHandler.Handle).Methods("POST")
	router.HandleFunc("/api/addresses/{id}", getAddressByIDHandler.Handle).Methods("GET")
	router.HandleFunc("/api/addresses", saveAddressHandler.Handle).Methods("POST")

	//Partner por ID
	//Partner por Query

	return router
}
