
-- Tabel user
CREATE TABLE IF NOT EXISTS app_user (
                          id_user        SERIAL PRIMARY KEY,
                          username       VARCHAR(50) UNIQUE NOT NULL,
                          email          VARCHAR(100) UNIQUE NOT NULL,
                          password_hash  VARCHAR(255) NOT NULL,
                          nama_lengkap   VARCHAR(100),
                          role           role_type NOT NULL DEFAULT 'karyawan',
                          status_akun    akun_status_type NOT NULL DEFAULT 'aktif',
                          created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                          updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
