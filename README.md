# DummyApp
a Dummy Web Application used to experiment CI/CD concepts. 

# Motivation
I needed a custom application with which to experiment with. 

# Build Status
TODO

# Features
* the applications listens on a given port for the following http requests:
  * /select performs a select on the read replica database
  * /insert insert a row in the master database
  * /update updates a row in the master database
  * /delete deletes a row in the master database
* uses 2 database connections (master and read replica)
* configurable through a settings file
* writes events in a log file

# Prerequisites
* The application requires a MySQL database that can be created using create_mysql_database_schema.sql as follows (run the command in the same directory as the create_mysql_database_schema.sql file):
```
mysql -h YOUR_DB_HOST -u YOUR_DB_ADMIN_USER -p < create_mysql_database_schema.sql
```

# How to get and build the source code
```
$ git clone https://github.com/dpaquette77/DummyApp.git
$ cd DummyApp
$ go build DummyApp.go
```

# Configuration
The application uses a json configuration file passed to the application using -c flag. See DummyApp.json.sample for the fairly self-explanatory configuration file. Things to configure are:
* database connection details for reads 
* database connection details for writes
* file name of the application log
* port on which the application listens

# Tests
TODO: Provide details on inculded tests and how to run them

# How to use?
```
$ DummyApp -c ./DummyApp.json &
$ curl localhost:8888/insert
inserted id: 1234
$ curl localhost:8888/select?id=1234
id: 1234, lastUpdateTime: 2020-10-28 18:00:23
```

