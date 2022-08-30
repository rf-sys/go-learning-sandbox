package log

import "github.com/rs/zerolog/log"

type Logger interface {
	Fatal(err error, msg string)
	Error(err error, msg string)
	Debug(msg string)
	Info(msg string)
}

type ZeroLogger struct {
}

func NewZeroLogger() Logger {
	return ZeroLogger{}
}

func (logger ZeroLogger) Fatal(err error, msg string) {
	log.Fatal().Err(err).Msg(msg)
}

func (logger ZeroLogger) Error(err error, msg string) {
	log.Err(err).Msg(msg)
}

func (logger ZeroLogger) Debug(msg string) {
	log.Debug().Msg(msg)
}

func (logger ZeroLogger) Info(msg string) {
	log.Info().Msg(msg)
}
