INSERT INTO type_of_law (id, name)
SELECT v.id, v.name
FROM (VALUES
    (1, 'Ley'),
    (2, 'Decreto'),
    (3, 'Resolución'),
    (4, 'Acuerdo'),
    (5, 'Ordenanza'),
    (6, 'Circular'),
    (7, 'Directiva'),
    (8, 'Reglamento'),
    (9, 'Estatuto'),
    (10, 'Constitución'),
    (11, 'Código'),
    (12, 'Decreto-Ley'),
    (13, 'Decreto Supremo'),
    (14, 'Auto'),
    (15, 'Sentencia'),
    (16, 'Protocolo'),
    (17, 'Convenio Internacional'),
    (18, 'Tratado'),
    (19, 'Resolución Ministerial'),
    (20, 'Resolución Presidencial')
) AS v(id, name)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name;