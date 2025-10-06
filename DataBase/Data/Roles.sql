INSERT INTO roles (id, name, created_at, updated_at)
SELECT v.id, v.name, NOW(), NOW()
FROM (VALUES
    (1, 'regular'),
    (2, 'university'),
    (3, 'business'),
    (4, 'admin')
) AS v(id, name)
ON CONFLICT (name) DO NOTHING;
