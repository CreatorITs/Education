package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
var HostUserPass = DB_Username + ":" + DB_Password + "@tcp" + "(" + DB_Host + ":" + DB_Port + ")/" + SqlDriverName + "?" + "parseTime=true&loc=Local"

func createDataBase(nameDB string) {
	db, err := gorm.Open(mysql.Open(HostUserPass), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.Exec("CREATE DATABASE " + nameDB + ";")
}

func createTable(nameDB, nameTable string, columnNumber int, ColName1, ColName2, ColName3, ColName4, ColName5 string) {
	var postsTable = "CREATE TABLE " + nameTable + "(" + ColName1 + " varchar(255)" + "," + ColName2 + " varchar(255)" + "," + ColName3 + " varchar(255)" + "," + ColName4 + " varchar(255)" + ")" + ";"
	var commentsTable = "CREATE TABLE " + nameTable + "(" + ColName1 + " varchar(255)" + "," + ColName2 + " varchar(255)" + "," + ColName3 + " varchar(255)" + "," + ColName4 + " varchar(255)" + "," + ColName5 + " varchar(255)" + ")" + ";"

	db, err := gorm.Open(mysql.Open(HostUserPass), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.Exec("USE " + nameDB + ";")

	switch columnNumber  {
	case 4:
		_ = db.Exec(postsTable)
	case 5:
		_ = db.Exec(commentsTable)
	default:
		fmt.Println("Set the number of columns (columnNumber) in the table with posts (TablePosts) to 4 and for the table with comments (TableComments) to 5")
		panic(err)
	}
}

func insertTable(nameDB, nameTable string, columnNumber int, UserId, PostId, Id, Title, Name, Email, Body interface{}) {
	postsTable := "INSERT INTO " + nameTable + "(UserId, Id, Title, Body) VALUES (?, ?, ?, ?);"
	commentsTable := "INSERT INTO " + nameTable + "(PostId, Id, Name, Email, Body) VALUES (?, ?, ?, ?, ?);"

	db, err := gorm.Open(mysql.Open(HostUserPass), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.Exec("USE " + nameDB + ";")

	switch columnNumber  {
	case 4:
		_ = db.Exec(postsTable, UserId, Id, Title, Body)
	case 5:
		_ = db.Exec(commentsTable, PostId, Id, Name, Email, Body)
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
