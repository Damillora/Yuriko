FROM golang:1.16-alpine AS build

WORKDIR /go/src/yuriko
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /yuriko

FROM scratch AS runtime
COPY --from=build /yuriko /

ENTRYPOINT ["/yuriko"]
