package ui

import (
	"fmt"
	"io"
	"text/template"
)

type (
	Layout string
	Page   string
)

const (
	LayoutBase Layout = "Base"
	LayoutHtmx Layout = "Htmx"
)

const (
	PageHome        Page = "Home"
	PageServices    Page = "Services"
	PageCatering    Page = "Catering"
	PageUnfriendlys Page = "Unfriendlys"
	PageEmployment  Page = "Employment"
	PageContact     Page = "Contact"
	PageOrder       Page = "Order"
)

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

type tabProps struct {
	Text  string
	HxGet string
}

var tabsData = []tabProps{
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

type templateData struct {
	TabsData    []tabProps
	ContactForm bool
}

func WriteComposedTemplate(w io.Writer, layout Layout, page Page) (*template.Template, error) {
	var data templateData

	tmpl, err := buildBaseTemplate(&data, layout)
	if err != nil {
		return nil, err
	}

	_, err = buildPageTemplate(tmpl, page)
	if err != nil {
		return nil, err
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func buildBaseTemplate(data *templateData, layout Layout) (*template.Template, error) {
	var tmpl *template.Template
	var err error

	switch layout {
	// For the base layout, add the template and set the tabs data
	case LayoutBase:
		tmpl, err = template.ParseFS(HtmlFiles, fmt.Sprintf(`src/html/layouts/%s.gohtml`, LayoutBase))
		if err != nil {
			return nil, err
		}
		// Data for executing the tabs needs to be added for the base template
		data.TabsData = tabsData
		_, err = tmpl.ParseFS(HtmlFiles, "src/html/components/tabs.gohtml")

	// For the htmx layout, just add the main content
	case LayoutHtmx:
		tmpl, err = template.ParseFS(HtmlFiles, fmt.Sprintf(`src/html/layouts/%s.gohtml`, LayoutHtmx))

	}

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func buildPageTemplate(tmpl *template.Template, page Page) (*template.Template, error) {
	var err error
	_, err = tmpl.ParseFS(HtmlFiles, fmt.Sprintf(`src/html/pages/%s.gohtml`, page))
	if err != nil {
		return nil, err
	}

	if page == PageContact || page == PageEmployment {
		tmpl, err = tmpl.ParseFS(HtmlFiles, "src/html/components/contact-form.gohtml")
		if err != nil {
			return nil, err
		}
	}

	return tmpl, nil
}
