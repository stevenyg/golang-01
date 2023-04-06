package api

import (
	internalConfig "github.com/stevenyg/golang-01/internal/config"
)

type InternalConfigInterface interface {
	ApplyConfig(c internalConfig.Config)
	GetConfig() internalConfig.Config
}
