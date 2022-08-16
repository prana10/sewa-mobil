package mobil

type AddMobilInput struct {
	Name            string `json:"nama_mobil" binding:"required"`
	Type            string `json:"jenis_mobil" binding:"required"`
	Plat            string `json:"plat_mobil" binding:"required"`
	RentalPrice     int    `json:"harga_sewa" binding:"required"`
	PublicationYear int    `json:"tahun_produksi"`
}
