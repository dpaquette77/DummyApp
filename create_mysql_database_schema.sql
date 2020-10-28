CREATE database DummyApp;

USE DummyApp;

CREATE TABLE dummy_data (
     id MEDIUMINT NOT NULL AUTO_INCREMENT,
     last_update_time DATETIME DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY (id)
);

