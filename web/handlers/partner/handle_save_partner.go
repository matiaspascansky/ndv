package partner

import (
	"encoding/json"
	"io/ioutil"
	"ndv/usecases/partner"
	"ndv/web/handlers"
	"net/http"
)

type SavePartnerHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type savePartnerHandler struct {
	savePartner partner.SavePartner
}

func NewSavePartnerHandler(savePartner partner.SavePartner) *savePartnerHandler {
	return &savePartnerHandler{savePartner: savePartner}
}

func (s savePartnerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	model := partner.SavePartnerModel{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handlers.InternalServerError(w, err)
	}
	err = json.Unmarshal(data, &model)
	if err != nil {
		handlers.BadRequest(w, "INVALID_JSON", err.Error())
	}
	_, err = s.savePartner.Execute(&model)
	if err != nil {
		handlers.InternalServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
