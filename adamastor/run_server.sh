#!/usr/bin/bash
go get github.com/a-h/templ &&\
go get github.com/go-playground/form &&\
go get github.com/spf13/cobra &&\
go get github.com/yuin/goldmark &&\
go get golang.org/x/crypto/acme/autocert
go get golang.org/x/crypto/acme

# Generate CSS
sass /adamastor/public/assets/static/css/global.scss /adamastor/public/assets/static/css/global.css

# Generate go templates
templ generate -path /adamastor/internal/server/components/
templ generate -path /adamastor/internal/server/templates/

# Build server
go build -o /adamastor
./adamastor serve
