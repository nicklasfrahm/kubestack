FROM golang:1.20 AS build
ARG VERSION
ARG BINARY

WORKDIR /app
COPY go.* ./
RUN go mod download

ADD Makefile .
ADD cmd/ cmd/
ADD api/ api/
ADD pkg/ pkg/
ADD controllers/ controllers/
RUN VERSION=$VERSION make bin/$BINARY \
  && mv bin/$BINARY app

FROM gcr.io/distroless/base:nonroot AS run
WORKDIR /
COPY --from=build /app/app .
USER 65532:65532
ENTRYPOINT [ "/app" ]
