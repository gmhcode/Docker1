FROM golang:1.15.2

WORKDIR src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

EXPOSE 8081