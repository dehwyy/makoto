package middleware

import "github.com/dehwyy/Makoto/backend/distributor/logger"

type middleware struct {
	l logger.AppLogger
}

func New(l logger.AppLogger) *middleware {
	return &middleware{
		l: l,
	}
}
