FROM alpine:latest
RUN apk --no-cache add ca-certificates
#COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY dist/ .
EXPOSE 50055
CMD ["/service"]