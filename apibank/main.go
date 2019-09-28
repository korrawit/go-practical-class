package main

import (
	"apibank/user"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	uu, err := user.FetchUsers()
	if err != nil {
		panic(err)
	}

	fmt.Println("================= USERS ==================")
	for _, u := range uu[:10] {
		fmt.Println(u)
	}

	// fmt.Println("================= ALBUMS ==================")
	// aa, err := fetchAlbums()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, a := range aa[:10] {
	// 	fmt.Println(a)
	// }

	// fmt.Println("================= PHOTOS ==================")
	// pp, err := fetchPhotos()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, p := range pp[:10] {
	// 	fmt.Println(p)
	// }

	// fmt.Println("================= TODOS ==================")
	// tt, err := fetchTodos()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, t := range tt[:10] {
	// 	fmt.Println(t)
	// }
}

func fetchPhotos() ([]Photo, error) {
	resp, err := http.Get("http://jsonplaceholder.typicode.com/photos")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pp []Photo
	if err := json.Unmarshal(body, &pp); err != nil {
		return nil, err
	}

	return pp, nil
}

func fetchTodos() ([]Todo, error) {
	resp, err := http.Get("http://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tt []Todo
	if err := json.Unmarshal(body, &tt); err != nil {
		return nil, err
	}

	return tt, nil
}
