package partner

import (
	"github.com/gorilla/mux"
	"ndv/usecases/partner"
	"ndv/web/handlers"
	"net/http"
	"strconv"
)

type GetPartnerByIDHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type getPatnerByIDHandler struct {
	getPartnerByID partner.GetPartnerByID
}

func NewGetPartnerByIDHandler(getPartnerByID partner.GetPartnerByID) *getPatnerByIDHandler {
	return &getPatnerByIDHandler{
		getPartnerByID: getPartnerByID,
	}
}

func (g getPatnerByIDHandler) Handle(writer http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	partnerID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		handlers.BadRequest(writer, "BAD_REQUEST", "id.invalid")
		return
	}
	p, err := g.getPartnerByID.Execute(partnerID)
	if err != nil {
		if err == partner.ErrNotFound {
			handlers.NotFound(writer, "NOT_FOUND")
			return
		}
		handlers.InternalServerError(writer, err)
		return
	}
	handlers.OK(writer, p)

}
