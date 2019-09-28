package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Album struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

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
	uu, err := fetchUsers()
	if err != nil {
		panic(err)
	}

	fmt.Println("================= USERS ==================")
	for _, u := range uu[:10] {
		fmt.Println(u)
	}

	fmt.Println("================= ALBUMS ==================")
	aa, err := fetchAlbums()
	if err != nil {
		panic(err)
	}

	for _, a := range aa[:10] {
		fmt.Println(a)
	}

	fmt.Println("================= PHOTOS ==================")
	pp, err := fetchPhotos()
	if err != nil {
		panic(err)
	}

	for _, p := range pp[:10] {
		fmt.Println(p)
	}

	fmt.Println("================= TODOS ==================")
	tt, err := fetchTodos()
	if err != nil {
		panic(err)
	}

	for _, t := range tt[:10] {
		fmt.Println(t)
	}
}

func fetchUsers() ([]User, error) {
	resp, err := http.Get("http://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func fetchAlbums() ([]Album, error) {
	resp, err := http.Get("http://jsonplaceholder.typicode.com/albums")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var aa []Album
	if err := json.Unmarshal(body, &aa); err != nil {
		return nil, err
	}

	return aa, nil
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
