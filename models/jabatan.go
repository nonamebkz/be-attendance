package models

type Jabatan struct {
	ID          int32   `gorm:"primaryKey;column:id_jabatan;type:serial4" json:"id_jabatan"`
	NamaJabatan string  `gorm:"column:nama_jabatan;size:100;not null" json:"nama_jabatan"`
	Deskripsi   *string `gorm:"column:deskripsi" json:"deskripsi"`
}
