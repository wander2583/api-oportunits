package handler

import "fmt"

func errParamIsRequire(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Role == "" {
		return errParamIsRequire("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequire("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequire("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequire("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequire("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequire("salary", "int64")
	}
	return nil
}
