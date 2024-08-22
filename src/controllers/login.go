package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/responses"
)

// FazerLogin utiliza o email e a senha do usuário para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/jsn", bytes.NewBuffer(usuario))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	var datasAuthentication models.DatasAuthentication
	if err := json.NewDecoder(response.Body).Decode(&datasAuthentication); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}