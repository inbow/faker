FROM golang:1.14 AS build
ARG NAME
ENV NAME ${NAME}
WORKDIR /opt/${NAME}
COPY . .
RUN ["make", "build"]

FROM debian:buster-slim
ARG NAME
ARG ENV
ARG VERSION
ARG LOG_LEVEL
ENV NAME ${NAME}
ENV ENV ${ENV}
ENV VERSION ${VERSION}
ENV LOG_LEVEL ${LOG_LEVEL}
WORKDIR /opt/${NAME}
COPY --from=build /opt/${NAME}/configs/${NAME}/*.yml ./configs/${NAME}/
COPY --from=build /opt/${NAME}/bin/${NAME} ./${NAME}
CMD ./${NAME} -e ${ENV} -v ${VERSION} -ll ${LOG_LEVEL}
