FROM golang as builder
EXPOSE 5688
COPY . .
RUN useradd scratchuser && \
    export GOPATH="" && go env  && go mod vendor  && cd src/main && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /go/src/main/main /
COPY --from=builder /etc/passwd /etc/passwd
USER scratchuser
CMD ["./main"]