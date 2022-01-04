package swagger

import (
	"dcswitch/internal/config"
	"testing"
)

func BenchmarkGetSwaggerContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getSwaggerContent()
	}
}

func getSwaggerContent() {
	_, err := config.SwaggerFS.ReadFile("swagger.json")
	if err != nil {
		panic(err)
	}
}
