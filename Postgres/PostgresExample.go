package main

import (
	"database/sql"
	d "database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // underscore denotes that we only import this package for it's "side effects". TODO: Better understand this concept.
)

type goUser struct {
	FirstName string
	LastName  string
	Age       int
}

const (
	host     = "{YOUR DB HOSTNAME HERE}"
	port     = 6432 // none standard port for my environment
	user     = "postgres"
	password = "postgres"
	dbname   = "Go_Example"
)

func main() {
	// Use the const values defined above to build a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Call into the sql package to connect to the databse
	db, err := sql.Open("postgres", psqlInfo)

	// Check if the err value was set, if so PANIC
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
		panic(err)
	}

	// Set the db.close method to be defered.
	// This means once the main() method completes, db.close() will be invoked
	defer db.Close()

	// Next we need to call the db.ping() method. See the below snip-it from the documentation as to why this is critical.

	// "It is vitally important that you call the Ping() method becuase the sql.Open() function call does not ever create a connection to the database. Instead, it simply validates the arguments provided."
	// "By calling db.Ping() we force our code to actually open up a connection to the database which will validate whether or not our connection string was 100% correct."
	err = db.Ping() // Notice we simply reuse the error variable declared above. Denoted by the "=" instead of the ":="
	if err != nil {
		log.Fatal("Failed to ping a DB connection: ", err)
		panic(err)
	}

	fmt.Println("Connected To Database Without Error.")

	user, err := getSingleGoUser(db, "Joseph")
	if err != nil {
		log.Fatal("Failed to retrieve goUser from postgres: ", err)
		panic(err)
	}

	fmt.Println(user)

	user, err = getSingleGoUser(db, "Cristina")
	if err != nil {
		log.Fatal("Failed to retrieve goUser from postgres: ", err)
		panic(err)
	}

	fmt.Println(user)

	var userToInsert = goUser{FirstName: "John", LastName: "Becker", Age: 26}
	err = insertGoUserToDatabase(db, userToInsert)
	if err != nil {
		log.Fatal("Failed to insert john into the DB: ", err)
		panic(err)
	}

	user, err = getSingleGoUser(db, "John")
	if err != nil {
		log.Fatal("Failed to retrieve goUser from postgres: ", err)
		panic(err)
	}

	fmt.Println(user)

	var allUsers = make([]goUser, 0)
	allUsers, err = getListOfGoUsers(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("********* Start Output From Multi-Row Query")
	for i := 0; i < len(allUsers); i++ {
		fmt.Println(allUsers[i])
	}
	fmt.Println("********* End Output From Multi-Row Query")

	// Need to delete all the johns, since we are hard coded to insert one per run of this app.
	err = deleteUsersByFirstName(db, "John")
	if err != nil {
		panic(err)
	}

	// Now get all users again to verify we don't have john anymore.
	allUsers, err = getListOfGoUsers(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("********* Start Output From Multi-Row Query ---- John should be gone now")
	for i := 0; i < len(allUsers); i++ {
		fmt.Println(allUsers[i])
	}
	fmt.Println("********* End Output From Multi-Row Query ---- John should be gone now")
}

func getSingleGoUser(db *d.DB, firstName string) (user goUser, err error) {
	var sql = fmt.Sprintf("SELECT * FROM public.go_users WHERE first_name = '%s'", firstName)

	err = db.QueryRow(sql).Scan(&user.FirstName, &user.LastName, &user.Age)

	if err != nil {
		log.Fatal("Failed to retrieve user from database.", err)
		return
	}

	return
}

func getListOfGoUsers(db *d.DB) (users []goUser, err error) {
	var sql = "SELECT * FROM public.go_users"

	result, err := db.Query(sql)

	for result.Next() {
		var user = goUser{}
		err := result.Scan(&user.FirstName, &user.LastName, &user.Age)

		if err != nil {
			log.Fatal("failed to get data from multi-row result set.")
			panic(err)
		}

		users = append(users, user)
	}

	return
}

func insertGoUserToDatabase(db *d.DB, user goUser) (err error) {
	var sql = "INSERT INTO go_users(first_name, last_name, age) VALUES ($1, $2, $3)"

	_, err = db.Exec(sql, user.FirstName, user.LastName, user.Age)

	if err != nil {
		log.Fatal("Failed to insert user to database.", err)
	}

	return
}

func deleteUsersByFirstName(db *d.DB, firstName string) (err error) {
	var sql = "DELETE FROM public.go_users WHERE first_name = $1"

	_, err = db.Exec(sql, firstName)
	if err != nil {
		log.Fatal("Failed to delete users from the database NAME = " + firstName)
	}

	return
}
