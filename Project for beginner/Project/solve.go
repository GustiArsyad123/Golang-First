package main

import (
	"expvar"

	"encoding/json"

	"bytes"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"go/build"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os/exec"
)

type myType struct {
	J []json.RawMessage
}

var pack map[string]string

type GoList struct {
	Imports []string
}

type Import struct {
	Dir        string
	ImportPath string
	Name       string
	Target     string
	Standard   bool
	Root       string
	GoFiles    []string
	Imports    []string
	Deps       []string
}

const contentTypeJSON = "application/json"

func main() {

	http.HandleFunc("/importgraph", func(w http.ResponseWriter, r *http.Request) { importGraph(w, r) })
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside handler")
	fmt.Fprintf(w, "Hello world from my Go program!")
}

func importGraph(w http.ResponseWriter, r *http.Request) {

	pack = make(map[string]string)

	var t myType
	cmd := exec.Command("go", "list", "-json")
	stdout, err := cmd.Output()
	if err != nil {

		println(err.Error())
		return
	}

	var list GoList
	err = json.Unmarshal(stdout, &list)

	for _, d := range list.Imports {
		//get the imports for each of the packages listed by go list -json
		t.imports(d)

	}

	var buff bytes.Buffer

	//concatenate the separate json.RawMessages together into json

	buff.WriteByte('[')

	for i, j := range t.J {

		if i != 0 {
			buff.WriteByte(',')
		}
		buff.Write([]byte(j))
	}
	buff.WriteByte(']')

	var buffer bytes.Buffer
	if err := json.Compact(&buffer, buff.Bytes()); err != nil {
		println(err.Error()) //error message: invalid character ',' looking for beginning of value
		return

	}

	w.Header().Set("Content-Type", contentTypeJSON)

	w.Write(buffer.Bytes())

}

func (myObj *myType) imports(pk string) error {

	cmd := exec.Command("go", "list", "-json", pk)
	stdout, _ := cmd.Output()

	pack[pk] = pk

	var deplist Import
	json.Unmarshal(stdout, &deplist)

	var newj json.RawMessage
	json.Unmarshal(stdout, &newj)
	myObj.J = append(myObj.J, newj)

	for _, imp := range deplist.Imports {

		if _, ok := pack[imp]; !ok {

			myObj.imports(imp) //recursive call to get the imports of the imports etc

		}
	}

	return nil

}
