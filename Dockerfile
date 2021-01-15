FROM golang:1.15-alpine AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/service

FROM scratch
COPY --from=build /bin/service /bin/service
EXPOSE 50055
ENTRYPOINT ["/bin/service"]
