FROM golang:1.20.0-bullseye  AS BUILDER

WORKDIR /app
COPY . .

RUN GOOS=linux; GOARCH=amd64; go build -o /app/bin/treebabel /app/cmd/treebabel/main.go

########################################

FROM debian:bullseye AS FINAL

COPY --from=BUILDER /app/bin/treebabel /usr/local/bin

WORKDIR /unbabel-data
COPY --from=BUILDER /app/testdata/treebabel/* /unbabel-data/