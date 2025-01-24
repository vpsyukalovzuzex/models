package models

type Request[T any] struct {
	UserId string `json:"user_id" validate:"required"`

	Request T `json:"request"`
}

// Cars

type AddCarRequest struct {
	EngineId string `json:"engine_id" validate:"required"`

	Name  string `json:"name" validate:"required,gte=1,lte=100"`
	Brand string `json:"brand" validate:"required,gte=1,lte=100"`
}

type GetCarsByIdsRequest struct {
	Ids []string `json:"ids"`
}

type GetCarsByBrandRequest struct {
	Brand string `json:"brand"`
}

// Engines

type AddEngineRequest struct {
	Name  string `json:"name" validate:"required,gte=1,lte=100"`
	Power int    `json:"power" validate:"required,min=0,max=2000"`
}

type GetEnginesByIdsRequest struct {
	Ids []string `json:"ids"`
}

// Users

type AddUserRequest struct {
	Name    string `json:"name" validate:"required,gte=1,lte=100"`
	Surname string `json:"surname" validate:"required,gte=1,lte=100"`
}

type AddRegistrationToUserIdRequest struct {
	CarId  string `json:"car_id" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
	Number string `json:"number" validate:"required,gte=1,lte=100"`
}

type GetUsersByIdsRequest struct {
	Ids []string `json:"ids"`
}
