FROM golang:1.14 AS build
ARG NAME
ENV NAME ${NAME}
WORKDIR /opt/${NAME}
COPY . .
RUN ["make", "build"]

FROM debian:buster-slim
ARG NAME
ARG ENV
ARG LOG_LEVEL
ENV NAME ${NAME}
ENV ENV ${ENV}
ENV LOG_LEVEL ${LOG_LEVEL}
WORKDIR /opt/${NAME}
COPY --from=build /opt/${NAME}/configs/${NAME}/*.yml ./configs/${NAME}/
COPY --from=build /opt/${NAME}/bin/${NAME} ./${NAME}
CMD ./${NAME} -e ${ENV} -ll ${LOG_LEVEL}
