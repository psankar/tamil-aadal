FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o /tamil-wordle

# Build section completed

# Deploy section begins
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /tamil-wordle /tamil-wordle
ADD ui1 /ui1
ADD ui2 /ui2
ADD ui3 /ui3

USER nonroot:nonroot

ENTRYPOINT ["/tamil-wordle"]