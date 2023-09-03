package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	DoctorName  string
	PatientName string
	Date        string
	Confirmed   bool
	Description string
}

type AppointmentResponse struct {
	ID          uuid.UUID `json:"id"`
	DoctorName  string    `json:"doctor_name"`
	PatientName string    `json:"patient_name"`
	Date        string    `json:"date"`
	Confirmed   bool      `json:"confirmed"`
	Description string    `json:"description"`
}
