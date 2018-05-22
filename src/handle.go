package main

import (
	"log"
	"net/http"
	"path"
)

var handleMap = map[string]http.HandlerFunc{}

var d = []string{}

const (
	_ uint = iota
	ResType_Actions
	ResType_Styles
	ResType_Pages
	ResType_Images
	ResType_Fonts
)

func root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "pages/index.html")
	} else {
		filenameWithSuffix := path.Base(r.URL.Path)
		fileSuffix := path.Ext(filenameWithSuffix)
		rt, err := getResType(fileSuffix)
		if err != nil {
			log.Println(err)
		}
		dir, err := getResourceDir()
		if err != nil {
			log.Println(err)
		}
		dir = path.Clean(dir)
		switch rt {
		case ResType_Actions:
			http.ServeFile(w, r, path.Join(dir, "actions", filenameWithSuffix))
		case ResType_Styles:
			http.ServeFile(w, r, path.Join(dir, "styles", filenameWithSuffix))
		case ResType_Images:
			http.ServeFile(w, r, path.Join(dir, "images", filenameWithSuffix))
		case ResType_Fonts:
			http.ServeFile(w, r, path.Join(dir, "fonts", filenameWithSuffix))
		case ResType_Pages:
			http.ServeFile(w, r, path.Join(dir, "pages", filenameWithSuffix))
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
