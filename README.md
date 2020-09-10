# README

`httmock` allows local testing and prototyping by mocking a HTTP server from a file based specification.

## Installing

1. Clone this repository: `git clone https://github.com/benmezger/httmock`:
2. Run `make install`

## Running

Make sure you have a `.http.yaml` in your current directory or pass a custom filename with `--api-file`

- `httmock serve` or `httmock --api-file <filename> serve`

Use `httmock serve --host localhost:8000` for using a different host/port.

## HTTP file syntax

For an example of the file structure, see [example-http.yaml](example-http.yaml).

## TODOs

- Handle context base requests and responses

## Example

Run `httmock` with the [example](example-http.yaml) file
Then run curl against the defined URL paths:

- `curl -X GET http://localhost:8000/?name=name-param -d '{"msg": "Body of GET request"}'`
  `{"msg": "Hello, from / GET response"}`
