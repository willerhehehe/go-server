package logger

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Option struct {
	LogPath string
}

func InitLogger(o Option) {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	var w io.Writer
	logrus.SetFormatter(customFormatter)
	if o.LogPath == "" {
		w = os.Stdout
	} else {
		w = getLogWriter(o.LogPath)
	}
	logrus.SetOutput(w)
	logrus.SetLevel(logrus.InfoLevel)
}

func getLogWriter(path string) io.Writer {
	_ = os.Mkdir(filepath.Dir(path), os.ModePerm)
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(60*60*24*15)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60*60*24)*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return writer
}
