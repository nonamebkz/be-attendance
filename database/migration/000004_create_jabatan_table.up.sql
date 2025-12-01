-- Tabel jabatan
CREATE TABLE jabatan (
                         id_jabatan    SERIAL PRIMARY KEY,
                         nama_jabatan  VARCHAR(100) UNIQUE NOT NULL,
                         deskripsi     TEXT
);
