FROM golang:1.22.2-alpine as builder
WORKDIR /data/voice-assistant-manager-code
RUN apk add --no-cache upx ca-certificates tzdata
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o voice-assistant-manager

FROM centos:7 as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /data/voice-assistant-manager-code/voice-assistant-manager /voice-assistant-manager
EXPOSE 9090
CMD ["/voice-assistant-manager"]
