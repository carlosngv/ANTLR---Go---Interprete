package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Input struct {
	Data string `json:"data"`
}

type Response struct {
	Status int `json:"status"`
	Msg string `json:"msg"`
}

func ProcessData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	var input Input
	var response Response

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": http.StatusBadRequest, "msg": "Error, revise los datos ingresados."})
		return
	}

	fmt.Println(input.Data)

	w.WriteHeader(http.StatusOK)

	response.Status = http.StatusOK
	response.Msg = "Entrada procesada con exito."

	jsonResponse, err := json.Marshal(response)

	if err != nil { return }

	w.Write(jsonResponse)

}
