package core

import (
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

const (
	moduleName = "injection"
)

// Injector wires all the application dependencies. This is done at runtime.
//
// It allows an easier setup in main, as well as more private class definitions (to enforce as most consistency as
// possible in the architecture).
//
//Given that it reuses the same standard constructors, it also lives perfectly alongside the tests.
//
var Injector = dig.New(dig.DeferAcyclicVerification())

// CheckInjection is in charge to validate if invocation contains errors
func CheckInjection(err error, instanceName string) {

	if err == nil {
		return
	}

	log.WithFields(log.Fields{
		"module":   moduleName,
		"instance": instanceName,
	}).Error(err)

	panic(err)
}
