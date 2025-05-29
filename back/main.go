package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path"

	_ "web_ilya/docs"
)

//	@title			Rosinka admin panel
//	@version		1.0
//	@description	This is a proxy server that serves data from DWH.

//	@license.name	Rosinka 1.0

//	@host		localhost:8080

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

//go:embed static/* templates/*
var content embed.FS

func main() {
	fmt.Println("running server... ")

	staticFiles, _ := fs.Sub(content, "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

	http.Handle("/index", authMiddleware(http.HandlerFunc(renderHTML("index.html"))))
	http.Handle("/catalog", authMiddleware(http.HandlerFunc(renderHTML("catalog.html"))))
	http.Handle("/contacts", authMiddleware(http.HandlerFunc(renderHTML("contacts.html"))))
	http.Handle("/about", authMiddleware(http.HandlerFunc(renderHTML("about.html"))))
	http.Handle("/guestbook", authMiddleware(http.HandlerFunc(renderHTML("guestbook.html"))))

	http.Handle("/cinema_item", authMiddleware(http.HandlerFunc(search)))
	http.Handle("/food_item", authMiddleware(http.HandlerFunc(search)))
	http.Handle("/kids_item", authMiddleware(http.HandlerFunc(search)))
	http.Handle("/shops_item", authMiddleware(http.HandlerFunc(search)))

	/*
		http.Handle("/offer", authMiddleware(http.HandlerFunc(renderHTML("offer.html"))))
		http.Handle("/production", authMiddleware(http.HandlerFunc(renderHTML("production.html"))))
	*/
	http.Handle("/search", authMiddleware(http.HandlerFunc(search)))

	//Auth
	http.HandleFunc("/reg", routereg)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)

	http.ListenAndServe(":8080", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	var err error
	query := r.URL.Query().Get("query")
	urlPath := r.URL.Path
	fmt.Printf("path: %s\n", urlPath)
	var data []interface{}
	switch urlPath {
	case "/shops_item":
		result, err := PGFindShops(query)
		if err != nil {
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		for _, item := range result {
			data = append(data, item)
		}
	case "/food_item":
		result, err := PGFindFoodItems(query)
		if err != nil {
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		for _, item := range result {
			data = append(data, item)
		}
	case "/kids_item":
		result, err := PGFindKidsZones(query)
		if err != nil {
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		for _, item := range result {
			data = append(data, item)
		}
	case "/cinema_item":
		result, err := PGFindCinemas(query)
		if err != nil {
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		for _, item := range result {
			data = append(data, item)
		}
	}

	/*
		if query == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
	*/
	/*
		funcMap := template.FuncMap{
			"formatDate": func(t time.Time) string {
				return t.Format("02 January 2006") // Формат DD Month YYYY
			},
		}
	*/
	//tmpl := template.Must(template.New(urlPath+".html").Funcs(funcMap).ParseFS(content, path.Join("templates", "pages", urlPath+".html")))
	fmt.Printf("path: %s \n", path.Join("templates", "pages", urlPath+".html"))
	tmpl, err := template.ParseFS(content, path.Join("templates", "pages", urlPath+".html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "error search ParseFS", http.StatusInternalServerError)
		return
	}
	uname := r.Context().Value(key("username"))

	tmpl.Execute(w, map[string]any{
		"User": uname,
		"Data": data,
	})
}
