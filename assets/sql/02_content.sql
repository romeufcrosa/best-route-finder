INSERT INTO nodes(node_id, name) VALUES
    (NULL, 'NODE_A'),
    (NULL, 'NODE_B'),
    (NULL, 'NODE_C'),
    (NULL, 'NODE_D'),
    (NULL, 'NODE_E'),
    (NULL, 'NODE_F'),
    (NULL, 'NODE_G'),
    (NULL, 'NODE_H'),
    (NULL, 'NODE_I');

INSERT INTO edges(edge_id, from_id, to_id, duration, cost) VALUES
    (NULL, 1, 3, 1, 20),
    (NULL, 1, 8, 10, 1),
    (NULL, 3, 2, 1, 12),
    (NULL, 1, 5, 30, 5),
    (NULL, 8, 5, 30, 1),
    (NULL, 5, 4, 3, 5),
    (NULL, 4, 6, 4, 50),
    (NULL, 6, 9, 45, 50),
    (NULL, 6, 7, 40, 50),
    (NULL, 9, 2, 65, 5),
    (NULL, 7, 2, 64, 73);
