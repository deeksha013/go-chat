package main

import(
       "fmt"
       "flag"
         "net/http"
          //"github.com/gorilla/websocket"
        "text/template"
        "sync"
        "os"
	"path/filepath"
        )

        // templ represents a single template
type templateHandler struct {
      once     sync.Once
      filename string
      templ    *template.Template
      }
        // ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r  *http.Request) {
    t.once.Do(func() { // not matter how many gorous call this, this func() will only be called once for the current instance of "templateHandler"
    t.templ =  template.Must(template.ParseFiles(filepath.Join("templates",
    t.filename)))
      })
    t.templ.Execute(w, nil)
  }





func main(){


  var addr = flag.String("addr", ":8080", "The addr of the  application.")
  flag.Parse() // parse the flags
  r := newRoom()
  http.Handle("/", &templateHandler{filename: "chat.html"})
  http.Handle("/room", r)
  // get the room going./chat -addr=":
  go r.run()
  // start the web server
  fmt.Println(*addr)
  fmt.Println(os.Getenv("PORT"))
if err:=http.ListenAndServe(*addr,nil);err!=nil {
  panic("unable to start server at 8080")

}
}
