package templates

import (
	"fmt"
	"html/template"
	// "io/fs"
	// "log"

	"adamastor/public"
)

func GetTmpl(file string) (*template.Template, error) {
	const base = "assets/templates"
	//    templatesFS, err := fs.Sub(public.Assets, base)
	// if err != nil {
	// 	log.Fatal(err)
	// }
    t := template.Must(template.ParseFS(
        public.Assets,
        "assets/templates/layout.html",
        "assets/routes/index/index.html",))
    // t := template.Must(template.New("t").ParseFS(templatesFS, "layout.html"))
	fmt.Println(t.DefinedTemplates())
	return t, nil
}

// tmpl := template.New("index")
// fileContents, err := public.Assets.ReadFile(file)
// if err != nil {
// 	return tmpl, fmt.Errorf("error: %w", err)
// }
//
// htmlString := string(fileContents)
//
// tmpl, err = tmpl.Parse(htmlString)
// if err != nil {
// 	return tmpl, fmt.Errorf("error: %w", err)
// }
