package models

type Basemessage[T any] struct{
	Success bool
	Data []T
	Message string

}

type Basemessageone[T any] struct{
	Success bool
	Data T
	Message string

}

type Basemessageerror struct{
	Success bool
	Error string
}