# GoLang REST API 

## REST API project written in Go that retrieves user information (i.e. coin balance) from a database using http requests. 

This project was made to create a custom REST API in Golang.  

The code displays how to do the following:

* Create a custom web server with Golang using chi Router
* Build Golang Code Structure
* Implement API Authentication
* Add Middleware
* Logging Errors
* Exporting External Packages
* Create custom JSON responses

## Installation Requirements:

* Go - https://go.dev/doc/install
* chi (http router), logrus (error logging), gorilla/schema (decoder to convert struct to and from form values)
    - All can be installed with command "go mod tidy" where go.mod file is located
 
## How To Run the Program

1. Clone the project.
2. In the terminal, navigate to /goapi folder.
3. Enter the command "go run ./cmd/api/main.go" .
4. While the server is running, open an API testing software such as Postman - https://www.postman.com/downloads/
5. Create a new HTTP GET request using http://localhost:8000/account/coins/?username={name}
     - Replace {name} with a person from the database (e.g. alex)
     - Add a header key: Authorization and value: {token} (e.g. 123ABC)
6. Send the http request and you should see a JSON response in return with the person's coin balance from the database.
