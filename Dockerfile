# build the app
FROM golang:alpine AS build-env
WORKDIR /src
COPY . .
RUN go build -o awesomeApp
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o awesomeApp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/awesomeApp /app/