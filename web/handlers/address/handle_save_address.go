package address

import (
	"encoding/json"
	"io/ioutil"
	"ndv/usecases/address"
	"ndv/web/handlers"
	"net/http"
)

type SaveAddressHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type saveAddressHandler struct {
	saveAddress address.SaveAddress
}

func NewSaveAddressHandler(saveAddress address.SaveAddress) *saveAddressHandler {
	return &saveAddressHandler{saveAddress: saveAddress}
}

func (s saveAddressHandler) Handle(w http.ResponseWriter, r *http.Request) {
	model := address.SaveAddressModel{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handlers.InternalServerError(w, err)
	}
	err = json.Unmarshal(data, &model)
	if err != nil {
		handlers.BadRequest(w, "INVALID_JSON", err.Error())
	}
	_, err = s.saveAddress.Execute(&model)
	if err != nil {
		handlers.InternalServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
