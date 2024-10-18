# Technical Test for Xsis

This project is for the technical test for Xsis.

## Requirements

You need to have Docker installed to run this program.

## Setup

Run the following command to start a PostgreSQL container:

```bash
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres
```

Then, you can tidy your Go modules:

```bash
go mod tidy
```

Finally, run the application:

```bash
go run main.go
```

You can customize the `.env` file according to your setup.

## Testing

To run the tests, use the following command:

```bash
go test -v ./test/
```
or
```bash
go test -v -count=1 ./test/  #If you don't wanna the test result come from cache
```



## Documentation

You can access the API documentation using Swagger at:

```
[HOST:PORT]/swagger/index.html
localhost:8089/swagger/index.html
```