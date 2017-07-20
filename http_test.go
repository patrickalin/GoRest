package http

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
		{"wrong google", rest, args{"http://www.dhnet.be/perdu"}, "404"},
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
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  HTTP
		args    args
		wantErr bool
	}{
		{"good google", rest, args{"http://www.google.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rest.Get(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("restHTTP.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restHTTP_GetBody(t *testing.T) {
	err := rest.Get("http://www.perdu.com/")
	checkErr(err, funcName(), "error with Get")
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
		{"good google", rest, "<html><head><title>Vous Etes Perdu ?</title></head><body><h1>Perdu sur l'Internet ?</h1><h2>Pas de panique, on va vous aider</h2><strong><pre>    * <----- vous &ecirc;tes ici</pre></strong></body></html>"},
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
		{"wrong google", rest, args{"http://www.googledsd.comZ", []byte("foo1")}, "no such host"},
		{"wrong google", rest, args{"http://www.dhnet.be/perdu", []byte("foo2")}, "404"},
		{"wrong google", rest, args{"https://api.bloomsky.com/api/skydata/", []byte("foo3")}, "405"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rest.PostJSON(tt.args.url, tt.args.buffer); !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("restHTTP.PostJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restHTTP_initLog(t *testing.T) {
	ll := initLog(nil)
	New(ll)
}

/////// Benchmark ///////

func Benchmark_GetGoogle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := rest.Get("http://www.google.com")
		checkErr(err, funcName(), "error Get", "http://www.google.com")
	}
}
