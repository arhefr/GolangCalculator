package handler

import (
	calc "calculator/pkg/calc"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {

	var request struct {
		Expression string `json:"expression"`
	}

	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, "{'error':'Internal server error'}", http.StatusInternalServerError)
		return
	}

	res, err := calc.Calc(request.Expression)
	if err != nil {
		http.Error(w, "{'error':'Expression is not valid'}", http.StatusUnprocessableEntity)
		return
	}

	http.Error(w, fmt.Sprintf("{'result':'%s'}", res), http.StatusOK)
}
