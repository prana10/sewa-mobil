package mobil

import "time"

type Mobil struct {
	ID              int       `json:"id"`
	Name            string    `json:"nama_mobil"`
	Type            string    `json:"jenis_mobil"`
	Plat            string    `json:"plat_mobil"`
	RentalPrice     int       `json:"harga_sewa"`
	PublicationYear int       `json:"tahun_produksi"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
