FROM golang:1.14 AS build
ARG NAME
ENV NAME ${NAME}
WORKDIR /opt/${NAME}
COPY . .
RUN ["make", "build"]

FROM debian:buster
ARG ENV
ARG VERSION
ENV NAME "rtb-mock"
WORKDIR /opt/${NAME}
COPY --from=build /opt/${NAME}/configs/*.yml ./configs/
COPY --from=build /opt/${NAME}/bin/${NAME} ./${NAME}
CMD ./${NAME}
