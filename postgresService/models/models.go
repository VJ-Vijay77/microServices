package models

var Schema = `CREATE TABLE users (
	id serial primary key,
	name text,
	email text,
	password text,
	age integer);`