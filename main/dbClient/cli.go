package main

import (
	"flag"
)

type Options struct {
	port     string
	logto    string
	loglevel string
	ip       string
}

func parseArgs() *Options {
	ip := flag.String("ip", "10.176.34.130", " server addr")
	port := flag.String("port", ":1024", "server port")
	logto := flag.String("log", "stdout", "Write log messages to this file. 'stdout' and 'none' have special meanings")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
	flag.Parse()

	return &Options{
		port:     *port,
		logto:    *logto,
		loglevel: *loglevel,
		ip:       *ip,
	}
}
