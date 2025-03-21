FROM golang:1.24.1-alpine AS build

WORKDIR /supermarket-historical-price-tracker

ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG SKAFFOLD_GO_GCFLAGS
RUN echo "Go gcflags: ${SKAFFOLD_GO_GCFLAGS}"
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -mod=readonly -v -o /app ./cmd/app

FROM gcr.io/distroless/static-debian11

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

WORKDIR /supermarket-historical-price-tracker
COPY --from=build /app ./app

ENTRYPOINT ["./app"]