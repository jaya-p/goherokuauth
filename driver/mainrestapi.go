// build: go build driver/mainrestapi.go
// run: go run driver/mainrestapi.go
// test:
//				POST: curl -d '{"Username":"me", "PasswordHash":"me"}' -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/auth
//				DELETE: curl -X DELETE http://localhost:8080/api/v1/auth

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jaya-p/goherokuauth"
)

func main() {
	httpPortNumber, ok := os.LookupEnv("PORT")
	if !ok {
		httpPortNumber = "8080"
	}

	fmt.Println("REST API web server is running on port " + httpPortNumber)

	httpPortNumberInUint, _ := strconv.ParseUint(httpPortNumber, 10, 32)
	goherokuauth.RestAPIWebserver(uint(httpPortNumberInUint))
}
