This is a GoLang Template that will be used 
as FDS ASYA PHILIPPINES INC. standard format.

#******************************************#
#*********** GoSwag is Required ***********#
#******************************************#

Usage

1. Add comments to your API source code, See Declarative Comments Format.
2. Download Swag for Go by using:
	go get -u github.com/swaggo/swag/cmd/swag
	# 1.16 or newer
	go install github.com/swaggo/swag/cmd/swag@latest

3. Run the Swag in the root folder of your Go project, which contains the main.go file. Then, The necessary files (docs folder and docs/doc.go) will be generated by Swag once it parses the comments.
	swag init
	If swag init command failed, type this then try swag init again.
	export PATH=$(go env GOPATH)/bin:$PATH
	then type swag init again

4. Download swagger by using:
	go get -u github.com/gofiber/swagger
	And import following in your code:
	import "github.com/gofiber/swagger" // swagger handler

5. Run it, browser to http://localhost:8080/docs/, you can see Swagger 2.0 Api documents.


#******************************************#
#*********** GoLang is Required ***********#
#******************************************#

1.  Download the template by using either:
        - git clone https://github.com/FDSAP-Git-Org/Go_Template.git
        or
        - Download the repository at Github
2.  Place the template in your GOROOT
        - To know your GOROOT, you can use this command:
                - go env
3.  Check or edit .env configuration before running the template
4.  Change the template to your desired application
    name by following this instruction:
        - Delete go.mod and go.sum
        - Initialize the application:
                - go mod init <Application name>
        - Change the imports from Template folder
          to your <Application name>
                - "Template/middleware"
                to
                - "<Application name>/middleware"
        - Get the application packages:
                - go mod tidy
        - Then run the application:
                - go run main.go
5.  Check if the application is running by using postman
    or visiting this link on a browser:
        - http://127.0.0.1:8000/api/public/v1
        or
        - localhost:8000/api/public/v1
6.  To deploy the application on a server,
    follow this instructions and command:
        - Upload the source code at the server
        - Navigate to the root folder
        - Run this command:
                - go build <Application Name> -o
                - nohup go run main.go
