package application

import (
	c "calculator/calc"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type request struct {
	Expression string `json:"expression"`
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "pong", 200)
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method is not allowed"}`, 405)
		return
	}

	var req request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "{'error':'Internal server error'}", 500)
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "{'error':'Internal server error'}", 500)
		return
	}

	res, err := c.Calc(req.Expression)
	if err != nil {
		http.Error(w, "{'error':'Expression is not valid'}", 422)
		return
	}

	http.Error(w, fmt.Sprintf("{'result':'%s'}", res), 200)
}
