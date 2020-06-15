// build: go build driver/restapiwebserver/mainrestapi.go
// run: go run driver/restapiwebserver/mainrestapi.go
// test:
//				GET: curl http://localhost:8080/api/v1/helloworld
//				POST: curl -d '{"name":"indonesia"}' -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/helloworld
//				DELETE: curl -X DELETE http://localhost:8080/api/v1/helloworld

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jaya-p/goheroku"
)

func main() {
	httpPortNumber, ok := os.LookupEnv("PORT")
	if !ok {
		httpPortNumber = "8080"
	}

	fmt.Println("REST API web server is running on port " + httpPortNumber)

	httpPortNumberInUint, _ := strconv.ParseUint(httpPortNumber, 10, 32)
	goheroku.HelloworldRestAPIWebserver(uint(httpPortNumberInUint))
}
