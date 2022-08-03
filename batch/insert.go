package batch

import "database/sql"

type IBatchInsert interface {
	Insert() error
}

type BatchInsert struct {
	db   *sql.DB
	data interface{}
}

func NewBatchInsert(db *sql.DB, data interface{}) IBatchInsert {
	return &BatchInsert{db: db, data: data}
}

func (batch *BatchInsert) Insert() error {
	_, err := batch.db.Exec("INSERT INTO mov_contable_detalle (monto, id_control) VALUES (?, ?)", batch.data)
	if err != nil {
		return err
	}
	return nil
}
