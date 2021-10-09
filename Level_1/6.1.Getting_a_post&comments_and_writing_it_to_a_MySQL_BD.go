package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Posts []struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Сomments []struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func createDataBase(nameDB string) {
	db, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_,err = db.Exec("CREATE DATABASE "+nameDB)
	if err != nil {
		panic(err)
	}
}

func createTablePosts(nameDB, nameTable string) {
	db, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_,err = db.Exec("USE "+nameDB)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE "+nameTable+"(UserId integer, Id integer, Title varchar(255), Body varchar(255));")
	if err != nil {
		panic(err)
	}
}

func createTableComments(nameDB, nameTable string) {
	db, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_,err = db.Exec("USE "+nameDB)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE "+nameTable+"(PostId integer, Id integer, Name varchar(255), Email varchar(255),Body varchar(255));")
	if err != nil {
		panic(err)
	}
}

func insertTablePosts(nameDB, nameTable string, UserId, Id, Title, Body interface{}) {
	db, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_,err = db.Exec("USE "+nameDB)
	if err != nil {
		panic(err)
	}

	ptable := "INSERT INTO "+nameTable+"(UserId, Id, Title, Body) VALUES (?, ?, ?, ?);"
	_, err = db.Exec(ptable, UserId, Id, Title, Body)
	if err != nil {
		panic(err)
	}
}

func insertTableComments(nameDB, nameTable string, PostId, Id, Name, Email, Body interface{}) {
	db, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_,err = db.Exec("USE "+nameDB)
	if err != nil {
		panic(err)
	}

	ctable := "INSERT INTO "+nameTable+"(PostId, Id, Name, Email, Body) VALUES (?, ?, ?, ?, ?);"
	_, err = db.Exec(ctable, PostId, Id, Name, Email, Body)
	if err != nil {
		panic(err)
	}
}

func jsonPlaceHolder(url string, webstruct interface{}) error {
	urlstruct, err := http.Get(url)
	if err != nil {
		return err
	}
	defer urlstruct.Body.Close()
	return json.NewDecoder(urlstruct.Body).Decode(webstruct)
}

func main() {

//Create DataBase and Tables
	createDataBase("DB_UserID_7")
	createTablePosts("DB_UserID_7", "TablePosts")
	createTableComments("DB_UserID_7", "TableComments")


// Posts struct
	var post Posts
	urlpost := "https://jsonplaceholder.typicode.com/posts?userId=7"
	jsonPlaceHolder(urlpost, &post)
	for p := 0; p < len(post); p++ {
		go insertTablePosts("DB_UserID_7", "TablePosts", post[p].UserId, post[p].Id, post[p].Title, post[p].Body)
	}

// Сomments struct
	var comment Сomments
	urlcomments := "https://jsonplaceholder.typicode.com/comments?postId=7"
	jsonPlaceHolder(urlcomments, &comment)
	for c := 0; c < len(comment); c++ {
		go insertTableComments("DB_UserID_7", "TableComments", comment[c].PostId, comment[c].Id, comment[c].Name, comment[c].Email, comment[c].Body)
	}
}
