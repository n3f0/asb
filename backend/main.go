package main
import "net/http"
func main() {
  //ini adalah listener http dengan memanfaatkan libarary http
  http.ListenAnServe(":80",nil)
}
