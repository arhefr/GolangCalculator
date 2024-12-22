package calc

import (
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

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				var response responseERROR
				response.Error = "Internal server error"
				responseJSON, err := json.Marshal(response)
				if err != nil {
					return
				}
				http.Error(w, string(responseJSON), 422)
				return
			}
		}()
		next.ServeHTTP(w, r)
	}
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	var exp post

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &exp)
	if err != nil {
		return
	}

	res, err := Calc(exp.Value)
	if err != nil {
		var response responseERROR
		response.Error = "Expression is not valid"
		responseJSON, err := json.Marshal(response)
		if err != nil {
			return
		}
		http.Error(w, string(responseJSON), 422)

	} else {
		var response responseOK
		response.Solve = res
		responseJSON, err := json.Marshal(response)
		if err != nil {
			return
		}
		http.Error(w, string(responseJSON), 200)
	}

}

func StartServer() {
	handler := RPC(http.HandlerFunc(calcHandler))

	http.HandleFunc("/api/v1/calculate", handler)
	http.ListenAndServe(":8080", handler)
}
