package casdoor

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/containous/traefik/v2/pkg/rules"
)

type Server struct {
	router *rules.Router
}

func NewServer(c Config) *Server {
	s := &Server{}
	var err error
	s.router, err = rules.NewRouter()
	if err != nil {
		panic(err)
	}

	for name, rule := range c.Rules {
		if rule.Action == "allow" {
			err = s.router.AddRoute(rule.Rule, 1, s.AllowHandler(name))
		} else {
			err = s.router.AddRoute(rule.Rule, 1, s.AuthHandler(rule.Application, name))
		}
	}

	s.router.NewRoute().Handler(s.AuthHandler(c.Application, "default"))

	return s
}

func (s *Server) AllowHandler(rule string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s%s, Allow, %s\n", r.Host, r.RequestURI, rule)
		w.WriteHeader(200)
	}
}

func (s *Server) AuthHandler(providerName, rule string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s%s, Deny, %s\n", r.Host, r.RequestURI, rule)
		w.WriteHeader(401)
	}
}


func (s *Server) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	r.Method = r.Header.Get("X-Forwarded-Method")
	r.Host = r.Header.Get("X-Forwarded-Host")

	if _, ok := r.Header["X-Forwarded-Uri"]; ok {
		r.URL, _ = url.Parse(r.Header.Get("X-Forwarded-Uri"))
	}

	s.router.ServeHTTP(w, r)
}
