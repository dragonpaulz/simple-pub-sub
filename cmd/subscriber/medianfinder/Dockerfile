FROM golang:1.13.5-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o medianfinder ./main.go

CMD ["app/medianfinder"]

FROM scratch

COPY --from=builder /app/medianfinder /app/medianfinder

ENTRYPOINT [ "/app/medianfinder" ]