package main

import (
	"fmt"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func main() {
	port := "8090"
	restful.Add(NewStaticService(port))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func NewStaticService(port string) *restful.WebService {
	path := "/0.1/static"
	posts := "/posts"
	fmt.Printf("Serving posts at http://%s:%s%s\n", "localhost", port, path+posts)

	service := new(restful.WebService)
	service.
		Path(path).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	//service.Route(service.POST(facebookCallback).To(FacebookCallback))
	service.Route(service.GET(posts).To(Posts))
	service.Filter(enableCORS)
	return service
}

func enableCORS(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	if origin := req.Request.Header.Get("Origin"); origin != "" {
		resp.AddHeader("Access-Control-Allow-Origin", origin)
	}
	chain.ProcessFilter(req, resp)
}

type test struct{ t string }

func Posts(r *restful.Request, response *restful.Response) {
	response.WriteEntity(createPosts())
}

func createPosts() []*Post {
	posts := []*Post{}
	for i := 0; i < 3; i++ {
		p := NewPost(i)
		posts = append(posts, p)
	}
	return posts
}

func NewPost(i int) *Post {
	return &Post{
		i,
		1,
		"https://www.example.com",
		"https://www.example.com",
		"test.html",
		"asdf",
		"https://www.example.com",
		"test title",
		"test-title",
		"test 1, 2",
		"asdf",
		"test 1, 2"}
}

type Post struct {
	Id          int
	Version     int
	ThumbImg    string
	PostImg     string
	Filename    string
	Date        string
	Url         string
	Title       string
	TitlePlain  string
	Excerpt     string
	Content     string
	DsqThreadId string
}
