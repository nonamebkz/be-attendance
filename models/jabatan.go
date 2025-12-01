package models

type Jabatan struct {
	ID          int32   `db:"id_jabatan" json:"id_jabatan"`
	NamaJabatan string  `db:"nama_jabatan" json:"nama_jabatan"`
	Deskripsi   *string `db:"deskripsi" json:"deskripsi"`
}
