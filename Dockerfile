FROM golang AS bob-the-builder

WORKDIR /app

COPY src/main.go .

RUN go build main.go

FROM scratch

WORKDIR /app

COPY --from=bob-the-builder /app/main .

ENTRYPOINT ["/app/main"]
