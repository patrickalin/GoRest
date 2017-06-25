package httpgo

import (
	"os"
	"strings"
	"testing"
)

var rest HTTP

func TestMain(m *testing.M) {

	rest = New(nil)
	os.Exit(m.Run())
}

func Test_restHTTP_GetWithHeaders(t *testing.T) {

	type args struct {
		url     string
		headers map[string][]string
	}
	m := make(map[string][]string)
	b := []string{"value"}
	m["key"] = b

	tests := []struct {
		name    string
		fields  HTTP
		args    args
		wantErr bool
	}{
		{"good google", rest, args{"http://www.google.com", m}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rest.Get(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("restHTTP.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restHTTP_GetError(t *testing.T) {
	rest := New(nil)
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  HTTP
		args    args
		wantErr string
	}{
		{"wrong google", rest, args{"http://www.googledsd.comZ"}, "no such host"},
		{"wrong google", rest, args{"http://www.services.alin.be/"}, "404"},
		{"wrong google", rest, args{"https://api.bloomsky.com/api/skydata/"}, "401"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rest.Get(tt.args.url); !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("restHTTP.Get() error = %v, wantErr %s", err, tt.wantErr)
			}
		})
	}
}

func Test_restHTTP_Get_GoodCase(t *testing.T) {
	a := New(nil)
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  HTTP
		args    args
		wantErr bool
	}{
		{"good google", a, args{"http://www.google.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := a.Get(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("restHTTP.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restHTTP_GetBody(t *testing.T) {
	perdu := New(nil)
	perdu.Get("http://www.perdu.com/")
	empty := New(nil)
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields HTTP
		want   string
	}{
		{"empty", empty, ""},
		{"good google", perdu, "<html><head><title>Vous Etes Perdu ?</title></head><body><h1>Perdu sur l'Internet ?</h1><h2>Pas de panique, on va vous aider</h2><strong><pre>    * <----- vous &ecirc;tes ici</pre></strong></body></html>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := string(tt.fields.GetBody()); strings.TrimSpace(got) != strings.TrimSpace(tt.want) {
				t.Errorf("restHTTP.GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restHTTP_PostJSON(t *testing.T) {
	a := New(nil)
	type args struct {
		url    string
		buffer []byte
	}
	tests := []struct {
		name    string
		fields  HTTP
		args    args
		wantErr string
	}{
		{"wrong google", a, args{"http://www.googledsd.comZ", []byte("foo1")}, "no such host"},
		{"wrong google", a, args{"http://www.services.alin.be/", []byte("foo2")}, "404"},
		{"wrong google", a, args{"https://api.bloomsky.com/api/skydata/", []byte("foo3")}, "405"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := a.PostJSON(tt.args.url, tt.args.buffer); !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("restHTTP.PostJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
