package main

import "fmt"
import "io/ioutil"
import "log"
import "os"

import "encoding/json"
import "net/http"
import "gopkg.in/ini.v1"

var cfg *ini.File

var appConfig DummyAppConfig
var configFilePath = "/Users/dp/go/DummyApp/DummyApp.json"

type DummyAppConfig struct {
	ReadDbConfig  ReadMySqlConnectionConfig  `json:"read_db"`
	WtiteDbConfig WriteMySqlConnectionConfig `json:"write_db"`
	Logfile       string                     `json:"logfile"`
	Port          int                        `json:"port"`
}

type ReadMySqlConnectionConfig struct {
	Server   string `json:"server"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type WriteMySqlConnectionConfig struct {
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

	// setup the http server stuff to handle requests to /
	http.HandleFunc("/", DummyHttpServer)

	log.Printf("attempting to listen on port %s", appConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), nil)
}

func DummyHttpServer(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("DummyApp request path: %s", r.URL.Path[1:])
	log.Print(response)
	fmt.Fprintf(w, response)
}
