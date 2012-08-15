package gosrestful

import (
	"net/http"
	"regexp"
	"strings"
	"html/template"
)

/*
 *Resource process interface
 *We promise one resource must have CRUD process
 * **I think it's right, if any resource lost one or more process
 * **may be you forgot it
 */
type GoSRestfulProcesser interface {
	Create(goSRestfulPara *GoSRestfulPara)
	Delete(goSRestfulPara *GoSRestfulPara)
	Update(goSRestfulPara *GoSRestfulPara)
	Get(goSRestfulPara *GoSRestfulPara)
}

type GoSRestfulRun struct {
	URIIds map[string] string
	URIMatchs map[string] string
	URIs map[string] GoSRestfulProcesser
}

type GoSRestfulPara struct {
	W http.ResponseWriter
	R *http.Request
	M string
	Para map[string] string
	ResultData interface{}
	Redirect string
}

type ErrorReport struct {
	Title string
	Id string
}

func (gosrestfulrun *GoSRestfulRun) ListenAndServe(HostAndPort string) {

	http.Handle("/", gosrestfulrun)
	http.ListenAndServe(HostAndPort, nil)
}

func (gosrestfulrun *GoSRestfulRun) AddURIProcess(id string, URI string, process GoSRestfulProcesser) {
	
	//If first set URI's process, create URI map
	if gosrestfulrun.URIs == nil {
		gosrestfulrun.URIIds = make(map[string] string)
		gosrestfulrun.URIMatchs = make(map[string] string)
		gosrestfulrun.URIs = make(map[string] GoSRestfulProcesser)
	}

	//Set URI's process
	gosrestfulrun.URIIds[URI] = id
	gosrestfulrun.URIs[URI] = process
	gosrestfulrun.URIMatchs[URI] = "^" + regexp.MustCompile(`\{[a-zA-Z0-9_-]+\}`).ReplaceAllString(URI, "[a-zA-Z0-9_-]+") + "$"
}

func (gosrestfulrun *GoSRestfulRun) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	for k, reg := range gosrestfulrun.URIMatchs {
	
		if regexp.MustCompile(reg).MatchString(r.URL.Path) {

			var requestPara GoSRestfulPara
			requestPara.W = w
			requestPara.R = r
			requestPara.Para = make(map[string] string)
			values := strings.Split(r.URL.Path, "/")
			keys := strings.Split(regexp.MustCompile(`[{|}]`).ReplaceAllString(k, ""), "/")
			matchKeys := strings.Split(reg, "/")
			for idx, v := range matchKeys {
				if v == "[a-zA-Z0-9_-]+" {
					requestPara.Para[keys[idx]] = values[idx]
				}
			}
			requestPara.Redirect = gosrestfulrun.URIIds[k]
			switch r.Method {
			case "GET":
				requestPara.M = "R"
				gosrestfulrun.URIs[k].Get(&requestPara)
			case "DELETE":
				requestPara.M = "D"
				gosrestfulrun.URIs[k].Delete(&requestPara)
			case "PUT":
				requestPara.M = "C"
				gosrestfulrun.URIs[k].Create(&requestPara)
			case "POST":
				requestPara.M = "U"
				gosrestfulrun.URIs[k].Update(&requestPara)
			}
			gosrestfulrun.redirect(&requestPara)
		}
	}
}

func (gosrestfulrun *GoSRestfulRun) redirect(requestPara *GoSRestfulPara) {
	if requestPara.ResultData == nil {
		requestPara.ResultData = &ErrorReport {
			Title : "404",
			Id : "123",
		}
		requestPara.Redirect = "404"
	}
	t, err := template.ParseFiles("./view/" + requestPara.Redirect + requestPara.M + ".tpl")
	if err != nil {
		http.Error(requestPara.W, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(requestPara.W, requestPara.ResultData)
	if err != nil {
		http.Error(requestPara.W, err.Error(), http.StatusInternalServerError)
	}
}
