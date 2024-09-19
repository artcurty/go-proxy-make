package unit

import (
	"bytes"
	"github.com/artcurty/go-proxy-make/pkg"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDefaultHTTPClient_DoRequest_AllMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.Header.Get("Custom-Header") != "HeaderValue" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "GET success"}`))
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			r.Body.Close()
			if !bytes.Equal(body, []byte(`{"key":"value"}`)) {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if r.Header.Get("Content-Type") != "application/json" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "POST success"}`))
		case http.MethodPut:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "PUT success"}`))
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "DELETE success"}`))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer server.Close()

	type args struct {
		req pkg.HTTPRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name: "Successful GET request with header",
			args: args{
				req: pkg.HTTPRequest{
					Host:    server.URL,
					Path:    "/",
					Method:  http.MethodGet,
					Headers: map[string]string{"Custom-Header": "HeaderValue"},
				},
			},
			want: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "GET success"}`)),
			},
			wantErr: false,
		},
		{
			name: "Successful POST request with body and header",
			args: args{
				req: pkg.HTTPRequest{
					Host:    server.URL,
					Path:    "/",
					Method:  http.MethodPost,
					Body:    map[string]interface{}{"key": "value"},
					Headers: map[string]string{"Content-Type": "application/json"},
				},
			},
			want: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "POST success"}`)),
			},
			wantErr: false,
		},
		{
			name: "Successful PUT request",
			args: args{
				req: pkg.HTTPRequest{
					Host:   server.URL,
					Path:   "/",
					Method: http.MethodPut,
				},
			},
			want: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "PUT success"}`)),
			},
			wantErr: false,
		},
		{
			name: "Successful DELETE request",
			args: args{
				req: pkg.HTTPRequest{
					Host:   server.URL,
					Path:   "/",
					Method: http.MethodDelete,
				},
			},
			want: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"message": "DELETE success"}`)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &pkg.DefaultHTTPClient{}
			got, err := c.DoRequest(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			body, _ := ioutil.ReadAll(got.Body)
			got.Body.Close()
			wantBody, _ := ioutil.ReadAll(tt.want.Body)
			tt.want.Body.Close()
			if !reflect.DeepEqual(body, wantBody) {
				t.Errorf("DoRequest() got = %v, want %v", string(body), string(wantBody))
			}
		})
	}
}
