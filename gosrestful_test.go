package gosrestful

import (
	"testing"
	"fmt"
	//"net/http"
)

type GoSRestful struct {
}

func (gorestful *GoSRestful) Create(goSRestfulPara *GoSRestfulPara) {

	fmt.Fprintf(*goSRestfulPara.w, "Create id:" + goSRestfulPara.para["id"])
}
func (gorestful *GoSRestful) Delete(goSRestfulPara *GoSRestfulPara) {
	fmt.Fprintf(*goSRestfulPara.w, "Delete id:" + goSRestfulPara.para["id"])
}
func (gorestful *GoSRestful) Update(goSRestfulPara *GoSRestfulPara) {
	fmt.Fprintf(*goSRestfulPara.w, "Update id:" + goSRestfulPara.para["id"])
}
func (gorestful *GoSRestful) Get(goSRestfulPara *GoSRestfulPara) {
	fmt.Fprintf(*goSRestfulPara.w, "Get id:" + goSRestfulPara.para["id"])
}

func TestCreate(t *testing.T) {
	gsrrun := new(GoSRestfulRun)
	gsr := new(GoSRestful)
	gsrrun.AddURIProcess("/{id}/{test}/", gsr)
	gsrrun.ListenAndServe("localhost:8080")
}
