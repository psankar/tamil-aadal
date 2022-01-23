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
COPY auth /auth
USER nonroot:nonroot
ENV GOOGLE_APPLICATION_CREDENTIALS /auth/vetchi-dev-firebase-adminsdk-cjf4f-80eddf2908.json
ENTRYPOINT ["/tamil-aadal"]
