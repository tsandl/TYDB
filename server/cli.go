package server

import (
	"flag"
)

type Options struct {
	port     string
	DBPath   string
	logto    string
	loglevel string
}

func parseArgs() *Options {
	port := flag.String("port", ":1024", "")
	DBPath := flag.String("dbPath", "F:\\data\\storage\\db4", "db save path")
	logto := flag.String("log", "stdout", "Write log messages to this file. 'stdout' and 'none' have special meanings")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
	flag.Parse()

	return &Options{
		port:     *port,
		DBPath:   *DBPath,
		logto:    *logto,
		loglevel: *loglevel,
	}
}
