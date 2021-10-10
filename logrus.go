package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func logrus1() {
	fmt.Println("========== logrus1() ==========")

	log.SetFormatter(&log.JSONFormatter{})

	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
}

func logrus2() {
	fmt.Println("========== logrus2() ==========")

	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Info("Something happened")
}

func logrus3() {
	fmt.Println("========== logrus3() ==========")

	log.SetLevel(log.TraceLevel)

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}
