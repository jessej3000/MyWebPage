package main

import (
	"fmt"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

//  Definition      :     verifyUser function to verify if user exist and password is correct in the database
//  Returns         :     (bool),(int) returns true or false, and returns id of user
func verifyUser(username string, pwd string) (bool, int) {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	res, err := db.Start("SELECT * FROM user WHERE username='" + username + "' AND pwd=SHA1('" + pwd + "')")
	if err != nil {
		panic(err)
	}

	row, err := res.GetRow()
	if err != nil {
		panic(err)
	}

	if len(row) > 0 {
		// Return true as success with the id number
		return true, row.Int(0)
	}

	// Return false if does not exist and 0
	return false, 0
}

//  Definition      :     checkUser function to verify if user exist in the database
//  Returns         :     (bool),(int) returns true or false, and returns id of user
func checkUser(username string) (bool, int) {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	res, err := db.Start("SELECT * FROM user WHERE username='" + username + "'")
	if err != nil {
		panic(err)
	}

	row, err := res.GetRow()
	if err != nil {
		panic(err)
	}

	if len(row) > 0 {
		// Return true as success with the id number
		return true, row.Int(0)
	}

	// Return false if user does not exist and 0
	return false, 0
}

func registerUser(person user) bool {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	if ok, _ := checkUser(person.username); ok {
		return false
	}

	query := "INSERT INTO user("
	query = query + "username,"
	query = query + "pwd,"
	query = query + "email,"
	query = query + "fullname,"
	query = query + "address,"
	query = query + "telephone,"
	query = query + "longitude,"
	query = query + "latitude,"
	query = query + "googleacc) "
	query = query + "VALUES ('"
	query = query + person.username + "',SHA1('"
	query = query + person.password + "'),'"
	query = query + person.email + "','"
	query = query + person.fullname + "','"
	query = query + person.address + "','"
	query = query + person.telephone + "',"
	query = query + "0,"
	query = query + "0,'"
	query = query + person.googleacc + "')"
	fmt.Println(query)
	_, res, err := db.Query(query)
	if res == nil {
		//Do nothing
	}
	if err != nil {
		panic(err)
	}

	return true

	/*res, err := db.Start("SELECT * FROM user WHERE email='" + person.email + "')")
	if err != nil {
		panic(err)
	}

	row, err := res.GetRow()
	if err != nil {
		panic(err)
	}

	if len(row) > 0 {
		// Return true as success with the id number
		return true, row.Int(0)
	}

	// Return false if does not exist and 0
	return false, 0*/
}
