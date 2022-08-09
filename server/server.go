package server

import (
	"net/http"
	"os"
	"github.com/Iffigues/mywebsite/config"
	"strings"

	"github.com/Iffigues/mywebsite/types"
	"github.com/Iffigues/mywebsite/linuxtool"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type HH interface {
	WWW(*Server)
}

type Server struct {
	Router *mux.Router
	Data   *types.Data
	Handle map[string]*Handle
	Give   []HH
}

type Handle struct {
	Role   int
	Route  string
	Method []string
	Handle http.Handler
}

func (s *Server) AddHH(p ...HH) {
	for _, val := range p {
		s.Give = append(s.Give, val)
	}
}

func (s *Server) StartHH() {
	for _, val := range s.Give {
		val.WWW(s)
	}
}

func NewServer(conf config.Config) *Server {
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return &Server{
		Data: &types.Data{
			Store: sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY"))),
			Conf:  conf,
			Commande: linuxtool.NewCommand(),
			//Db:    pk.NewPk(conf["pk"]),
		},
		Router: router,
		Handle: make(map[string]*Handle),
	}
}

func (r *Server) NewR(route, key string, method []string, handler http.Handler, i int) {
	route = strings.ToLower(route)
	r.Handle[key] = &Handle{Method: method, Route: route, Handle: handler, Role: i}
}

func (g *Server) Servers(conf config.Config) (srv *http.Server) {
	g.makeEndPoint()
	g.StartHH()
	for _, h := range g.Handle {
		g.Router.Handle(h.Route, g.Middleware(h.Handle, h)).Methods(h.Method...)
	}
	return &http.Server{
		Addr:    ":80",
		Handler: g.Router,
	}
}
