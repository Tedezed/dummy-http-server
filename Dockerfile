FROM golang:1.22 AS builder
COPY src/ .
RUN go build dummy-http-server.go

FROM debian:stable-slim
ARG UNAME=dummy
ARG UID=1000
ARG GID=1000
ARG DUMMY_HOST=""
ARG DUMMY_PORT="8080"
ENV DUMMY_HOST = $DUMMY_HOST
ENV DUMMY_PORT = $DUMMY_PORT
# Debian
RUN groupadd -f -g ${GID} ${UNAME} \
    && useradd -l -u ${UID} -g ${UNAME} -r ${UNAME}
# Alpine
# RUN addgroup -g ${GID} ${UNAME} \
#     && adduser -G ${UNAME} -u ${UID} ${UNAME} -D
COPY --from=builder /go/dummy-http-server /dummy-http-server
RUN chmod +x dummy-http-server
USER 1000
CMD ["./dummy-http-server"]
