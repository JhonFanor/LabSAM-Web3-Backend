INSERT INTO university_types (id, name)
SELECT v.id, v.name
FROM (VALUES
    (1, 'Privada'),
    (2, 'Publica')
) AS v(id, name)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name;

