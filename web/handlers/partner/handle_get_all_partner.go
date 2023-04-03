package partner

import (
	"ndv/usecases/partner"
	"ndv/web/handlers"
	"net/http"
)

type GetAllPartnersHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type getAllPartnersHandler struct {
	getAllPartners partner.GetAllPartners
}

func NewGetAllPartnersHandler(getAllPartners partner.GetAllPartners) *getAllPartnersHandler {
	return &getAllPartnersHandler{
		getAllPartners: getAllPartners,
	}
}

func (g getAllPartnersHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	partners, err := g.getAllPartners.Execute()
	if err != nil {
		if err == partner.ErrNotFound {
			handlers.NotFound(writer, "NOT_FOUND")
			return
		}
		handlers.InternalServerError(writer, err)
		return
	}
	handlers.OK(writer, partners)

}
