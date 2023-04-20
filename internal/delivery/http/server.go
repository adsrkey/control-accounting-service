package http

import (
	"control-accounting-service/internal/delivery/http/config"
	"errors"
)

func (de *Delivery) Start(config *config.Config) error {
	if config.IsEmpty() {
		return errors.New("config is empty")
	}
	return de.engine.Run()
}
