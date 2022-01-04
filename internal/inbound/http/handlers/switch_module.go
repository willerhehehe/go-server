package handlers

import (
	"dcswitch/internal/inbound/http/payloads"
	"dcswitch/internal/outbound/db"
	"dcswitch/internal/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateModuleDetailTask(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /task/switch/module/detail ModuleDetailTask
	//
	// 模块明细 start/success/fail
	//
	// Responses:
	//  201:
	//  403:
	//  500:
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	md := payloads.ModuleDetailTask{}
	err := json.Unmarshal(body, &md.Body)
	if err != nil {
		setReqBodyError(err, string(body), w)
		return
	}
	checkPass := md.CheckParam()
	if !checkPass {
		setReqBodyError(err, string(body), w)
		return
	}
	// TODO: REPO未写具体逻辑
	repo := db.ModuleDetailDBRepo{}
	svc := service.SwitchModuleSvc{SmRepo: repo}
	err = svc.Start(md.Body.Name)
	if err != nil {
		setServerError(err, w)
		return
	}
	setPostSuccess(w)
}
