package models

import "fmt"

type AppError struct {
	Message string `json:"message"`
	Ctx     string `json:"context"`
}

func (a AppError) Error() string {
	return fmt.Sprintf("AppError: Message - %s Ctx - %s", a.Message, a.Ctx)
}
