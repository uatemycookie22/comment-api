FROM golang:1.22.0

WORKDIR /app

# copy mod \
COPY go.mod go.sum ./

# download
RUN go mod download

# copy code
COPY *.go ./
COPY ./models/*.go ./models/

# copy db
COPY identifier.sqlite ./

RUN go build -o /comments-api

EXPOSE 8080

CMD ["/comments-api"]