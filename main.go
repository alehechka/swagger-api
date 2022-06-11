package main

import (
	_ "github.com/alehechka/swagger-api/docs"
	"github.com/alehechka/swagger-api/rest"
)

// @title User API documentation
// @version 1.0.0
// @host localhost:3000
// @BasePath /api/v1

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	check(rest.SetupRouter().Run())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
