-- Tabel password_reset (lupa password)
CREATE TABLE password_reset (
                                id_reset    SERIAL PRIMARY KEY,
                                user_id     INT NOT NULL REFERENCES app_user(id_user) ON DELETE CASCADE,
                                token       VARCHAR(255) UNIQUE NOT NULL,
                                expired_at  TIMESTAMPTZ NOT NULL,
                                used_at     TIMESTAMPTZ,
                                created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
