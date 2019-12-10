FROM golang:1.13.5-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o sumfinder ./main.go

CMD ["go/src/sumfinder"]

FROM scratch

COPY --from=builder /go/src/sumfinder /go/src/sumfinder

ENTRYPOINT [ "/go/src/sumfinder" ]