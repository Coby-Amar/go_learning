package handlers

import "github.com/coby-amar/go_learning/database"

type RegistrationJson struct {
	Username    string `json:"username" validate:"required,email"`
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phonenumber" validate:"required"`
	Password    string `json:"password" validate:"required,min=8,max=70"`
}

type LoginJson struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=70"`
}

type UserCreateReportWithEntries struct {
	Report  database.CreateReportParams          `json:"report" validate:"required"`
	Entries []database.CreateReportEntriesParams `json:"entries" validate:"required,min=1,max=20"`
}
