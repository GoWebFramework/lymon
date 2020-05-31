# Lymon

Adavanced Web API Framework

## Installation

```sh
go get -u github.com/GoWebFramework/lymon
```

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
	c.DB.C("users").FindOne(ctx, lymon.M{}).Decode(&result)

	w.Write(c.JSON(result))
}

func main() {
	web := lymon.Context{}
    // web.UseDefaultConfig()
    web.UseConfig(Config{
		MongoURI: "mongodb://127.0.0.1:27017",
		RedisURI: "redis://127.0.0.1:6379/0",
		Listen:   "127.0.0.1:8080",
	})

	// set default mongo database
	web.Database = "my_site"

	web.HandleFunc("/users", "GET", users)

    log.Println("starting server...")
	web.Start()
}
```

## To Do List

- [x] Include MongoDB & Redis
- [ ] Auth Support
- [ ] Cache Support
- [ ] Simplify Wrapper
