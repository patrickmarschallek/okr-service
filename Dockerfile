FROM alpine:3.3
MAINTAINER Patrick Marschallek <patrick.marschallek@gmail.com>
ADD config.yaml config.yaml
ADD okr-service okr-service
ENV PORT 8080
EXPOSE 8080 8080

ENTRYPOINT ["/okr-service"]