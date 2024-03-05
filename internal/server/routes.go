package server

import (
	"net/http"

	"github.com/JonnyLoughlin/pio/internal/ui"
	"github.com/angelofallars/htmx-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Configure CORS Handler
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Private Routes
	r.Group(func(r chi.Router) {
		// Handle Serving JS
		jsFileServer := http.FileServer(http.FS(ui.JsFiles))
		r.Handle("/src/js/*", jsFileServer)
		// Handle Serving CSS
		cssFileServer := http.FileServer(http.FS(ui.CssFiles))
		r.Handle("/src/css/*", cssFileServer)
		// Handle Serving Static Files
		assetFileServer := http.FileServer(http.FS(ui.AssetFiles))
		r.Handle("/src/assets/*", assetFileServer)

		// Page Routes
		r.Get(string(ui.RouteHome), s.HomeHandler)
		r.Get(string(ui.RouteServices), s.ServicesHandler)
		r.Get(string(ui.RouteCatering), s.CateringHandler)
		r.Get(string(ui.RouteUnfriendlys), s.UnfriendlysHandler)
		r.Get(string(ui.RouteEmployment), s.EmploymentHandler)
		r.Get(string(ui.RouteContact), s.ContactHandler)
		r.Get(string(ui.RouteOrder), s.OrderHandler)
	})

	return r
}

func (s *Server) PageHandler(w http.ResponseWriter, r *http.Request, route ui.Route, page ui.Page) {
	// If the request is an htmx request, write the maincontent only
	if htmx.IsHTMX(r) {
		err := htmx.NewResponse().PushURL(string(route)).Write(w)
		if err != nil {
			panic(err)
		}
		_, err = ui.WriteComposedTemplate(w, ui.LayoutHtmx, page)
		if err != nil {
			panic(err)
		}
		return
	}
	// else, write the content in the base layout
	_, err := ui.WriteComposedTemplate(w, ui.LayoutBase, ui.PageHome)
	if err != nil {
		panic(err)
	}
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, ui.RouteHome, ui.PageHome)
}

func (s *Server) ServicesHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, ui.RouteServices, ui.PageServices)
}

func (s *Server) CateringHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, ui.RouteCatering, ui.PageCatering)
}

func (s *Server) UnfriendlysHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, ui.RouteUnfriendlys, ui.PageUnfriendlys)
}

func (s *Server) EmploymentHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, ui.RouteEmployment, ui.PageEmployment)
}

func (s *Server) ContactHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, ui.RouteContact, ui.PageContact)
}

func (s *Server) OrderHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, ui.RouteOrder, ui.PageOrder)
}
