package main

import (
	"api/router"
)

func main() {

	r := router.SetupRouter()
	r.Run(":8080")

}

// Make a rest API with this URL (/v1/movie/<movie-id>)
// Tech stack Golang and DB mysql
// The output of the rest  API is as follows:-
// {
// "id": <movie-id>,
// "title": "<movie-name>",
// "poster_path": "<image-url>",
// "language": "<movie-language>",
// "overview": "string",
// "release_date" : "date"
// }
// gin framework with clean architecture
// Redis

// sanjay.agrawal@bobble.ai
// clean
