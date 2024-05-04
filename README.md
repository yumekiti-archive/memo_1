## Backend

### Create

```
$ go mod init backend
go: creating new go.mod: module backend
$ go get -u github.com/gin-gonic/gin
```

### Run

```
// https://gin-gonic.com/docs/quickstart/
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```

```
$ go run main.go
```