package payloads

import "dcswitch/internal/domain"

// swagger:parameters GetAllVersionsParam
type GetAllVersionsParam struct{}

// swagger:response GetAllVersionsResp
type GetAllVersionsResp struct {
	// in: body
	Body struct {
		Versions []domain.SwitchVersion `json:"versions"`
	}
}

// swagger:parameters EditVersionName
type EditVersionName struct {
	// Required: true
	// in: path
	Id int64 `json:"id"`
	// Required: true
	// in: body
	Body struct {
		Name string `json:"name"`
	}
}

// swagger:parameters ModuleDetailTask
type ModuleDetailTask struct {
	// Required: true
	// in: body
	// example: {"type": "start", "name": "DBASwitch1"}
	Body struct {
		Name string `json:"name"`
		// Required: true
		// pattern: ^(start|success|fail)$
		Type string `json:"type"` // start/success
	}
}

func (p ModuleDetailTask) CheckParam() bool {
	checkPass := false
	typeOptions := []string{"start", "success", "fail"}
	for _, v := range typeOptions {
		if p.Body.Type == v {
			checkPass = true
			return checkPass
		}
	}
	return checkPass
}
