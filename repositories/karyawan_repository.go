package repositories

type KaryawanRepository interface {
}

type karyawanRepository struct {
}

func NewKaryawanRepository() KaryawanRepository {
	return &karyawanRepository{}
}
