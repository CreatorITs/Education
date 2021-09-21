package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Posts []struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func jsonPlaceHolder(url string, webpost interface{}) error {
	urlpost, err := http.Get(url)
	if err != nil {
		return err
	}
	defer urlpost.Body.Close()
	return json.NewDecoder(urlpost.Body).Decode(webpost)
}

func main() {
	var post Posts
	postch := make(chan interface{}, 4)
	var jsondir string = "./storage/posts/"
	jsonfile := []string{
		jsondir + "1.txt",
		jsondir + "2.txt",
		jsondir + "3.txt",
		jsondir + "4.txt",
		jsondir + "5.txt"}

	for p := 0; p < 5; p++ {
		go func() {
			jsonPlaceHolder("https://jsonplaceholder.typicode.com/posts", &post)
			postch <- post[p].UserId
			postch <- post[p].Id
			postch <- post[p].Title
			postch <- post[p].Body
		}()

		file, err := os.Create(jsonfile[p])
		if err != nil {
			panic(err)
		}
		defer file.Close()

		postwriter := io.MultiWriter(os.Stdout, file)

		fmt.Fprintln(postwriter, "UserID:", <-postch, "\nID:", <-postch, "\nTitle:", <-postch, "\nBody:", <-postch)
		fmt.Println("---------------------------------------------------------------------")
	}
}
