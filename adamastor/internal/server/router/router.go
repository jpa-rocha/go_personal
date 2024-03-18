package router

import (
	"io/fs"
	"log"
	"net/http"

	"adamastor/internal/server/templates"
	"adamastor/internal/server/utilities"
	"adamastor/public"

	"github.com/go-playground/form/v4"
)

type Router struct {
	Mux *http.ServeMux
	Err error
}

func NewRouter() *Router {
	router := Router{
		Mux: http.NewServeMux(),
		Err: nil,
	}
	return &router
}

func (r *Router) SetStaticPaths() {
	static := []string{
		"/static/",
	}

	const base = "assets"
	staticFS, err := fs.Sub(public.Assets, base)
	if err != nil {
		r.Err = err
		return
	}
	for _, path := range static {
		r.Mux.Handle(path, http.FileServer(http.FS(staticFS)))
	}
}

func (r *Router) HandleRoutes() {
	r.SetStaticPaths()
	r.Mux.HandleFunc("/cv", HandleCV)
	r.Mux.HandleFunc("/projects", HandleProjects)
	r.Mux.HandleFunc("/littleprofessor", HandleLittleProfessor)
	r.Mux.HandleFunc("/start_professor", startProfessor)
	r.Mux.HandleFunc("/play_round", playRound)
	r.Mux.HandleFunc("/", HandleIndex)
}

//TODO: Check templ Render error and display 500 in case of error

func HandleCV(w http.ResponseWriter, r *http.Request) {
    templates.Layout(templates.CV()).Render(r.Context(), w)
}

func HandleProjects(w http.ResponseWriter, r *http.Request) {
    templates.Layout(templates.Project()).Render(r.Context(), w)
}

func HandleLittleProfessor(w http.ResponseWriter, r *http.Request) {
    templates.Professor().Render(r.Context(), w)
}

func startProfessor(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Panic(err)
    }
    decoder := form.NewDecoder()

	var game utilities.Game

	err = decoder.Decode(&game, r.Form)
	if err != nil {
		log.Panic(err)
	}

	game.Operations = []string{"+", "-", "*", "/"}
    game.MakeRound()
    templates.StartProfessor(game).Render(r.Context(), w)
}

func playRound(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        log.Panic(err)
    }
    decoder := form.NewDecoder()

	var game utilities.Game

	err = decoder.Decode(&game, r.Form)
	if err != nil {
		log.Panic(err)
	}
    if game.Answer == game.Result {
        game.Win += 1
    } else {
        game.Loss += 1
    }
    game.NumRounds -= 1
    log.Println(game)
    game.MakeRound()
    templates.PlayRound(game).Render(r.Context(), w)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    templates.Layout(templates.Index()).Render(r.Context(), w)
}
