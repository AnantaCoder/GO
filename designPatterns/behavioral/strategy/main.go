package main

import "fmt"

// idbconnect is an interface and its calling connect method for cll inherited structures
type IDBconnection interface {
	Connect()
}

type SQLconnection struct {
	connectionString string
}
func (c SQLconnection) Connect() {
	fmt.Println("SQL"+c.connectionString)
}
type OracleConnection struct {
	connectionString string
}

func (c OracleConnection) Connect(){
	fmt.Println("Oracle"+ c.connectionString)
}
type MongoDBConnection struct{
	connectionString string
}

func (c MongoDBConnection) Connect(){
	fmt.Println("MongoDB"+ c.connectionString)
}
/*
This is like a manager 👨‍💼.

He doesn’t care how you connect.

He just says: “You have a Connect() button? Okay, I’ll call it!”
*/
type DBConnection struct{
	db IDBconnection
}


func (c DBConnection) DBConnect(){
	c.db.Connect()
}
func main() {
	SQLconnection := SQLconnection{connectionString: " Connection is connected "}
	con := DBConnection{db: SQLconnection}
	con.DBConnect()


	OracleConnection := OracleConnection{connectionString: " is connected "}
	con = DBConnection{db:OracleConnection}
	con.DBConnect()


	MongoDBConnection:= MongoDBConnection{connectionString: " is connected "}
	con = DBConnection{db: MongoDBConnection}
	con.DBConnect()
}

/*

The Strategy pattern basically allows an object to select and execute at runtime without knowing how to implement the business logic it wants, instead of using different types of algorithms that can be used to perform each operation.
*/