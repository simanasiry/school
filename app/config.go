package app

import "time"

type config struct {
	Http struct {
		Address                   string
		GracefullyShutdownTimeout time.Duration
	}

	Sql struct {
		Dialect          string
		Host             string
		Port             string
		Username         string
		Password         string
		DbName           string
		ConnectionString string
	}
}

func (c *config) Initialize() {
	if c.Http.GracefullyShutdownTimeout == 0 {
		c.Http.GracefullyShutdownTimeout = time.Second * 10
	}
}
