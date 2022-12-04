package infra

import "log"

type SimpleLogger struct{}

func (a SimpleLogger) AddLog(stm string) {
	log.Printf("Log entry added %v...", stm)

}
