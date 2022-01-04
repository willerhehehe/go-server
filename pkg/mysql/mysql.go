package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type Mysql struct {
	db          *sql.DB
	initialized bool
}

var DB = Mysql{db: nil, initialized: false}

type DBConnOptions struct {
	User        string
	Pwd         string
	Host        string
	Port        int
	Database    string
	MaxOpenConn int
	MaxIdleConn int
}

func (d *DBConnOptions) EnsureDefaults() {
	if d.MaxOpenConn == 0 {
		d.MaxOpenConn = 30
	}
	if d.MaxIdleConn == 0 {
		d.MaxIdleConn = 15
	}
}

func (m *Mysql) InitConn(o DBConnOptions) {
	o.EnsureDefaults()
	m.InitDBConn(o)
}

func (m *Mysql) InitDBConn(o DBConnOptions) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true", o.User, o.Pwd, o.Host, o.Port, o.Database)

	m.CloseConn()
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Error("Mysql init conn error")
		panic(err)
	}
	db.SetMaxOpenConns(o.MaxOpenConn)
	db.SetMaxIdleConns(o.MaxIdleConn)

	err = db.Ping()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Error("Mysql Ping error")
		panic(err)
	}
	m.db = db
	m.initialized = true
	log.Infoln("Mysql初始化")
}

func (m *Mysql) CloseConn() {
	if m.initialized {
		m.db.Close()
		m.db = nil
		m.initialized = false
	}
}

func (m *Mysql) checkInit() bool {
	if m.initialized == false {
		log.WithFields(log.Fields{"err": "DB还没有初始化"}).Error("Mysql GET conn error")
	}
	return m.initialized
}

func (m *Mysql) GetConn() *sql.DB {
	if init := m.checkInit(); init == false {
		return nil
	}
	return m.db
}

func (m *Mysql) GetTx() *sql.Tx {
	if init := m.checkInit(); init == false {
		return nil
	}
	tx, err := m.db.Begin()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Error("Mysql GET conn error")
		return nil
	}
	return tx
}

func (m *Mysql) RowExists(query string, args ...interface{}) (bool, error) {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := m.db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Errorf("error checking if row exists '%s' %v", args, err)
		return false, err
	}
	return exists, nil
}
