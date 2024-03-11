package server

import (
	"log"
	"net/http"

	"github.com/JonnyLoughlin/pio/internal/ui"
	"github.com/JonnyLoughlin/pio/internal/ui/src/templates"
	"github.com/JonnyLoughlin/pio/internal/ui/src/templates/components"
	"github.com/JonnyLoughlin/pio/internal/ui/src/templates/pages"

	"github.com/a-h/templ"

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
		r.Get(string(RouteHome), s.HomeHandler)
		r.Get(string(RouteServices), s.ServicesHandler)
		r.Get(string(RouteCatering), s.CateringHandler)
		r.Get(string(RouteUnfriendlys), s.UnfriendlysHandler)
		r.Get(string(RouteEmployment), s.EmploymentHandler)
		r.Get(string(RouteContact), s.ContactHandler)
		r.Get(string(RouteOrder), s.OrderHandler)
	})

	return r
}

type Route string

const (
	RouteHome        Route = "/"
	RouteServices    Route = "/Services"
	RouteCatering    Route = "/Catering"
	RouteUnfriendlys Route = "/Unfriendlys"
	RouteEmployment  Route = "/Employment"
	RouteContact     Route = "/Contact"
	RouteOrder       Route = "/Order"
)

var TabsData = []components.TabProps{
	{
		Text:  "Home",
		HxGet: string(RouteHome),
	},
	{
		Text:  "Services",
		HxGet: string(RouteServices),
	},
	{
		Text:  "Catering",
		HxGet: string(RouteCatering),
	},
	{
		Text:  "Unfriendly's Ice Cream",
		HxGet: string(RouteUnfriendlys),
	},
	{
		Text:  "Employment",
		HxGet: string(RouteEmployment),
	},
	{
		Text:  "Contact Us",
		HxGet: string(RouteContact),
	},
	{
		Text:  "Order Here",
		HxGet: string(RouteOrder),
	},
}

func (s *Server) PageHandler(w http.ResponseWriter, r *http.Request, route Route, page templ.Component) {
	// If the request is an htmx request, write the maincontent only
	if htmx.IsHTMX(r) {
		log.Println("Is HTMX")
		err := htmx.NewResponse().PushURL(string(route)).Write(w)
		if err != nil {
			panic(err)
		}
		// else, write the content in the base layout
		err = page.Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
		return
	}
	err := templates.Base(TabsData, page).Render(r.Context(), w)
	if err != nil {
		panic(err)
	}
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, RouteHome, pages.Home())
}

func (s *Server) ServicesHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, RouteServices, pages.Services())
}

func (s *Server) CateringHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, RouteCatering, pages.Catering())
}

func (s *Server) UnfriendlysHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, RouteUnfriendlys, pages.Unfriendlys())
}

func (s *Server) EmploymentHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, RouteEmployment, pages.Employment())
}

func (s *Server) ContactHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, RouteContact, pages.Contact())
}

func (s *Server) OrderHandler(w http.ResponseWriter, r *http.Request) {
	s.PageHandler(w, r, RouteOrder, pages.Order())
}
