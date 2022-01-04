package static

import "embed"

//go:embed swagger.json

//go:embed swagger-ui-bundle.js
//go:embed swagger-ui.css

//go:embed redoc.standalone.js
//go:embed redoc.standalone.css

//go:embed dcswitch.png
var SwaggerFS embed.FS
