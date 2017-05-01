package main

import "net/http"

func main() {
	serveHTTP()
}

func serveHTTP() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/dashboard", handleDashboard)

	http.ListenAndServe("localhost:1000", nil)
}
