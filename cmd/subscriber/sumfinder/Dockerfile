FROM golang:1.13.5-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o sumfinder ./main.go

CMD ["app/sumfinder"]

FROM scratch

COPY --from=builder /app/sumfinder /app/sumfinder

ENTRYPOINT [ "/app/sumfinder" ]