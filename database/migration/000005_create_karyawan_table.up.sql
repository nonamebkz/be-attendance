-- Tabel karyawan (profil detail)
CREATE TABLE karyawan (
                          id_karyawan    SERIAL PRIMARY KEY,
                          user_id        INT UNIQUE REFERENCES app_user(id_user) ON DELETE CASCADE,
                          nip            VARCHAR(50) UNIQUE,
                          nama           VARCHAR(100) NOT NULL,
                          tanggal_lahir  DATE,
                          jenis_kelamin  VARCHAR(10),
                          jabatan_id     INT REFERENCES jabatan(id_jabatan),
                          alamat         TEXT,
                          no_telepon     VARCHAR(20),
                          foto           VARCHAR(255)
);
