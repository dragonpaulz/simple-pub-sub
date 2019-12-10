FROM golang:1.13.5-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o publisher ./main.go

CMD ["go/src/publisher"]

FROM scratch

COPY --from=builder /go/src/publisher /go/src/publisher

ENTRYPOINT [ "/go/src/publisher" ]