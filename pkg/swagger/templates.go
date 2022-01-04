package swagger

const (
	redocJs       = "https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js"
	redocCss      = "https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700"
	redocTemplate = `<!DOCTYPE html>
<html>
  <head>
	<link rel="shortcut icon" href="{{ .IconURL }}">
    <title>{{ .Title }}</title>
		<!-- needed for adaptive design -->
		<meta charset="utf-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link href="{{ .RedocCssURL }}" rel="stylesheet">

    <!--
    ReDoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='{{ .SpecSwaggerJsonURL }}'></redoc>
    <script src="{{ .RedocJsURL }}"> </script>
  </body>
</html>
`
	docsTemplate = `<!DOCTYPE html>
<html>
<head>
<link type="text/css" rel="stylesheet" href="{{ .DocsCssURL }}">
<link rel="shortcut icon" href="{{ .IconURL }}">
<title>{{ .Title }} - Swagger UI</title>
</head>
<body>
<div id="swagger-ui">
</div>
<script src="{{ .DocsJsURL }}"></script>
<script>
const ui = SwaggerUIBundle({
    url: '{{ .DocsSwaggerJsonURL }}',
	oauth2RedirectUrl: window.location.origin + '/docs/oauth2-redirect',
    dom_id: '#swagger-ui',
    presets: [
    SwaggerUIBundle.presets.apis,
    SwaggerUIBundle.SwaggerUIStandalonePreset
    ],
    layout: "BaseLayout",
    deepLinking: true,
    showExtensions: true,
    showCommonExtensions: true
})
</script>
</body>
</html>
`
)
