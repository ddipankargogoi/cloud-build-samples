// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample run-helloworld is a minimal Cloud Run service.

package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
)
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})
// 	log.Println("Listening on localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
func main() {
	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewHelloAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()
	server.Port = 8080
	// Implement the CheckHealth handler
	// api.CheckHealthHandler = operations.CheckHealthHandlerFunc(
	// 	func(user operations.CheckHealthParams) middleware.Responder {
	// 		return operations.NewCheckHealthOK().WithPayload("OK")
	// 	})
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)
	// Implement the GetHelloUser handler
	// api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(
	// 	func(user operations.GetHelloUserParams) middleware.Responder {
	// 		return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
	// 	})
	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)
	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
//Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}
//GetHelloUser returns Hello + your name
func GetHelloUser(user operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
}
