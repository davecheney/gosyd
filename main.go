package main

import (
	"log"
	"net"
	"net/http"
	"text/template"
)

var (
	greeting = template.Must(template.New("greeting").Parse(`
<html>
   <head>
      <title>Hello {{.Name}}</title>
   </head>
   <body>
      <p>Hola {{.Name}}, ¿cómo estás?</p>
   </body>
</html>
`))

	question = template.Must(template.New("question").Parse(`
<html>
   <head>
      <title>Securebot 9000</title>
   </head>
   <body>
      <p>Halt! Who goes there? <form method="POST"><input type="text" name="Name" /></form></p>
   </body>
</html>
`))
)

// handle /favicon.ico requests
func favicon(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Not found", 404)
}

// greeting handler
func greeter(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// send question
		question.Execute(w, nil)
		return
	case "POST":
		if err := req.ParseForm(); err != nil {
			http.Error(w, err.Error(), 504)
			return
		}
		// send greeting
		greeting.Execute(w, req.Form)
		return
	default:
		http.Error(w, "Method not allowed", 405)
	}
}

func main() {
	// setup handlers
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/", greeter)

	l, err := net.Listen("tcp4", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("http server listening on %v", l.Addr())

	if err := http.Serve(l, nil); err != nil {
		log.Fatal(err)
	}
}
