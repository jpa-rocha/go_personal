FROM debian:stable as build

ENV GOLANG_VERSION  1.22.0
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
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /adamastor
RUN go mod init adamastor

#RUN go get github.com/a-h/templ &&\
#    go get github.com/go-playground/form &&\
#    go get github.com/spf13/cobra &&\
#    go get github.com/yuin/goldmark &&\
#    go get golang.org/x/crypto/acme/autocert@latest

RUN go mod tidy && go mod vendor
ENV PATH=$PATH:$WORKDIR

CMD [ "./run_server.sh" ]
