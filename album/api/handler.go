package api

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
	"net/http"
	"strconv"
)

var connection *gorm.DB

func init()  {
	name := "paviyarasu95@gmail.com"
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", name, name, "143.110.190.177", name)
	connection,err := gorm.Open("mysql", uri)
	if err != nil {
		fmt.Println("Failed to create database connection error: ", err)
		return
	}
	fmt.Println("connection created ", connection)
}

func getData(w http.ResponseWriter, r *http.Request){
	renderer := render.New()

	getAlbumsResponse := new(struct{
		Albums []Album
		Photos []Photo
	})


	var albums []Album

	response,err := externalRequest("GET", "https://jsonplaceholder.typicode.com/albums")

	err = json.NewDecoder(response.Body).Decode(&albums)
	if err != nil {
		message,_ := fmt.Println("Failed to read response error: ",err)
		renderer.JSON(w, 500, message)
		return
	}

	connection.Create(albums)
	getAlbumsResponse.Albums = albums

	for _,album := range albums{
		var photos []Photo
		response,err = externalRequest("GET", fmt.Sprint("https://jsonplaceholder.typicode.com/photos?albumId=",strconv.Itoa(album.ID)))

		err = json.NewDecoder(response.Body).Decode(&photos)
		if err != nil {
			message,_ := fmt.Println("Failed to read response error: ",err)
			renderer.JSON(w, 500, message)
			return
		}

		connection.Create(photos)
		getAlbumsResponse.Photos = append(getAlbumsResponse.Photos, photos...)
	}

	renderer.JSON(w, 200, getAlbumsResponse)
}

func externalRequest(method string,url string) (*http.Response,error){

	requestClient := http.Client{}

	request,err := http.NewRequest(method,url,nil)
	if err != nil{
		fmt.Println("Failed to create request error: ", err)
		return nil,err
	}

	response,err := requestClient.Do(request)
	if err != nil {
		fmt.Println("Failed to get data from api request error: ",err)
		return nil,err
	}

	return response,nil
}

func search(w http.ResponseWriter, r *http.Request)  {
	renderer := render.New()

	var tableName,id,albumId string
	tableName = r.URL.Query().Get("type")
	id = r.URL.Query().Get("id")
	albumId = r.URL.Query().Get("album")

	var response interface{}
	switch tableName {
	case "album":
		var album Album
		connection.Where("id=?",id).Find(&album)
		response = album
	case "photo":
		var photo Photo
		connection.Where("id=? and album=?",id,albumId).Find(&photo)
		response = photo
	default:
		renderer.JSON(w, 404, "invalid search")
		return
	}
	renderer.JSON(w, 200, response)
}
