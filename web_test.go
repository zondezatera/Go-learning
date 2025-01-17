package main
import (
	"contact"
	"io/ioutil"
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
			"tel" : "001112222",
			"email" : "h@mail.com"
		}`))
	if resp.StatusCode != http.StatusOK {
		t.Error("Expect status code is 200 OK")
	}
	if len(contacts) != 1 {
		t.Error("Expect contacts has 1 contacts")
	}
	c := contacts[0]
	if c.Name != "hello"  {
		t.Error("Expect name is hello but got",c.Name)
	}
	if c.Tel != "001112222"  {
		t.Error("Expect tel is 001112222 but got",c.Tel)
	}
	if c.Email != "h@mail.com"  {
		t.Error("Expect email is h@mail.com but got",c.Email)
	}
}

func TestGetContact(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(addressBookHandler))
	defer ts.Close()

	contacts = []contact.Contact{
		contact.Contact{
			Name: "wingyplus",
			Tel: "00001112222",
			Email: "h@mail.com",
		},
	}

	resp , _ := http.Get(ts.URL + "/contacts?name=wingyplus")
	if resp.StatusCode != http.StatusOK {
		t.Error("Expect status is 200")
	}

	b , _ := ioutil.ReadAll(resp.Body)
	if string(b) != `{"name":"wingyplus","tel":"00001112222","email":"h@mail.com"}`+"\n" {
		t.Error(`Expect json {"name":"hello","tel":"32423423","email" : "test@test.com"}`)
	}

	resp, _ = http.Get(ts.URL + "/contacts?name=haha")
	if resp.StatusCode != http.StatusNotFound {
		t.Error("Expect status code is 404")
	}
}