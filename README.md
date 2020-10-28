# DummyApp
a Dummy Web Application used to experiment CI/CD concepts. 

# Motivation
I needed a custom application with which to experiment with. 

# Build Status
TODO

# Features
* the applications listens on a given port for the following http requests:
  * /select performs a select on the database read database
  * /insert insert a row in the master database
  * /update updates a row in the master database
  * /delete deletes a row in the master database
* uses 2 database connections (master and read replica)
* configurable through a settings file
* writes events in a log file

# Installation
TODO: Provide installation steps

# Tests
TODO: Provide details on inculded tests and how to run them

# How to use?
TODO: Provide example on how to use the deployed version of the app, basically this could be the curl commands to trigger each one of the supported operation (select, insert, update, delete). If a client test script is eventually added, an example of using the script should be included here


