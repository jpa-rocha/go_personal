#!/usr/bin/bash

# Generate CSS
sass /adamastor/public/assets/static/css/global.scss /adamastor/public/assets/static/css/global.css

# Generate go templates
templ generate -path /adamastor/internal/server/components/
templ generate -path /adamastor/internal/server/templates/

# Build server
go build -o /adamastor
./adamastor serve
