package models

type Response[T any] struct {
	Function string `json:"function"`

	Response T `json:"response"`
}

type ErrorResponse struct {
	Function string `json:"function"`

	Code int `json:"code"`

	Message string `json:"message"`
}

// Cars

type AddCarResponse struct {
	Car CarResponse `json:"car"`
}

type GetCarsByIdsResponse struct {
	Cars []CarResponse `json:"cars"`
}

type GetCarsByBrandResponse struct {
	Cars []CarResponse `json:"cars"`
}

// Engines

type AddEngineResponse struct {
	Engine EngineResponse `json:"engine"`
}

type GetEnginesByIdsResponse struct {
	Engines []EngineResponse `json:"engines"`
}

// Users

type AddUserResponse struct {
	User UserResponse `json:"user"`
}

type AddRegistrationToUserIdResponse struct {
	Registration RegistrationResponse `json:"registration"`
}

type GetUsersByIdsResponse struct {
	Users []UserResponse `json:"users"`
}

// Gateway

type GetCarsEnginesByBrandResponse struct {
	Engines []EngineResponse `json:"engines"`
}

type GetCarEngineByCarIdResponse struct {
	Engine EngineResponse `json:"engine"`
}

type GetEnginesResponse struct {
	Engines []EngineResponse `json:"engines"`
}

type GetUserAllCarsResponse struct {
	Registration []RegistrationResponse `json:"registration"`
}

type GetUserAllEnginesResponse struct {
	Engines []EngineResponse `json:"engines"`
}
