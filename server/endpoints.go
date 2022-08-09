package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"github.com/Iffigues/mywebsite/server/endpoint"
	"text/template"
)

func jointure(r *http.Request, w http.ResponseWriter, are interface{}, ar ...string) {
	var joins []string
	for _, ok := range ar {
		joins = append(joins, filepath.Join("template", ok))
	}
	tmpl, err := template.ParseFiles(joins...)
	fmt.Println(err)
	tmpl.ExecuteTemplate(w, "layout", are)
}

func methode(a ...string) (b []string) {
	return a
}

func myRole(s *Server, r *http.Request) (role int) {
	session, err := s.Data.Store.Get(r, "user")
	if err != nil {
		fmt.Println(err)
		role = 1
	} else {
		if session.Values["role"] == nil {
			role = 1
		} else {
			role = session.Values["role"].(int)
		}
	}

	return
}

func (s *Server) Middleware(next http.Handler, a *Handle) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := myRole(s, r)
		fmt.Println(a)
		if a.Role <= role {
			println("Hello", a.Role, " ",role)
			next.ServeHTTP(w, r)
		} else {

			http.Redirect(w, r, "/", 303)
		}
	})
}

func (g *Server) makeEndPoint() {
	g.NewR("/", "home", methode("GET"), endpoint.Health(g.Data), 1)
	g.NewR("/gallerie", "gallerie", methode("GET"), endpoint.Health(g.Data), 1)
	g.NewR("/article", "article", methode("GET"), endpoint.Health(g.Data), 1)
}
