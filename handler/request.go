package handler

import (
	"fmt"

	"github.com/google/uuid"
)

type CreateAppointmentRequest struct {
	ID          uuid.UUID `json:"id"`
	PatientName string    `json:"patient_name"`
	DoctorName  string    `json:"doctor_name"`
	Date        string    `json:"date"`
	Confirmed   *bool     `json:"confirmed"`
	Description string    `json:"description"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("Param: %s (type: %s) is required", name, typ)
}

func (r *CreateAppointmentRequest) Validate() error {
	if r.PatientName == "" {
		return errParamIsRequired("patient_name", "string")
	}
	if r.DoctorName == "" {
		return errParamIsRequired("doctor_name", "string")
	}
	if r.Date == "" {
		return errParamIsRequired("date", "string")
	}
	if r.Confirmed == nil {
		return errParamIsRequired("confirmed", "bool")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}

type UpdateAppointmentRequest struct {
	ID          uuid.UUID `json:"id"`
	PatientName string    `json:"patient_name"`
	DoctorName  string    `json:"doctor_name"`
	Date        string    `json:"date"`
	Confirmed   *bool     `json:"confirmed"`
	Description string    `json:"description"`
}

func (r *UpdateAppointmentRequest) Validate() error {
	if r.PatientName != "" || r.DoctorName != "" || r.Date != "" || r.Confirmed != nil || r.Description != "" {
		return nil
	}
	return fmt.Errorf("params: at least one valid field must be provided")
}
