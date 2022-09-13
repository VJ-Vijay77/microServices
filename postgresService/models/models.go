package models

var Schema = `CREATE TABLE IF NOT EXISTS users (
	id serial primary key,
	name text,
	email text,
	password text,
	age integer);`

type Users struct {
	Id int `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
