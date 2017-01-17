# VYNYL Golang HTTP Scaffold

## Building from source

* Install Go version >= 1.6 to support vendoring
* Ensure you have configured [a proper Go workspace](https://golang.org/doc/code.html#Organization)
* [Install Glide](https://github.com/Masterminds/glide)
* Run `$ glide up` to download vendor packages
* Copy .env file and fill in fields
* Run `$ go build .`
* `$ ./go-http-scaffold`

## Testing

* `$ go test ./...`