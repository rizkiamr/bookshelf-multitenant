FROM golang:1.21.0 AS build

RUN mkdir -p /opt/build

WORKDIR /opt/build

# Copy only necessary files
COPY go.mod ./
RUN go mod download

# Copy the rest of the files
COPY . .

# Do the build
WORKDIR /opt/build/cmd/api
RUN CGO_ENABLED=0 go build -o main .

FROM gcr.io/distroless/static
USER nobody:nobody
WORKDIR /
COPY --from=build /opt/build/cmd/api/main .
ENTRYPOINT []
