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

 - Headers:

   `curl --request GET -b cookie.txt -H "Origin: blinnikov.com" --verbose http://localhost:8080/private/whoami`

### Migrations
- create
   ``` bash
   migrate create -ext sql -dir migrations migration_name  
   ```
- run migration
   ``` bash
    migrate -path migrations -database "postgresql://postgres:changeme@localhost/restapi_test?sslmode=disable" up
   ```

### Serving through TLS
- install `openssl`
   ``` bash
   brew update
   brew install openssl
   ```
- genereate Self-Signed Certificate
   ``` bash
   openssl req -newkey rsa:4096 \
            -x509 \
            -sha256 \
            -days 3650 \
            -nodes \
            -out go-rest-api.crt \
            -keyout go-rest-api.key
   ```
- use `http.ListenAndServeTLS`
   ``` go
   return http.ListenAndServe(config.BindAddr, srv)
   ```
   ->
   ``` go
   return http.ListenAndServeTLS(config.BindAddr, "go-rest-api.crt", "go-rest-api.key", srv)
   ```

### Deploy to local k8s
- Give `create-configmap.sh` script execution permissions
   ``` bash
   chmod +x ./k8s/create-configmap.sh 
   ```
- Run this script to create config maps for certificate files
   ``` bash
   ./k8s/create-configmap.sh
   ```

- Install all kubernetes objects
   ``` bash
   kubectl apply -f ./k8s --recursive
   ```