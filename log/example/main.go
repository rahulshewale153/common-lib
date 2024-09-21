package main

import "github.com/rahulshewale153/common-lib/log"

func main() {
	// Directly use global logger
	// Defaul level is log.ERROR
	// log.SetLogLevel(log.INFO)
	log.SetOutputFile("log.txt")
	log.Info("This info log")

	// Create a new logger
	l := log.NewLogger()
	l.SetLogLevel(log.INFO)
	l.SetOutputFile("log.txt")
	l.Warn("This is warning")

	l.Critical("This is critical")
	l.Fatal("This is fatal")

}
