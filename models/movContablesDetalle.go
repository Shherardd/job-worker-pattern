package models

type MovContableDetalle struct {
	Monto     int
	IdControl int
}

type MovContableDetalles []*MovContableDetalle
