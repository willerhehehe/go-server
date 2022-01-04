package test

import (
	"fmt"
	eg "github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestErrorGroup(t *testing.T) {
	ExampleGroupJustErrors()
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExampleGroupJustErrors()
	}
}

func TestWrapError(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		//fmt.Printf("original error: %T %v+\n", err, err)
		fmt.Printf("original error: %T %v\n", eg.Cause(err), eg.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
	}
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		//return nil, fmt.Errorf("open failed %w\n", err)
		return nil, eg.Wrap(err, "open failed")
	}
	defer f.Close()
	return nil, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	//return config, fmt.Errorf("could not read config %w\n", err)
	return config, eg.WithMessage(err, "could not read config")
}

func ExampleGroupJustErrors() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org",
		"http://www.google.com",
		"http://www.baidu.com",
	}
	for _, url := range urls {
		g.Go(func() error {
			resp, err := http.Get(url)
			fmt.Println(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err)
	}

}

func StrConvert(split string, a ...interface{}) string {
	str := ""

	for index := 0; index < len(a); index++ {
		str1 := fmt.Sprintf("%v", a[index])
		if index > 0 {
			str += split + str1
		} else {
			str += str1
		}
	}
	return str
}

func TestJoin(t *testing.T) {
	fmt.Println(StrConvert(", ", "a", 1, 313, 3.1))

}
