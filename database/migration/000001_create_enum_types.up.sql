CREATE TYPE role_type AS ENUM ('admin', 'karyawan');
CREATE TYPE akun_status_type AS ENUM ('aktif', 'nonaktif');
CREATE TYPE absensi_status_type AS ENUM ('hadir', 'izin', 'sakit', 'alpha', 'cuti');
CREATE TYPE izin_jenis_type AS ENUM ('izin', 'sakit', 'cuti');
CREATE TYPE izin_status_type AS ENUM ('pending', 'disetujui', 'ditolak');
