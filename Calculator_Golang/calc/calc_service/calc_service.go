package calc_service

import (
	c "Calculator_Golang/calc"
	"encoding/json"
	"io"
	"net/http"
)

type post struct {
	Value string `json:"expression"`
}

type responseOK struct {
	Solve float64 `json:"result"`
}

type responseERROR struct {
	Error string `json:"error"`
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	var exp post
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "{'error':'Internal server error'}", 500)
		return
	}
	err = json.Unmarshal(body, &exp)
	if err != nil {
		http.Error(w, "{'error':'Internal server error'}", 500)
		return
	}

	res, err := c.Calc(exp.Value)
	if err != nil {
		http.Error(w, "{'error':'Expression is not valid'}", 422)
		return
	}
	var response responseOK
	response.Solve = res
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "{'error':'Internal server error'}", 500)
		return
	}
	http.Error(w, string(responseJSON), 200)

}

func StartServer() {
	handler := http.HandlerFunc(calcHandler)

	http.HandleFunc("/api/v1/calculate", handler)
	http.ListenAndServe(":8080", handler)
}
