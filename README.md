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

    "github.com/GoWebFramework/lymon"
)

func users(w http.ResponseWriter, r *http.Request, c Context) {
    var result lymon.M
    
    ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
    c.Database.Collection("users").FindOne(ctx, lymon.M{}).Decode(&result)
    
    j, _ := json.Marshal(result)
    fmt.Fprint(w, j)
}

func main() {
    web := lymon.Context{}
    web.UseDefaultConfig()

    web.HandleFunc("/users", "GET", users)

    web.Start()
}
```
