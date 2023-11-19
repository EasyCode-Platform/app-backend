package controller

import (
	"github.com/EasyCode-Platform/app-backend/src/storage"
	"github.com/EasyCode-Platform/app-backend/src/utils/tokenvalidator"
)

type Controller struct {
	Storage               *storage.Storage
	RequestTokenValidator *tokenvalidator.RequestTokenValidator
}

func NewControllerForBackend(storage *storage.Storage, validator *tokenvalidator.RequestTokenValidator) *Controller {
	return &Controller{
		Storage:               storage,
		RequestTokenValidator: validator,
	}
}

func NewControllerForBackendInternal(storage *storage.Storage, validator *tokenvalidator.RequestTokenValidator) *Controller {
	return &Controller{
		Storage:               storage,
		RequestTokenValidator: validator,
	}
}
