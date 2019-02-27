package models

import (
	"gopkg.in/mgo.v2"
)

// Config will hold configurable pieces of the application
// this can be defined either by enviroment variables or
// command line arguments when running the app.
// --
// Env - Environment the app is running. (prod, dev, test)
// Port - The port in which the app will run. (:8081, :9000, :8085)
// Version - Specifies the version of the app
// DatabaseName - Specifies the name of the database the app will connect to.
// ConnectionString - Points to the database the app will be connecting to.
type Config struct {
	Env              string
	Port             int
	Version          string
	DatabaseName     string
	ConnectionString string
}

type DatabaseConfig struct {
	DB               *mgo.Database
	Collection       string
	DatabaseName     string
	ConnectionString string
}
