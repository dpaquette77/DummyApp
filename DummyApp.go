package main

import "fmt"
import "log"
import "os"
import "net/http"
import "gopkg.in/ini.v1"

var cfg *ini.File

func main() {
	var err error
	cfg, err = ini.Load("DummyApp.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		log.Fatal(err)
	}

	// Print configuration parameters to stdout
	fmt.Println(cfg.Section("read_db").Key("server").String())
	fmt.Println(cfg.Section("read_db").Key("database").String())
	fmt.Println(cfg.Section("read_db").Key("username").String())
	fmt.Println(cfg.Section("read_db").Key("password").String())

	fmt.Println(cfg.Section("write_db").Key("server").String())
	fmt.Println(cfg.Section("write_db").Key("database").String())
	fmt.Println(cfg.Section("write_db").Key("username").String())
	fmt.Println(cfg.Section("write_db").Key("password").String())

	fmt.Println(cfg.Section("DummyApp").Key("logfile").String())
	fmt.Println(cfg.Section("DummyApp").Key("response_prefix").String())
	fmt.Println(cfg.Section("DummyApp").Key("port").String())

	// initialize the log file
	file, err := os.OpenFile(cfg.Section("DummyApp").Key("logfile").String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	// setup the http server stuff to handle requests to /
	http.HandleFunc("/", DummyHttpServer)

	log.Printf("attempting to listen on port %s", cfg.Section("DummyApp").Key("port").String())
	http.ListenAndServe(":"+cfg.Section("DummyApp").Key("port").String(), nil)
}

func DummyHttpServer(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("%s request path: %s", cfg.Section("DummyApp").Key("response_prefix").String(), r.URL.Path[1:])
	log.Print(response)
	fmt.Fprintf(w, response)
}
