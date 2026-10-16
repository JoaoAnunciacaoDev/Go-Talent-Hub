package handler

import (
	"fmt"
	"uuid"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Role        string    `json:"role"`
	Location    string    `json:"location"`
	IsRemote    *bool     `json:"remote"`
	IsActive    *bool     `json:"isActive"`
	IsSecret    *bool     `json:"isSecret"`
	Link        string    `json:"link"`
	Salary      int64     `json:"salary"`
}

type UpdateOpeningRequest struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Role        string    `json:"role"`
	Location    string    `json:"location"`
	IsRemote    *bool     `json:"remote"`
	IsActive    *bool     `json:"isActive"`
	IsSecret    *bool     `json:"isSecret"`
	Link        string    `json:"link"`
	Salary      int64     `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}

	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.IsRemote == nil {
		return errParamIsRequired("isRemote", "bool")
	}

	if r.IsActive == nil {
		return errParamIsRequired("isActive", "bool")
	}

	if r.IsSecret == nil {
		return errParamIsRequired("isSecret", "bool")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Title == "" || r.Description == "" || r.Company == "" || r.Role == "" || r.Location == "" || r.Link == "" || r.IsRemote == nil || r.IsActive == nil || r.IsSecret == nil || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be provided for update")
}
