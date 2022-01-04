package swagger

// RedocOpts configures the Redoc middlewares
type RedocOpts struct {
	// SpecSwaggerJsonURL the url to find the spec for
	SpecSwaggerJsonURL string
	// RedocJsURL for the js that generates the redoc site, defaults to: https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js
	RedocJsURL string
	// Title for the documentation site, default to: API documentation
	RedocCssURL string
	IconURL     string
	Title       string
}

type DocsOpts struct {
	DocsSwaggerJsonURL string
	DocsJsURL          string
	DocsCssURL         string
	IconURL            string
	Title              string
}

func (d *DocsOpts) EnsureDefaults() {
	if d.DocsSwaggerJsonURL == "" {
		d.DocsSwaggerJsonURL = "/static/swagger.json" // 此处的swagger.json需要自行提供
	}
	if d.DocsJsURL == "" {
		d.DocsJsURL = "https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui-bundle.js"
	}
	if d.DocsCssURL == "" {
		d.DocsCssURL = "https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui.css"
	}
	if d.IconURL == "" {
		d.IconURL = "https://fastapi.tiangolo.com/img/favicon.png"
	}
	if d.Title == "" {
		d.Title = "API documentation"
	}
}

// EnsureDefaults in case some options are missing
func (r *RedocOpts) EnsureDefaults() {
	if r.SpecSwaggerJsonURL == "" {
		r.SpecSwaggerJsonURL = "/static/swagger.json"
	}
	if r.RedocJsURL == "" {
		r.RedocJsURL = redocJs
	}
	if r.RedocCssURL == "" {
		r.RedocCssURL = redocCss
	}
	if r.IconURL == "" {
		r.IconURL = "https://fastapi.tiangolo.com/img/favicon.png"
	}
	if r.Title == "" {
		r.Title = "API documentation"
	}
}
