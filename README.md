<img align="right" width="100" height="100" src="https://www.pngmart.com/files/1/Lemon-PNG-Transparent-Image.png">

# Lymon

[![Go Report Card](https://goreportcard.com/badge/github.com/GoWebFramework/lymon)](https://goreportcard.com/report/github.com/GoWebFramework/lymon)

Advanced Web API Framework for Go, Project Status : WIP

```sh
go get github.com/GoWebFramework/lymon
```

[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/xS1D-tTrKmQ/0.jpg)](https://www.youtube.com/watch?v=xS1D-tTrKmQ)

_Book Store CRUD With Lymon._

## Feature

- [x] Include MongoDB & Redis
- [x] BeforeAll middleware
- [x] StatusCodeHandler middleware
- [ ] cmd feature
- [ ] CRUD Generator 
- [ ] Built-in cache wrapper
- [ ] Simplify Wrapper
- [ ] Auth Sample

## Usage

```go
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/GoWebFramework/lymon"
)

func users(w http.ResponseWriter, r *http.Request, c lymon.Context) {
	var result lymon.M

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c.DB().C("users").FindOne(ctx, lymon.M{}).Decode(&result)

	w.Write(c.JSON(result))
}

func main() {
	web := lymon.Context{}
	// web.UseDefaultConfig()
	web.UseConfig(lymon.Config{
		MongoURI: "mongodb://127.0.0.1:27017",
		RedisURI: "redis://127.0.0.1:6379/0",
		Listen:   "127.0.0.1:8080",
	})

	// set default mongo database
	web.Database = "my_site"

	web.HandleFunc("/users", "GET", users)

	// before all middleware support
	web.BeforeAll(func(w http.ResponseWriter, r *http.Request, c lymon.Context) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	})

	// handle specific custom code, currently only applied for 404
	web.HandleStatusCode(404, func(w http.ResponseWriter, r *http.Request, c lymon.Context) {
		w.Write([]byte("this is custom 404 page"))
	})

	log.Println("starting server...")
	web.Start()
}
```

## Credits

- [Lemon Image https://www.pngmart.com/files/1/Lemon-PNG-Transparent-Image.png](https://www.pngmart.com/files/1/Lemon-PNG-Transparent-Image.png)
