package handlers

import (
	"dcswitch/internal/inbound/http/payloads"
	"dcswitch/internal/outbound/db"
	"dcswitch/internal/service"
	"dcswitch/pkg/mysql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetAllSwitchVersions(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/versions GetAllVersionsParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetAllVersionsResp
	//  500:
	repo := db.SwitchVersionDBRepo{}
	svs := service.SwitchVersionService{SwRepo: repo}
	vResp := payloads.GetAllVersionsResp{}
	versions, err := svs.GetAll()
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Versions = versions

	resp, err := json.Marshal(vResp.Body)
	if err != nil {
		setServerError(err, w)
		return
	}
	_, err = fmt.Fprintf(w, string(resp))
	if err != nil {
		setServerError(err, w)
		return
	}
	setGetSuccess(w)
}

func EditSwitchVersionName(w http.ResponseWriter, r *http.Request) {
	// swagger:route PATCH /switch/version/name/{id} EditVersionName
	//
	// 编辑切换版本名
	//
	// Responses:
	//  200:
	//  404:
	//  500:
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	sv := payloads.EditVersionName{}
	err = json.Unmarshal(body, &sv.Body)
	if err != nil {
		setReqBodyError(err, string(body), w)
		return
	}
	repo := db.SwitchVersionDBRepo{}
	svs := service.SwitchVersionService{SwRepo: repo}
	_, err = svs.EditName(id, sv.Body.Name)
	switch err.(type) {
	case mysql.NotFoundError:
		setNotFound(err, w)
		return
	}
	if err != nil {
		setServerError(err, w)
		return
	}
	setPutSuccess(w)
}
