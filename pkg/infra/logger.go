package infra

import "log"

type TestApacheLogger struct{}

func (a TestApacheLogger) AddLog(stm string) {
	log.Printf("Item added %v...", stm)

}
