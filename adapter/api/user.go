package api

import (
	"encoding/json"
	"net/http"
	"rest_api_go/usecase"
)

type UserApi struct {
	uc usecase.UserUsecase
}

func NewUserApi(uc usecase.UserUsecase) *UserApi {
	return &UserApi{uc: uc}
}

func (a *UserApi) Fetch(w http.ResponseWriter, r *http.Request) {
	var id int64 = 1
	u, err := a.uc.Fetch(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(u)
}
