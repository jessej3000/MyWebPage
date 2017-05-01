package main

const (
	// DBUser MySQL database username
	DBUser = "root" //"test" //
	// DBPassword MySQL database password
	DBPassword = "" //"test" //
	// DBName database name
	DBName = "apidb"
	// DBHost MySQL host server
	DBHost = "127.0.0.1"
	// DBPort MySQL port
	DBPort = "3306"
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
