package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

// swagger:parameters HealthCheckParam
type HealthCheckParam struct {
}

// swagger:response HealthCheckResp
type HealthCheckResp struct {
	// HealthCheck Response
	// in: body
	Body struct {
		// Example: 200
		Status int `json:"status"`
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /healthz HealthCheckParam
	//
	// health check url
	//
	// Responses:
	//  200: HealthCheckResp
	code := http.StatusOK
	w.WriteHeader(code)
	res := HealthCheckResp{}
	res.Body.Status = http.StatusOK
	ret, _ := json.Marshal(res.Body)
	_, err := fmt.Fprintln(w, string(ret))
	if err != nil {
		logrus.Error(fmt.Sprintf("Error: %v\n", err))
	}
}
