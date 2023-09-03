FROM alpine:latest

RUN mkdir /app

# copy the backendApp binary to the container
COPY ./cmd/backendApp /app

CMD ["/app/backendApp"]