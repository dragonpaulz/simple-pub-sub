FROM golang:1.13.4-alpine AS builder

WORKDIR /app/
COPY . .

RUN go build -o publisher ./main.go

CMD ["/app/publisher"]

FROM scratch

COPY --from=builder /app/publisher /app/publisher

ENTRYPOINT [ "/app/publisher" ]