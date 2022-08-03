package repos

import (
	"database/sql"
	"worker-pools/models"
)

type IMovContableDetalleRepository interface {
	GetAll() (*models.MovContableDetalles, error)
}

type MovContableDetalleRepository struct {
	DB *sql.DB
}

func NewMovContableDetalleRepository(db *sql.DB) IMovContableDetalleRepository {
	return &MovContableDetalleRepository{DB: db}
}

func (mov *MovContableDetalleRepository) GetAll() (*models.MovContableDetalles, error) {
	rows, err := mov.DB.Query("SELECT * FROM mov_contable_detalle")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movContableDetalles models.MovContableDetalles
	for rows.Next() {
		var movContableDetalle models.MovContableDetalle
		err := rows.Scan(&movContableDetalle.Monto, &movContableDetalle.IdControl)
		if err != nil {
			return nil, err
		}
		movContableDetalles = append(movContableDetalles, &movContableDetalle)
	}
	return &movContableDetalles, nil
}
