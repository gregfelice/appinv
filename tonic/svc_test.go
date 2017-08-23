package tonic

import (
	"bytes"
	"encoding/json"
	//"io"
	"net/http"
	//"os"
	dao "appinv/dao"
	"testing"
)

/*
https://dlintw.github.io/gobyexample/public/http-client.html
https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

*** >>> https://kev.inburke.com/kevin/golang-json-http/
**** https://blog.golang.org/json-and-go

@todo - write some tests that just show the encoding and unmarshalling stuff.

*/
func TestSVCApplications(t *testing.T) {

	appName := "Application ZeroTen"
	bizUnit := "Devops"

	t.Run("create", func(t *testing.T) {

		u := dao.Application{ApplicationName: appName, BusinessUnit: bizUnit}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(u)
		res, _ := http.Post("http://localhost:8080/applications", "application/json; charset=utf-8", b)
		//io.Copy(os.Stdout, res.Body)
		t.Log(res)
	})
}
