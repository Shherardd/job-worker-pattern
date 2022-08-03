package jobs

import (
	"errors"
	"worker-pools/batch"
	"worker-pools/repos"
)

type TransferirMovContableDetalles interface {
	Do() error
	initMovContablesDetalle() error
}

type transferirMovContableDetalles struct {
	repo         *repos.MovContableDetalleRepository
	idControl    int
	keypartition int64
}

func NewTransferirMovContableDetalles(repo *repos.MovContableDetalleRepository,
	id int, key int64) TransferirMovContableDetalles {
	return &transferirMovContableDetalles{
		repo:         repo,
		idControl:    id,
		keypartition: key,
	}
}

func (t *transferirMovContableDetalles) Do() error {
	movContables, err := t.repo.GetAll()

	if err != nil {
		return errors.New("error al obtener los movimientos contables")
	}
	batch := batch.NewBatchInsert(t.repo.DB, movContables)
	err = batch.Insert()
	if err != nil {
		return errors.New("error al insertar los movimientos contables")
	}
	return nil
}

func (t *transferirMovContableDetalles) initMovContablesDetalle() error {

	return nil
}
