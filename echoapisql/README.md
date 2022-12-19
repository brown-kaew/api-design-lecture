run server
```{shell}
DATABASE_URL="postgres://root:password@tcp(127.0.0.1:3306)/dbname" go run server.go
```

run test
```{shell}
AUTH_TOKEN="Basic YXBpZGVzaWduOjQ1Njc4" go test -v -tags=integration
```