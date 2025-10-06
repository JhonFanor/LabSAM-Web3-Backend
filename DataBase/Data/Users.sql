INSERT INTO users (email, avatar, password, role_id, email_verified, created_at, updated_at)
VALUES ('admin@admin.com', '', '$2a$10$ZZQMa/w76ZbEL8qXYQ6CKOD8Up8i4DH2cpaSj6KUPCfSRGfajhORq', 4, true, NOW(), NOW())
ON CONFLICT (email) DO UPDATE
SET avatar = EXCLUDED.avatar,
    password = EXCLUDED.password,
    role_id = EXCLUDED.role_id,
    email_verified = EXCLUDED.email_verified,
    updated_at = NOW();

