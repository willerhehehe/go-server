package db

import (
	"dcswitch/internal/domain"
	"dcswitch/pkg/mysql"
	log "github.com/sirupsen/logrus"
)

type SwitchVersionDBRepo struct{}

func (repo SwitchVersionDBRepo) Get(id int64) (domain.SwitchVersion, error) {
	var version domain.SwitchVersion
	sql := "SELECT id, version_time, name FROM switch_version WHERE id = ?"
	exists, _ := repo.CheckExist(id)
	if !exists {
		return version, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"id":  id,
		}).Error("Query Error")
		return version, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&version.Id, &version.Time, &version.Name)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return version, err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return version, err
	}
	return version, nil
}

func (repo SwitchVersionDBRepo) CheckExist(id int64) (bool, error) {
	sql := "SELECT id, version_time, name FROM switch_version WHERE id = ?"
	exists, err := mysql.DB.RowExists(sql, id)
	if !exists {
		return false, mysql.NotFoundError{}
	}
	return exists, err
}

func (repo SwitchVersionDBRepo) GetAll() ([]domain.SwitchVersion, error) {
	sql := "SELECT id, version_time, name FROM switch_version"
	var versions []domain.SwitchVersion
	db := mysql.DB.GetConn()

	rows, err := db.Query(sql)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
		}).Error("Query Error")
		return versions, err
	}
	defer rows.Close()
	for rows.Next() {
		var v domain.SwitchVersion
		err := rows.Scan(&v.Id, &v.Time, &v.Name)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return versions, err
		}
		versions = append(versions, v)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return versions, err
	}
	return versions, nil
}

func (repo SwitchVersionDBRepo) Add(v domain.SwitchVersion) (int64, error) {
	sql := "INSERT INTO switch_version (version_time, name) VALUES (?, ?)"
	db := mysql.DB.GetConn()
	res, err := db.Exec(sql, v.Time.Format("2006-01-02 15:04:05"), v.Name)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":          sql,
			"version_time": v.Time,
			"name":         v.Name,
		}).Error(err.Error())
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (repo SwitchVersionDBRepo) EditName(id int64, name string) (int64, error) {
	sql := "UPDATE switch_version SET name = ? WHERE ID = ?"
	db := mysql.DB.GetConn()
	exists, err := repo.CheckExist(id)
	if err != nil || !exists {
		return -1, err
	}
	res, err := db.Exec(sql, name, id)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":  sql,
			"id":   id,
			"name": name,
		}).Error(err.Error())
		return -1, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rowsAffected, nil
}
