FROM golang:1.15-buster AS build
ARG ACTION
WORKDIR /opt/faker
COPY . .
RUN make $ACTION

FROM debian:buster-slim
ARG ENVIRONMENT
ARG LOG_LEVEL
ENV ENV ${ENV}
ENV LOG_LEVEL ${LOG_LEVEL}
WORKDIR /opt/faker
COPY --from=build /opt/faker/bin/faker ./faker
COPY --from=build /opt/faker/configs/faker/*.yml ./configs/faker/
CMD ./faker
