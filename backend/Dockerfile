FROM golang:1.17-buster AS build
WORKDIR /app
COPY . ./
RUN go build -o /tamil-aadal tamil-aadal.go
# Build section completed


# Deploy section begins
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /tamil-aadal /tamil-aadal
ADD ui1 /ui1
ADD ui2 /ui2
ADD ui3 /ui3
USER nonroot:nonroot
ENTRYPOINT ["/tamil-aadal"]
