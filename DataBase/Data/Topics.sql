INSERT INTO topics (id, name, created_at, updated_at)
SELECT v.id, v.name, NOW(), NOW()
FROM (VALUES
    (1, 'Blockchain'),
    (2, 'Metaverse'),
    (3, 'Inteligencia artificial'),
    (4, 'Computación en la nube')
) AS v(id, name)
ON CONFLICT (id) DO UPDATE 
SET name = EXCLUDED.name,
    updated_at = NOW();

