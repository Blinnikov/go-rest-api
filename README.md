# go-rest-api
Learning Go through creating simple REST API HTTP server


Not to forget links:
 - [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
 - [Codeship Golang Best Practices](https://github.com/codeship/go-best-practices)
 - [SQL library to be used](http://go-database-sql.org/index.html)

 Commands:
 - Create user:

    `curl --request POST --data '{"email":"user@domain.com","password":"1234567"}' --verbose http://localhost:8080/users`
 - Log In:

    `curl --request POST --data '{"email":"user@domain.com","password":"1234567"}' -c cookie.txt  --verbose http://localhost:8080/sessions`
 - Get current user name (passed by middleware through context):

    `curl --request GET -b cookie.txt --verbose http://localhost:8080/private/whoami`