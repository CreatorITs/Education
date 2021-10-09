package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Posts []struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var post Posts
	url, err := http.Get("https://jsonplaceholder.typicode.com/posts/") // Выдача GET по url
	if err != nil {
		panic(err)
	}

	urlbody, error := ioutil.ReadAll(url.Body) // Чтение данных из url
	defer url.Body.Close()
	if error != nil {
		fmt.Println("ReadAll url error")
	}

	err = json.Unmarshal(urlbody, &post) // Декодирование JSON данных

	fmt.Println("------------------------------------------------------") // Вывод данных постов №98-100
	fmt.Printf("UserID: %+v\nID: %+v\nTitle: %+v\nBody: %+v\n", post[97].UserId, post[97].Id, post[97].Title, post[97].Body)
	fmt.Println("------------------------------------------------------")
	fmt.Printf("UserID: %+v\nID: %+v\nTitle: %+v\nBody: %+v\n", post[98].UserId, post[98].Id, post[98].Title, post[98].Body)
	fmt.Println("------------------------------------------------------")
	fmt.Printf("UserID: %+v\nID: %+v\nTitle: %+v\nBody: %+v\n", post[99].UserId, post[99].Id, post[99].Title, post[99].Body)
	fmt.Println("------------------------------------------------------")
}
