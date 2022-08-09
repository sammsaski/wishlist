/**
* This is our application entry point. It will be responsible for reading environment
* variables, setting up the database as well as starting and stopping the API server.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	// "encoding/json"
	// "github.com/gorilla/mux"
	"github.com/sammsaski/wishlist-golang/backend/db"
	"github.com/sammsaski/wishlist-golang/backend/handler"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Homepage!")
// 	fmt.Println("Endpoint hit: homePage")
// }

// func returnItem(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint hit: returnItem")

// 	// get the id requested
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	// Create the DB query for the entry with id=key
// 	dbQuery := fmt.Sprintf("SELECT * FROM wishlist WHERE id = %d", key)
// 	// TODO: Implement a handler that holds a DB reference, so that we can make a query
// 	fmt.Println(dbQuery)
// }

// func returnAllItems(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint hit: returnAllItems")

// 	// TODO: Make a query to our database for all entries in the wishlist db.
// 	items := 0

// 	json.NewEncoder(w).Encode(items)
// }

// func handleRequests() {
// 	// http.HandleFunc("/", homePage)
// 	// http.HandleFunc("/items", returnAllItems)
// 	// log.Fatal(http.ListenAndServe(":10000", nil))
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/items", returnAllItems)
// 	log.Fatal(":10000", myRouter)
// }

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}

func main()  {
	// Establish DB connection
	addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}
	
	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := db.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database is it this error: %v", err)
	}
	defer database.Conn.Close()

	httpHandler := handler.NewHandler(database)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")

	/**
	* WHEN QUERYING, DON'T USE A SPRINTF() TO BUILD THE QUERY. THIS INTRODUCES
	* AN SQL INJECTION RISK.
	*/
}

