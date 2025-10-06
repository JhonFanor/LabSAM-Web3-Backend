INSERT INTO subtopics (id, name, topic_id, created_at, updated_at)
SELECT v.id, v.name, v.topic_id, NOW(), NOW()
FROM (VALUES
    -- Subtopics del topic 1 (Blockchain)
    (1, 'Blockchain wallets', 1),
    (2, 'Stablecoins', 1),
    (3, 'Smartcontracts', 1),
    (4, 'Dapps', 1),
    (5, 'Blockchain platforms', 1),
    (6, 'NFT', 1),
    (7, 'Consensus mechanisms', 1),
    (8, 'Decentralyzed identity', 1),
    (9, 'Tokenization', 1),
    (10, 'Blockchain interoperability', 1),
    (11, 'Oracles', 1),
    (12, 'DAOs', 1),
    (13, 'Secure Multiparty Computation', 1),
    (14, 'Zero knowledge proofs', 1),
    (15, 'Real World Assets (RWA)', 1),
    (16, 'Layer 2', 1),

    -- Subtopics del topic 2 (Metaverse)
    (17, 'Realidad virtual', 2),
    (18, 'Realidad aumentada', 2),

    -- Subtopics del topic 3 (Inteligencia artificial)
    (19, 'Machine Learning', 3),
    (20, 'Analitica de datos', 3),
    (21, 'Redes neuronales artificiales', 3),
    (22, 'Procesamiento de lenguaje natural', 3),

    -- Subtopics del topic 4 (Computación en la nube)
    (23, 'Docker', 4),
    (24, 'Kubernetes', 4),
    (25, 'Infraestructura como servicio (LaaS)', 4),
    (26, 'Plataformas como servicio (PaaS)', 4),
    (27, 'Software como servicio (SaaS)', 4)
) AS v(id, name, topic_id)
ON CONFLICT (id) DO UPDATE 
SET name = EXCLUDED.name,
    topic_id = EXCLUDED.topic_id,
    updated_at = NOW();

