FROM debian:stable as build

ENV GOLANG_VERSION  1.21.5
ENV SASS_VERSION 1.69.5 
ENV CERTS=/usr/local/share/ca-certificates

RUN apt-get update && apt-get upgrade -y &&\
    apt-get -y install --no-install-recommends wget\
                                               git\ 
                                               ca-certificates\
                                               openssl &&\
    rm -rf /var/lib/apt/lists/*

RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${CERTS}/github.crt
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${CERTS}/proxy.golang.crt
RUN update-ca-certificates

RUN wget --no-check-certificate --progress=dot:giga "https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" &&\
            rm -rf /usr/local/go && tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && rm go${GOLANG_VERSION}.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

RUN wget --progress=dot:giga https://github.com/sass/dart-sass/releases/download/${SASS_VERSION}/dart-sass-${SASS_VERSION}-linux-x64.tar.gz &&\
    tar -xzf dart-sass-${SASS_VERSION}-linux-x64.tar.gz && rm dart-sass-${SASS_VERSION}-linux-x64.tar.gz

RUN go install github.com/a-h/templ/cmd/templ@latest

ENV PATH=$PATH:/dart-sass
ENV PATH=$PATH:/root/go/bin

FROM build

WORKDIR /adamastor/
ENV PATH=$PATH:$WORKDIR
COPY adamastor/ .
RUN go get -u github.com/a-h/templ
RUN go mod vendor && go mod tidy

CMD [ "./run_server.sh" ]
