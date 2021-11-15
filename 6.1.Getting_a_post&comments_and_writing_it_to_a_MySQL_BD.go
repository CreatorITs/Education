package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
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

var SqlDriverName = "mysql"
var DB_Username = "root"
var DB_Password = "root1234"
var DB_Host = "127.0.0.1"
var DB_Port = "3306"
var HostUserPass = DB_Username + ":" + DB_Password + "@tcp" + "(" + DB_Host + ":" + DB_Port + ")/"

func createDataBase(nameDB string) {
	db, err := sql.Open(SqlDriverName, HostUserPass)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + nameDB + ";")
	if err != nil {
		panic(err)
	}
}

func createTable(nameDB, nameTable string, columnNumber int, ColName1, ColName2, ColName3, ColName4, ColName5 string) {
	var postsTable = "CREATE TABLE " + nameTable + "(" + ColName1 + " varchar(255)" + "," + ColName2 + " varchar(255)" + "," + ColName3 + " varchar(255)" + "," + ColName4 + " varchar(255)" + ")" + ";"
	var commentsTable = "CREATE TABLE " + nameTable + "(" + ColName1 + " varchar(255)" + "," + ColName2 + " varchar(255)" + "," + ColName3 + " varchar(255)" + "," + ColName4 + " varchar(255)" + "," + ColName5 + " varchar(255)" + ")" + ";"

	db, err := sql.Open(SqlDriverName, HostUserPass)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + nameDB + ";")
	if err != nil {
		panic(err)
	}

	switch columnNumber  {
	case 4:
		_, err = db.Exec(postsTable)
		if err != nil {
			panic(err)
		}
	case 5:
		_, err = db.Exec(commentsTable)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("Set the number of columns (columnNumber) in the table with posts (TablePosts) to 4 and for the table with comments (TableComments) to 5")
		panic(err)
	}
}

func insertTable(nameDB, nameTable string, columnNumber int, UserId, PostId, Id, Title, Name, Email, Body interface{}) {
	postsTable := "INSERT INTO " + nameTable + "(UserId, Id, Title, Body) VALUES (?, ?, ?, ?);"
	commentsTable := "INSERT INTO " + nameTable + "(PostId, Id, Name, Email, Body) VALUES (?, ?, ?, ?, ?);"

	db, err := sql.Open(SqlDriverName, HostUserPass)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + nameDB + ";")
	if err != nil {
		panic(err)
	}

	switch columnNumber  {
	case 4:
		_, err = db.Exec(postsTable, UserId, Id, Title, Body)
		if err != nil {
			panic(err)
		}
	case 5:
		_, err = db.Exec(commentsTable, PostId, Id, Name, Email, Body)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("Set the number of columns (columnNumber) in the table with posts (TablePosts) to 4 and for the table with comments (TableComments) to 5")
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

	//Create DataBase and Tables for Posts and Comments
	createDataBase("DB_UserID_7")
	createTable("DB_UserID_7", "TablePosts", 4, "UserId", "Id", "Title", "Body", "Empty")
	createTable("DB_UserID_7", "TableComments", 5, "PostId", "Id", "Name", "Email", "Body")

	// 	Posts struct
	var post Posts
	urlpost := "https://jsonplaceholder.typicode.com/posts?userId=7"
	jsonPlaceHolder(urlpost, &post)
	for p := 0; p < len(post); p++ {
		go insertTable("DB_UserID_7", "TablePosts", 4, post[p].UserId, 0, post[p].Id, post[p].Title, 0, 0, post[p].Body)
	}

	// Сomments struct
	var comment Сomments
	urlcomments := "https://jsonplaceholder.typicode.com/comments?postId=7"
	jsonPlaceHolder(urlcomments, &comment)
	for c := 0; c < len(comment); c++ {
		go insertTable("DB_UserID_7", "TableComments", 5, 0, comment[c].PostId, comment[c].Id, 0, comment[c].Name, comment[c].Email, comment[c].Body)
		time.Sleep(1001 * time.Microsecond)
	}
}
