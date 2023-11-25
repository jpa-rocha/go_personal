package utilities

import (
	"fmt"
	"html/template"

	"adamastor/public"
)

func GetTmpl(file string) (*template.Template, error) {
	tmpl := template.New("index")

	fileContents, err := public.Assets.ReadFile(file)
	if err != nil {
		return tmpl, fmt.Errorf("error: %w", err)
	}

	htmlString := string(fileContents)

	tmpl, err = tmpl.Parse(htmlString)
	if err != nil {
		return tmpl, fmt.Errorf("error: %w", err)
	}

	return tmpl, nil
}
