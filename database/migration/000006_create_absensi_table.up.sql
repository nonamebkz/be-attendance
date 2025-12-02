-- Tabel absensi
CREATE TABLE absensi (
                         id_absensi    SERIAL PRIMARY KEY,
                         karyawan_id   INT NOT NULL REFERENCES karyawan(id_karyawan) ON DELETE CASCADE,
                         tanggal       DATE NOT NULL,
                         waktu_masuk   TIMESTAMPTZ,
                         waktu_pulang  TIMESTAMPTZ,
                         status        absensi_status_type NOT NULL DEFAULT 'hadir',
                         keterangan    TEXT,
                         lokasi        VARCHAR(255),
                         created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                         UNIQUE (karyawan_id, tanggal)
);
