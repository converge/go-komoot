# Go Komoot library

This is an easy way to communicate your user with Komoot.

Via the Go Komoot library you can:

[ work in progress ] login into your Komoot account  
[ work in progress ] get Komoot user information

### Usage

```go
komootConfig := goKomoot.Config{
"client_id": "",
"redirect_url": ""
}

client := goKomoot.NewClient(komootConfig)

userData := client.GetUserProfile()

fmt.Println(userData.Name)
```

### Developers Guide

### Tests

Running tests:

```go
go test -v./...
```

### todo:

[ ] have fun!  
[ ] add CI  
[ ] add semantic release  