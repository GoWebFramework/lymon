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
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GoWebFramework/lymon"
)

func users(w http.ResponseWriter, r *http.Request, c lymon.Context) {
	var result lymon.M

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c.Database.Collection("users").FindOne(ctx, lymon.M{}).Decode(&result)

	j, _ := json.Marshal(result)
	fmt.Fprint(w, string(j))
}

func main() {
	web := lymon.Context{}
	web.UseDefaultConfig()

	// set default mongo database
	web.Database = web.Mongo.Database("my_site")

	web.HandleFunc("/users", "GET", users)

	web.Start()
}
```

## To Do List

- [x] Include MongoDB & Redis
- [ ] Auth Support
- [ ] Cache Support
- [ ] Simplify Wrapper
