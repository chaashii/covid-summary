package repository

import (
	"covid/db"
)

type Reposirory struct{}

func (r Reposirory) CreateCovid(dto interface{}) error {

	if err := db.Conn.Create(dto).Error; err != nil {
		return err
	}

	return nil
}

func (r Reposirory) DeleteCovid(dto interface{}) error {

	if err := db.Conn.Exec("DELETE FROM COVID_SUMMARIES").Error; err != nil {
		return err
	}

	return nil
}
