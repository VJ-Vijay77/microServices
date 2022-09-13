package models

var Schema = `CREATE TABLE users (
	id serial primary key,
	name text,
	email text,
	password text,
	age integer);`

type Users struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
