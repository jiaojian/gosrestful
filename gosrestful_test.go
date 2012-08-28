package gosrestful

import (
	"testing"
	"fmt"
	//"net/http"
)

type GoSRestful struct {
}

func (gorestful *GoSRestful) Create(goSRestfulPara *GoSRestfulPara) {

	fmt.Fprintf(goSRestfulPara.W, "Create id:" + goSRestfulPara.Para["id"])
}
func (gorestful *GoSRestful) Delete(goSRestfulPara *GoSRestfulPara) {
	fmt.Fprintf(goSRestfulPara.W, "Delete id:" + goSRestfulPara.Para["id"])
}
func (gorestful *GoSRestful) Update(goSRestfulPara *GoSRestfulPara) {
	fmt.Fprintf(goSRestfulPara.W, "Update id:" + goSRestfulPara.Para["id"])
}
func (gorestful *GoSRestful) Get(goSRestfulPara *GoSRestfulPara) {
	fmt.Fprintf(goSRestfulPara.W, "Get id:" + goSRestfulPara.Para["id"])
}

func TestCreate(t *testing.T) {
	gsrrun := new(GoSRestfulRun)
	gsr := new(GoSRestful)
	gsrrun.AddURIProcess("index", "/{id}/{test}/", gsr)
	gsrrun.ListenAndServe("localhost:8080")
}
