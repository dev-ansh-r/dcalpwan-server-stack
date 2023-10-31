package web

import (
	"net/http"
)

type redirector struct {
	Path     string
	Code     int
	Location string
}

func (red redirector) RegisterRoutes(s *Server) {
	s.RootRouter().Path(red.Path).Handler(http.RedirectHandler(red.Location, red.Code))
}

// Redirect returns a Registerer that redirects requests to the given path to
// the given location with the given code.
func Redirect(path string, code int, location string) Registerer {
	return redirector{
		Path:     path,
		Code:     code,
		Location: location,
	}
}
