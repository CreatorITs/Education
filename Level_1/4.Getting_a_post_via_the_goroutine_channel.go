package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func printPosts(postch chan interface{}) {
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("UserID:", <-postch, "\nId:", <-postch, "\nTitle:", <-postch, "\nBody:", <-postch)
}

func main() {
	var post Posts
	postch := make(chan interface{}, 4)
	for p := 0; p < 100; p++ {
		go func() {
			jsonPlaceHolder("https://jsonplaceholder.typicode.com/posts", &post)
			postch <- post[p].UserId
			postch <- post[p].Id
			postch <- post[p].Title
			postch <- post[p].Body
		}()
		printPosts(postch)
	}
}