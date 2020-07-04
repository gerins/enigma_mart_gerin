package message

import (
	"net/http"
)

type ResponeMessage struct {
	Message string
	Code    int
	Status  string
	Results interface{}
}

func Respone(msg string, code int, rslt interface{}) *ResponeMessage {
	return &ResponeMessage{
		Message: msg,
		Code:    code,
		Status:  http.StatusText(code),
		Results: rslt,
	}
}

/*
=== Penggunaan ===
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.Respone("Update Failed", http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message.Respone("Update Success", http.StatusOK, "ResultYangMauDitampilin"))
*/
