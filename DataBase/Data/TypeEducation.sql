INSERT INTO type_education (id, name)
SELECT v.id, v.name
FROM (VALUES
    (1, 'Curso'),
    (2, 'Taller'),
    (3, 'Seminario'),
    (4, 'Diplomado'),
    (5, 'Conferencia'),
    (6, 'Webinar'),
    (7, 'Charla'),
    (8, 'Panel'),
    (9, 'Simposio'),
    (10, 'Foro'),
    (11, 'Especialización'),
    (12, 'Maestría'),
    (13, 'Doctorado'),
    (14, 'Formación continua'),
    (15, 'Programa técnico'),
    (16, 'Programa tecnológico'),
    (17, 'Pregrado'),
    (18, 'Posgrado'),
    (19, 'Capacitación'),
    (20, 'Entrenamiento')
) AS v(id, name)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name;
