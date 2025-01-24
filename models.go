package models

import "time"

type Timestamp struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// Cars

type CarResponse struct {
	Id string `json:"id"`

	Name  string `json:"name"`
	Brand string `json:"brand"`

	Engine EngineResponse `json:"engine"`

	Timestamp `json:"timestamp"`
}

// Engines

type EngineResponse struct {
	Id string `json:"id"`

	Name  string `json:"name"`
	Power int    `json:"power"`

	Timestamp `json:"timestamp"`
}

// Esers

type UserResponse struct {
	Id string `json:"id"`

	Name    string `json:"name"`
	Surname string `json:"surname"`

	Registrations []RegistrationResponse `json:"registrations"`

	Timestamp `json:"timestamp"`
}

type RegistrationResponse struct {
	Id string `json:"id"`

	Number string `json:"number"`

	Car CarResponse `json:"car"`

	Timestamp `json:"timestamp"`
}
