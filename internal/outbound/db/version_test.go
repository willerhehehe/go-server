package db

import (
	"dcswitch/internal/config"
	"dcswitch/internal/domain"
	"dcswitch/pkg/mysql"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestNotFoundError(t *testing.T) {
	swRepo := SwitchVersionDBRepo{}
	v, err := swRepo.Get(12312321)
	_, ok := err.(mysql.NotFoundError)
	if !ok {
		t.Errorf("NotFoundError Wrong %v\n", v)
	}
}

func TestAddSwitchVersion(t *testing.T) {
	swRepo := SwitchVersionDBRepo{}
	version := domain.SwitchVersion{Name: "test001", Time: time.Now()}
	id, err := swRepo.Add(version)
	if err != nil {
		t.Error(err)
	}
	t.Logf("add id : %v\n", id)
}

func TestEditSwitchVersion(t *testing.T) {
	swRepo := SwitchVersionDBRepo{}
	v101, err := swRepo.Get(101)
	_, ok := err.(mysql.NotFoundError)
	if ok {
		addV := domain.SwitchVersion{Id: 101, Name: "test", Time: time.Now()}
		_, err := swRepo.Add(addV)
		if err != nil {
			t.Error(err)
		}
		v101, err = swRepo.Get(101)
		if err != nil {
			t.Error(err)
		}
	}
	randName := RandStr(10)
	v101.Name = randName
	_, err = swRepo.EditName(101, randName)
	if err != nil {
		t.Error(err)
	}
	v, err := swRepo.Get(101)
	if err != nil {
		t.Error(err)
	}

	if v.Name != randName {
		t.Errorf("EditError %v != %v", v.Name, randName)
	}
}

// TODO: 增加测试数据库及测试前清表功能
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
	config.GlobalConf.InitConfig()
	config.InitDB()
	logrus.Println("setup")

}

func shutdown(m *testing.M) {
	logrus.Println("shutdown")
}

func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
