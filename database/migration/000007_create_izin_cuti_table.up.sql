-- Tabel izin / cuti
CREATE TABLE izin_cuti (
                           id_izin            SERIAL PRIMARY KEY,
                           karyawan_id        INT NOT NULL REFERENCES karyawan(id_karyawan) ON DELETE CASCADE,
                           tanggal_mulai      DATE NOT NULL,
                           tanggal_selesai    DATE NOT NULL,
                           jenis              izin_jenis_type NOT NULL,
                           alasan             TEXT,
                           file_bukti         VARCHAR(255),
                           status_persetujuan izin_status_type NOT NULL DEFAULT 'pending',
                           created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                           approved_at        TIMESTAMPTZ
);
