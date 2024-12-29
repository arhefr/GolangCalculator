package handler_test

import (
	"bytes"
	app "calculator/internal/application"
	"calculator/internal/application/handler"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	for i, testCase := range testCases {
		w, r := httptest.NewRecorder(), httptest.NewRequest(testCase.requestMethod, Config["urlHandlerCalc"], bytes.NewBuffer([]byte(testCase.request)))
		handler.CalcHandler(w, r)

		if strings.ReplaceAll(w.Body.String(), "\n", "") != testCase.response {
			t.Fatalf("#%d: %s expected response %s, but got %s", i, testCase.name, testCase.response, strings.ReplaceAll(w.Body.String(), "\n", ""))
		} else if w.Result().StatusCode != testCase.responseStatus {
			t.Fatalf("#%d: %s expected status %d, but got %d", i, testCase.name, testCase.responseStatus, w.Result().StatusCode)
		}
	}
}

var Config = app.Config
var testCases = []struct {
	name string

	request       string
	requestMethod string

	response       string
	responseStatus int
}{
	{
		name: "correct request",

		request:       "{\"expression\":\"2+2\"}",
		requestMethod: "POST",

		response:       "{'result':'4'}",
		responseStatus: 200,
	},
	{
		name: "incorrect math expression",

		request:       "{\"expression\":\"2+(\"}",
		requestMethod: "POST",

		response:       "{'error':'Expression is not valid'}",
		responseStatus: 422,
	},
	{
		name: "incorrect JSON",

		request:       "\"expression\":\"2+2\"",
		requestMethod: "POST",

		response:       "{'error':'Internal server error'}",
		responseStatus: 500,
	},
}
