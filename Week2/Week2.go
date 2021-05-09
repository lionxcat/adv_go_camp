package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// connTo creates a connection to database by dns and source name.
// The returned DB has been veryfied by calling Ping().
func connTo(dns string, name string) (db *sql.DB, err error) {
	db, err = sql.Open(name, dns)
	if err != nil {
		db = nil
		return
	}

	err = db.Ping()
	if err != nil {
		db = nil
	}
	return
}

// global query string for testing purpose only
var queryString = "select User from user where Host=?;"

// queryUsersByHost returns slice of user names under given host.
// If no users found, a zero length of user slice returns.
// If a db error occors when querying, user slice will be nil.
// If an error returns, values of user slice is not reliable.
func queryUsersByHost(db *sql.DB, host string) ([]string, error) {
	rows, err := db.Query(queryString, host)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return make([]string, 0), nil
		default:
			return nil, err
		}
	}
	defer rows.Close()

	var users []string
	var aUser string
	for rows.Next() {
		err = rows.Scan(&aUser)
		if err != nil {
			break
		}
		users = append(users, aUser)
	}
	return users, err
}

// doQueryTest does the query test
func doQueryTest(dns string, host string) {
	const sourceName = "mysql"
	db, err := connTo(dns, sourceName)
	if err != nil {
		// unrecoverable error
		panic(err)
	}
	defer db.Close()

	users, err := queryUsersByHost(db, host)
	if err != nil {
		fmt.Printf("Error occurs when querying users from host \"%s\" :\r\n", host)
		fmt.Println(err)
		//os.Exit(2)
		return
	}

	numberOfUsers := len(users)
	if numberOfUsers < 1 {
		fmt.Printf("No users found under host \"%s\" !\r\n", host)
	} else {
		fmt.Printf("%d users found under host \"%s\" :\r\n", numberOfUsers, host)
		for _, user := range users {
			println(user)
		}
	}
}

func main() {
	// normal circumstance
	println("[Round 1]")
	doQueryTest("root:123456@tcp(127.0.0.1:3306)/mysql", "localhost")

	// no result, ErrNoRows circumstance
	println("\r\n\r\n[Round 2]")
	doQueryTest("root:123456@tcp(127.0.0.1:3306)/mysql", "127.0.0.1")

	// wrong query string
	queryString = "select User111 from user where Host=?;"
	println("\r\n\r\n[Round 3]")
	doQueryTest("root:123456@tcp(127.0.0.1:3306)/mysql", "127.0.0.1")

	// access denied
	println("\r\n\r\n[Round 4]")
	doQueryTest("lionxcat:123456@tcp(127.0.0.1:3306)/mysql", "localhost")
}
