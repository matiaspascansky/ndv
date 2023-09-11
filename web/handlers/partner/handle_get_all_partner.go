package partner

import (
	"ndv/usecases/partner"
	"ndv/web/handlers"
	"net/http"
	"strconv"
)

type GetAllPartnersHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type getAllPartnersHandler struct {
	getAllPartners   partner.GetAllPartners
	getPartnerByType partner.GetPartnerByType
}

func NewGetAllPartnersHandler(getAllPartners partner.GetAllPartners, getPartnerByType partner.GetPartnerByType) *getAllPartnersHandler {
	return &getAllPartnersHandler{
		getAllPartners:   getAllPartners,
		getPartnerByType: getPartnerByType,
	}
}

func (g getAllPartnersHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	partnerType := request.URL.Query().Get("partnerType")
	if partnerType == "" {
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
	p, _ := strconv.ParseInt(partnerType, 10, 64)

	partners, err := g.getPartnerByType.Execute(int(p))
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
