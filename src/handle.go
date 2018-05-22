package main

import (
	"net/http"
	"path"
)

var handleMap = map[string]http.HandlerFunc{}

func root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "pages/index.html")
	} else {
		filenameWithSuffix := path.Base(r.URL.Path)
		fileSuffix := path.Ext(filenameWithSuffix)
		switch fileSuffix {
		case ".js":
			http.ServeFile(w, r, "actions/"+filenameWithSuffix)
		case ".css":
			http.ServeFile(w, r, "styles/"+filenameWithSuffix)
		case ".jpg", ".jpeg", ".gif", ".png", ".svg":
			http.ServeFile(w, r, "images/"+filenameWithSuffix)
		case ".eot", ".ttf", ".woff":
			http.ServeFile(w, r, "fonts/"+filenameWithSuffix)
		}
	}
}

func initHandle() {
	handleMap["/"] = root
	handleMap["/data"] = root
	for pattern, fn := range handleMap {
		http.HandleFunc(pattern, fn)
	}
}
