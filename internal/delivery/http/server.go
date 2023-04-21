package http

import (
	"control-accounting-service/internal/config"
	"fmt"
)

func (de *Delivery) Start(config *config.Config) error {
	return de.engine.Run(fmt.Sprintf("0.0.0.0:%d", config.Port))
}
