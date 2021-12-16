FROM public.ecr.aws/docker/library/golang:1.17 AS builder
WORKDIR /go/src/github/mmmknt-sandbox/github-actions-demo/
COPY *.go go.* ./
COPY logic/* ./logic/
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM public.ecr.aws/docker/library/alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github/mmmknt-sandbox/github-actions-demo/app .
EXPOSE 8080
CMD ["./app"]
