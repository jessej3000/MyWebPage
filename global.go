package main

const (
	// DBUser MySQL database username
	DBUser = "root" //"test" //
	// DBPassword MySQL database password
	DBPassword = "root" //"test" //
	// DBName database name
	DBName = "apidb"
	// DBHost MySQL host server
	DBHost = "127.0.0.1"
	// DBPort MySQL port
	DBPort = "8889"
)

type result struct {
	message string
}

type user struct {
	username  string
	password  string
	email     string
	fullname  string
	address   string
	telephone string
	longitude float64
	latitude  float64
	googleacc string
}

// MyID global id of the user will later be set during login
var MyID = 0
