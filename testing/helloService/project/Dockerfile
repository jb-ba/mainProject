FROM golang:alpine as builder
EXPOSE 3000
COPY src .
RUN adduser -D -H -u 10001 scratchuser && \
    cd /go/main && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /go/main/main /
COPY --from=builder /etc/passwd /etc/passwd
USER scratchuser
CMD ["./main"]