package gosrestful

import (
	"net/http"
	"regexp"
	"strings"
)

type GoSRestfulProcesser interface {
	Create(goSRestfulPara *GoSRestfulPara)
	Delete(goSRestfulPara *GoSRestfulPara)
	Update(goSRestfulPara *GoSRestfulPara)
	Get(goSRestfulPara *GoSRestfulPara)
	//Create(w http.ResponseWriter, r *http.Request)
	//Delete(w http.ResponseWriter, r *http.Request)
	//Update(w http.ResponseWriter, r *http.Request)
	//Get(w http.ResponseWriter, r *http.Request)
}

type GoSRestfulRun struct {
	URIMatchs map[string] string
	URIs map[string] GoSRestfulProcesser
}

type GoSRestfulPara struct {
	w *http.ResponseWriter
	r *http.Request
	para map[string] string
}

func (gosrestfulrun *GoSRestfulRun) ListenAndServe(HostAndPort string) {

	http.Handle("/", gosrestfulrun)
	http.ListenAndServe(HostAndPort, nil)
}

func (gosrestfulrun *GoSRestfulRun) AddURIProcess(URI string, process GoSRestfulProcesser) {
	
	//If first set URI's process, create URI map
	if gosrestfulrun.URIs == nil {
		gosrestfulrun.URIMatchs = make(map[string] string)
		gosrestfulrun.URIs = make(map[string] GoSRestfulProcesser)
	}

	//Set URI's process
	gosrestfulrun.URIs[URI] = process
	gosrestfulrun.URIMatchs[URI] = "^" + regexp.MustCompile(`\{[a-zA-Z0-9_-]+\}`).ReplaceAllString(URI, "[a-zA-Z0-9_-]+") + "$"
}

func (gosrestfulrun *GoSRestfulRun) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	for k, reg := range gosrestfulrun.URIMatchs {
	
		if regexp.MustCompile(reg).MatchString(r.URL.Path) {

			var requestPara GoSRestfulPara
			requestPara.w = &w
			requestPara.r = r
			requestPara.para = make(map[string] string)
			values := strings.Split(r.URL.Path, "/")
			keys := strings.Split(regexp.MustCompile(`[{|}]`).ReplaceAllString(k, ""), "/")
			matchKeys := strings.Split(reg, "/")
			for idx, v := range matchKeys {
				if v == "[a-zA-Z0-9_-]+" {
					requestPara.para[keys[idx]] = values[idx]
				}
			}
			switch r.Method {
			case "GET":
				gosrestfulrun.URIs[k].Get(&requestPara)
			case "DELETE":
				gosrestfulrun.URIs[k].Delete(&requestPara)
			case "PUT":
				gosrestfulrun.URIs[k].Create(&requestPara)
			case "POST":
				gosrestfulrun.URIs[k].Update(&requestPara)
			}
		}
	}
}
