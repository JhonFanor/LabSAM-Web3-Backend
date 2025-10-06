INSERT INTO roles (name, created_at, updated_at)
SELECT v.name, NOW(), NOW()
FROM (VALUES
    ('regular'),
    ('university'),
    ('business'),
    ('admin')
) AS v(name)
ON CONFLICT (name) DO NOTHING;

