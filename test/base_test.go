package test

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	setup(m)
	code := m.Run()
	shutdown(m)
	os.Exit(code)
}

func setup(m *testing.M) {
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
	}
	switch runtime.GOOS {
	case "windows":
		_ = os.Setenv("ENV", "DEV")
	case "darwin":
		_ = os.Setenv("ENV", "DEV")
	default:
	}
	log.Println("setup")
}

func shutdown(m *testing.M) {
	log.Println("shutdown")
}
