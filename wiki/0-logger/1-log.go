package main

import "log"

func main() {
	log.Print("INFO: message")
	log.Panic("Panic: message + stack trace")
	log.Fatal("FATAL: message")
}
