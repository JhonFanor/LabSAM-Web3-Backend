INSERT INTO permissions (id, name, created_at, updated_at)
SELECT v.id, v.name, NOW(), NOW()
FROM (VALUES
    (1, 'bank-of-resume:create'),
    (2, 'company:create'),
    (3, 'documentation:create'),
    (4, 'educational-offer:create'),
    (5, 'event:create'),
    (6, 'investigation:create'),
    (7, 'job-board:create'),
    (8, 'legislation:create'),
    (9, 'news:create')
) AS v(id, name)
ON CONFLICT (id) DO UPDATE 
SET name = EXCLUDED.name,
    updated_at = NOW();

