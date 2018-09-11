FROM zkrhm/golang:1.10.4-alpine

ARG app_name
ARG package_name

ENV APP_NAME=${app_name}
ENV PACKAGE=${package_name}

RUN mkdir -p $GOPATH/src/$PACKAGE
WORKDIR $GOPATH/src/$PACKAGE/

COPY . .
COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

RUN make install
ENTRYPOINT [ ${app_name} ]

# EXPOSE 8000