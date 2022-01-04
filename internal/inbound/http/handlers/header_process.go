package handlers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func setGetSuccess(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func setPutSuccess(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func setPostSuccess(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
}

func setDelSuccess(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func setServerError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err = fmt.Fprintf(w, "error: %v", err)
	if err != nil {
		logrus.Error(err)
	}
}

func setParamsError(err error, vars interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	_, err = fmt.Fprintf(w, "Error: %v, Unsupportted Params: %v \n", err, vars)
	if err != nil {
		logrus.Error(err)
	}
}

func setQueriesError(key string, value string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	_, err := fmt.Fprintf(w, "QueriesError Unsupportted Queries: {key: %s, value: %s}\n", key, value)
	if err != nil {
		logrus.Error(err)
	}
}

func setReqBodyError(err error, body string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	_, err = fmt.Fprintf(w, "Error: %v, Error Request Body: %v \n", err, body)
	if err != nil {
		logrus.Error(err)
	}
}

func setNotFound(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}
