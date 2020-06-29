package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	error := http.ListenAndServe(":8000", nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}

	/*	for i, c := range words {
		if string(c) != '' {
			fmt.Println(i, string(c))
		}
	*/
}

func index(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "Hello.")
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	froll := r.FormValue("rollno")
	fmsg := r.FormValue("msg")

	d := struct {
		Rollno  string
		Message string
	}{
		Rollno:  froll,
		Message: fmsg,
	}

	tpl.ExecuteTemplate(w, "processor.html", d)
	fmt.Println(fmsg)

	j := 0
	count := 0
	words := strings.Fields(fmsg)
	for i, w := range words {
		for range w {
			j++
		}
		fmt.Println(i, w)
		count++
	}

	fmt.Println("Total words are: ", count)
	fmt.Println("Total characters are: ", j)

	t := struct {
		Nword int
		Nchar int
	}{
		Nword: count,
		Nchar: j,
	}
	tpl.ExecuteTemplate(w, "processor.html", t)
}
