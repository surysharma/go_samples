package infra

import "log"

type SimpleLogger struct{}

func (a SimpleLogger) AddLog(stm string) {
	log.Printf("Item added %v...", stm)

}
