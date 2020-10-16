package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "time"

import "database/sql"
import "encoding/json"
import "net/http"
import _ "github.com/go-sql-driver/mysql"

var appConfig DummyAppConfig
var configFilePath = "/Users/dp/go/DummyApp/DummyApp.json"
var readDb *sql.DB
var writeDb *sql.DB

// DummyAppConfig struct is the top level struct that reprensents
// the content of the json config file
type DummyAppConfig struct {
	ReadDbConfig  ReadMySQLConnectionConfig  `json:"read_db"`
	WtiteDbConfig WriteMySQLConnectionConfig `json:"write_db"`
	Logfile       string                     `json:"logfile"`
	Port          int                        `json:"port"`
}

// ReadMySQLConnectionConfig contains all elements of a read only mysql conn
type ReadMySQLConnectionConfig struct {
	Server   string `json:"server"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// WriteMySQLConnectionConfig contains all elements of a read only mysql conn
type WriteMySQLConnectionConfig struct {
	Server   string `json:"server"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		fmt.Printf("Error while attempting to read config file: %w", err)
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// convert jsonFile to a byte array
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error parsing config file: %w", err)
		log.Fatal(err)
	}

	// Unmarshal json to appConfig struct
	// TODO: add error checks
	json.Unmarshal(jsonBytes, &appConfig)

	// initialize the log file
	logFile, err := os.OpenFile(appConfig.Logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// attempt to connect to both read and write db
	readDb, err = sql.Open("mysql", "admin:0uYMW%rQ1ZJxQE972gDW@tcp(dummyappdatabase.cogjdqpkoljl.ca-central-1.rds.amazonaws.com:3306)/DummyApp")
	if err != nil {
		log.Fatal(err)
	}
	defer readDb.Close()

	readDb.SetConnMaxLifetime(time.Minute * 3)
	readDb.SetMaxOpenConns(10)
	readDb.SetMaxIdleConns(10)

	err = readDb.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connection to read database succeeded")
	writeDb, err = sql.Open("mysql", "admin:0uYMW%rQ1ZJxQE972gDW@tcp(dummyappdatabase.cogjdqpkoljl.ca-central-1.rds.amazonaws.com:3306)/DummyApp")
	if err != nil {
		log.Fatal(err)
	}
	defer writeDb.Close()

	writeDb.SetConnMaxLifetime(time.Minute * 3)
	writeDb.SetMaxOpenConns(10)
	writeDb.SetMaxIdleConns(10)

	err = writeDb.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connection to write database succeeded")
	// setup the http server stuff to handle requests to /
	http.HandleFunc("/", DummyHttpServer)
	http.HandleFunc("/insert", DummyHttpServerInsert)
	http.HandleFunc("/select", DummyHttpServerSelect)

	log.Printf("attempting to listen on port %s", appConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), nil)
}

func DummyHttpServer(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("DummyApp request path: %s", r.URL.Path[1:])
	log.Print(response)
	fmt.Fprintf(w, response)
}

func DummyHttpServerInsert(w http.ResponseWriter, r *http.Request) {
	// Prepare statement for insert
	stmtInsert, err := writeDb.Prepare("INSERT INTO test_writes(id) VALUES (NULL)")
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}
	defer stmtInsert.Close() // Close the statement when we leave main() / the program terminates

	result, err := stmtInsert.Exec()
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	fmt.Printf("id=%d", id)
	response := fmt.Sprintf("inserted id: %d", id)
	log.Print(response)
	fmt.Fprintf(w, response)
}

func DummyHttpServerSelect(w http.ResponseWriter, r *http.Request) {
	// Prepare statement for select
	id := -1
	stmtSelect, err := readDb.Prepare("SELECT id FROM test_writes WHERE id=?")
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}
	defer stmtSelect.Close() // Close the statement when we leave main() / the program terminates

	err = stmtSelect.QueryRow(5).Scan(&id)
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	log.Printf("retreived id=%d", id)

	response := fmt.Sprintf("inserted id: %d", 5)
	log.Print(response)
	fmt.Fprintf(w, response)
}
