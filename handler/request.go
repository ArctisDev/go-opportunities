package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

// Create Opening
type CreateOpeningRequest struct {
	Role        string `json:"role"`
	Company     string `json:"company"`
	Description string `json:"description"`
	Remote      *bool  `json:"remote"`
	Location    string `json:"location"`
	URL         string `json:"url"`
	Salary      int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" &&
		r.Company == "" &&
		r.Description == "" &&
		r.Location == "" &&
		r.URL == "" &&
		r.Salary <= 0 &&
		r.Remote == nil {
		return fmt.Errorf("Request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.URL == "" {
		return errParamIsRequired("url", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "*bool")
	}
	return nil

}

type UpdateOpeningRequest struct {
	Role        string `json:"role"`
	Company     string `json:"company"`
	Description string `json:"description"`
	Remote      *bool  `json:"remote"`
	Location    string `json:"location"`
	URL         string `json:"url"`
	Salary      int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, it must be valid
	if r.Role != "" ||
		r.Company != "" ||
		r.Description != "" ||
		r.Location != "" ||
		r.URL != "" ||
		r.Salary > 0 ||
		r.Remote == nil {

		return nil
	}
	return fmt.Errorf("At least one field must be provided for update")
}