# Go simple CRUD API

This repo demonstrate on how you can create APIs and backend services with go.

## Installation
Make sure you have go installed if not checkout [Go installation](https://go.dev/doc/install).

Fork the repo.

install [gorilla/mux](https://github.com/gorilla/mux).
```bash
go mod init <project name>
go get -u github.com/gorilla/mux
```

## Usage

```bash
go build
go run main.go
```

## Docker

The project can also be run with Docker. Make sure that you have Docker installed.

```bash
docker build --tag go-movies-crud .
docker run -p 8080:8080 -d --name go-movies-crud go-movies-crud
```

To stop the container simply run:
```bash
docker stop go-movies-crud
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[GNU](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)
