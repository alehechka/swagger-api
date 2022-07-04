# BUILD

FROM golang:1.18-alpine as go-builder

WORKDIR /app
COPY . ./

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.2

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN swag init

RUN go build main.go

# SERVE

FROM busybox

COPY --from=go-builder /app/main server

RUN mkdir docs
COPY --from=go-builder /app/docs docs

ENV PORT=80
ENV GO_ENV="production"
ENV GIN_MODE="release"

EXPOSE 80
CMD [ "/server" ]