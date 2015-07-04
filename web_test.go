package main
import (
	"net/http"
	"testing"
	"strings"
	"net/http/httptest"
)

func TestAddContacts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(addressBookHandler))
	defer ts.Close()

	resp, _ := http.Post(ts.URL+"/contacts","application/json",strings.NewReader(`{
			"name" : "hello",
			"tel" : "0234234234",
			"email" : "h@mail.com"
		}`))
	if resp.StatusCode != http.StatusOK {
		t.Error("Expect status code is 200 OK")
	}
}
