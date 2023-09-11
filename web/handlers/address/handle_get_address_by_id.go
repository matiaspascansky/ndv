package address

import (
	"github.com/gorilla/mux"
	"ndv/usecases/address"
	"ndv/web/handlers"
	"net/http"
	"strconv"
)

type GetAddressByIdHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type getAddressByIdHandler struct {
	getAddressByID address.GetAddressByID
}

func NewGetAddressByIDHandler(getAddressByID address.GetAddressByID) *getAddressByIdHandler {
	return &getAddressByIdHandler{getAddressByID: getAddressByID}
}

func (g getAddressByIdHandler) Handle(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	addressID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		handlers.BadRequest(w, "BAD_REQUEST", "id.invalid")
		return
	}
	result, err := g.getAddressByID.Execute(addressID)

	if err != nil {
		if err == address.ErrNotFound {
			handlers.NotFound(w, "address not found with id: ", string(addressID))
			return
		}
		handlers.InternalServerError(w, err)
		return
	}

	handlers.OK(w, result)
	return
}
