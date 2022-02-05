# Go Komoot library

This is an easy way to communicate your user with Komoot.

Via the Go Komoot library you can:

[ work in progress ] login into your Komoot account  
[ work in progress ] get Komoot user information

### usage

```go
komootConfig := goKomoot.Config{
"client_id": "",
"redirect_url": ""
}

client := goKomoot.NewClient(komootConfig)

userData := client.GetUserProfile()
fmt.Println(userData.Name)
```

