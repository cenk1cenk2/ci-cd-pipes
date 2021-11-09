FROM alpine:latest

ARG PIPE_FILE

COPY ${PIPE_FILE} /pipe

RUN chmod +x /pipe

ENTRYPOINT [ "/pipe" ]
