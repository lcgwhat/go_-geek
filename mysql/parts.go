package mysql

import (
	"database/sql"
)

type ActiveRecord interface {
	getById(id int)(Part,error)
	getByName()
	create()
}
type Part struct {
	Id int
	Name string
}
type parts struct {
	db *sql.DB
}

func NewPart(db *sql.DB)*parts{
	return &parts{
		db: db,
	}
}

func (p parts)GetById(id int) (Part, error)  {
	var part Part
	sqlStr := "select id, name from parts where id = ?"
	err := p.db.QueryRow(sqlStr,id).Scan(&part.Id, &part.Name)
	if err!=nil {
		return Part{}, err
	}
	return part, nil
}

func (p parts)GetByName()  {
}

func (p parts)Create()  {
}