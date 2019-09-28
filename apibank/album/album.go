package album

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Album struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

func FetchAlbums() ([]Album, error) {
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
