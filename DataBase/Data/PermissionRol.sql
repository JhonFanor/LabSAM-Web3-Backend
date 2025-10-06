INSERT INTO permission_rol (permission_id, role_id)
SELECT v.permission_id, v.role_id
FROM (VALUES
    (1, 1),
    (4, 2),
    (5, 2),
    (6, 2),
    (2, 3),
    (5, 3),
    (7, 3),
    (2, 4),
    (3, 4),
    (4, 4),
    (5, 4),
    (6, 4),
    (7, 4),
    (8, 4),
    (9, 4)
) AS v(permission_id, role_id)
ON CONFLICT (permission_id, role_id) DO NOTHING;

