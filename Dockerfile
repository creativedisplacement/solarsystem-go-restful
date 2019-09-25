FROM golang:latest as builder  
RUN mkdir -p /go/src/folder   
WORKDIR /go/src/folder    
COPY . .
WORKDIR /go/src/folder/api
#RUN apk add --no-cache git
RUN go get github.com/emicklei/go-restful
RUN go get github.com/jmoiron/sqlx  
RUN go get github.com/mattn/go-sqlite3 
RUN go get github.com/solarsystem-go-restful/data
RUN go get github.com/solarsystem-go-restful/planet
RUN go get github.com/solarsystem-go-restful/planets   

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest  
WORKDIR /
COPY --from=builder /go/src/folder/app .  
CMD ["main.go"]