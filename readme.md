## How to setup project?

- Find location for `GOPATH`

- Inside GOPATH create directory `github.com/shivkumarsingh7/spicework-assignment` with all subdirectories

- Paste all file in `spicework-assignment` directory

- run `make` command inside root of project in terminal

Open URL in browser http://localhost:8080

Wait for 20 seconds :-  resources health status will updated

By default all resources health will return `200` status

for testing `500` status please change URL for any resource with wrong path
Edit any resource status in `spicework-assignment/util/monitor.go`

Available endpoint on application

http://localhost:8080/

http://localhost:8080/hello