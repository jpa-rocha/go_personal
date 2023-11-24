package serve

import (
	"io/fs"
	"log"
	"net/http"

	"adamastor/public"
)

func (s *Server) setStaticPaths() {
	static := []string{"/static/", "/components/", "/routes/"} // TODO: should be in config.

	const base = "assets"
	staticFS, err := fs.Sub(public.Assets, base)
	if err != nil {
		s.Err = err
		return
	}
	for _, path := range static {
		s.Router.Handle(path, http.FileServer(http.FS(staticFS)))
	}
}

func (s *Server) handleRoutes() {
	s.setStaticPaths()
	s.Router.HandleFunc("/cv", handleCV)

	s.Router.HandleFunc("/components/test", handleComponentTest)
	s.Router.HandleFunc("/components/nav", handleComponentNav)

	s.Router.HandleFunc("/", handleIndex)
}

func handleCV(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := getTmpl("assets/routes/cv/cv.html")
	if err != nil {
		log.Println(err)
	}

	if err = tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func handleIndex(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := getTmpl("assets/routes/index/index.html")
	if err != nil {
		log.Println(err)
	}
	if err = tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}
