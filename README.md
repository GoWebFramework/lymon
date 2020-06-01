<img align="right" width="100" height="100" src="https://github.com/GoWebFramework/lymon/blob/master/lemon.png?raw=true">

# Lymon

[![Go Report Card](https://goreportcard.com/badge/github.com/GoWebFramework/lymon)](https://goreportcard.com/report/github.com/GoWebFramework/lymon)

Advanced Web API Framework for Go, Project Status : WIP

```sh
go get github.com/GoWebFramework/lymon
```

## Feature

- [x] Include MongoDB & Redis
- [x] JSON & Form payload validation
- [x] BeforeAll middleware
- [x] StatusCodeHandler middleware
- [ ] cmd feature
- [ ] CRUD Generator 
- [ ] Built-in cache wrapper
- [ ] Simplify Wrapper
- [ ] Sample

## Usage

```go
package main

import (
	"context"
	"log"
	"time"

	"./lymon"
)

func users(ctx lymon.Context, g *lymon.Global) {
	var result lymon.M

	to, _ := context.WithTimeout(context.Background(), 5*time.Second)
	g.DB().C("users").FindOne(to, lymon.M{}).Decode(&result)

	ctx.JSON(result)
}

func index(ctx lymon.Context, g *lymon.Global) {
	if ctx.V.IsValidated {
		ctx.JSON(map[string]string{
			"message": "valid payload!",
		})
	} else {
		ctx.W.Write([]byte(`{"message": "invalid payload !"}`))
	}
}

func main() {
	web := lymon.Global{}
	// web.UseDefaultConfig()
	web.UseConfig(lymon.Config{
		MongoURI: "mongodb://127.0.0.1:27017",
		RedisURI: "redis://127.0.0.1:6379/0",
		Listen:   "127.0.0.1:8080",
	})

	// set default mongo database
	web.Database = "my_site"

	web.HandleFunc("/users", "GET", users, nil)

	// Look at https://github.com/asaskevich/govalidator#validatestruct-2
	// for validation reference
	web.HandleFunc("/", "POST", index, map[string]interface{}{
		// curl request for test
		// invalid payload :  curl -d '{"name": "val2Conta1nNumb3r"}' http://localhost:8080
		// valid payload :  curl -d '{"name": "alphaOnly"}' http://localhost:8080
		"name": "required,alpha",
	})

	// before all middleware support
	web.BeforeAll(func(ctx lymon.Context, g *lymon.Global) {
		ctx.W.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.W.Write([]byte("before all modifier \n"))
		return
	})

	// handle specific custom code, currently only applied for 404
	web.HandleStatusCode(404, func(ctx lymon.Context, g *lymon.Global) {
		ctx.W.Write([]byte("this is custom 404 page"))
	})

	log.Println("starting server...")
	web.Start()
}
```

## Credits

- [Lemon Image https://www.pngmart.com/files/1/Lemon-PNG-Transparent-Image.png](https://www.pngmart.com/files/1/Lemon-PNG-Transparent-Image.png)
