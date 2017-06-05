package myRest

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_restHTTP_GetWithHeaders(t *testing.T) {
	type fields struct {
		status string
		header http.Header
		body   []byte
	}
	type args struct {
		url     string
		headers map[string][]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &restHTTP{
				status: tt.fields.status,
				header: tt.fields.header,
				body:   tt.fields.body,
			}
			if err := r.GetWithHeaders(tt.args.url, tt.args.headers); (err != nil) != tt.wantErr {
				t.Errorf("restHTTP.GetWithHeaders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restHTTP_Get(t *testing.T) {
	type fields struct {
		status string
		header http.Header
		body   []byte
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.

	//tt.args.url = "www.google.com"
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &restHTTP{
				status: tt.fields.status,
				header: tt.fields.header,
				body:   tt.fields.body,
			}
			if err := r.Get(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("restHTTP.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_restHTTP_GetBody(t *testing.T) {
	type fields struct {
		status string
		header http.Header
		body   []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &restHTTP{
				status: tt.fields.status,
				header: tt.fields.header,
				body:   tt.fields.body,
			}
			if got := r.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restHTTP.GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
