package swagger

import (
	"bytes"
	"html/template"
	"net/http"
)

var DefaultDocs = WrappedDocsHandler(DocsOpts{})
var DefaultReDoc = WrappedReDocHandler(RedocOpts{})

func WrappedDocsHandler(opts DocsOpts) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		opts.EnsureDefaults()

		tmpl := template.Must(template.New("docs").Parse(docsTemplate))

		buf := bytes.NewBuffer(nil)
		_ = tmpl.Execute(buf, opts)
		b := buf.Bytes()

		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.WriteHeader(http.StatusOK)

		_, _ = rw.Write(b)
		return
	}
}

func WrappedReDocHandler(opts RedocOpts) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		opts.EnsureDefaults()

		tmpl := template.Must(template.New("redoc").Parse(redocTemplate))

		buf := bytes.NewBuffer(nil)
		_ = tmpl.Execute(buf, opts)
		b := buf.Bytes()

		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.WriteHeader(http.StatusOK)

		_, _ = rw.Write(b)
		return
	}
}
