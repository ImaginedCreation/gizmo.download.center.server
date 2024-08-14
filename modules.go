package main

import (
	"github.com/go-playground/validator/v10"
)

type TRegister_P struct {
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type TLogin_P struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TGlobalConfig struct {
	PORT string `json:"port"`
}

func ParseValidatorError(err error) string {
	if err.Error() == "EOF" {
		return "please fill request parameters ."
	}
	if _, ok := err.(validator.ValidationErrors); ok {
		return "paramter validator failed ."
	}
	return err.Error()
}

type RESCODE int

const (
	RESCODEOK    RESCODE = 200
	RESCODEERROR RESCODE = 500
)

type RES[T any] struct {
	Code    RESCODE `json:"code"`
	Message string  `json:"message"`
	Data    T       `json:"data"`
}

func RES_OK_FN[T any](data T) RES[T] {
	return RES[T]{
		Code:    RESCODEOK,
		Message: "OK.",
		Data:    data,
	}
}

func RES_ERROR_FN(error_str string) RES[string] {
	return RES[string]{
		Code:    RESCODEERROR,
		Message: error_str,
		Data:    "",
	}
}
