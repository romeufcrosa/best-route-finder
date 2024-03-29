FROM golang:1.11 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/romeufcrosa/best-route-finder
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app ./cmd/...

#FROM scratch
FROM golang:1.11
WORKDIR /release
ENV ENV=docker
ENV MYSQL_ADDRESS=host.docker.internal
COPY --from=builder /go/src/github.com/romeufcrosa/best-route-finder/assets/sql/*.sql /assets/sql/
COPY --from=builder /go/src/github.com/romeufcrosa/best-route-finder/ /go/src/github.com/romeufcrosa/best-route-finder/
COPY --from=builder /go/src/github.com/romeufcrosa/best-route-finder/configurations/*.json ./configurations/
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
