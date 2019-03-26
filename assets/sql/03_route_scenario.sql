INSERT INTO nodes(node_id, name) VALUES
    (NULL, 'Node_A'),
    (NULL, 'Node_B'),
    (NULL, 'Node_C'),
    (NULL, 'Node_D');

INSERT INTO edges(edge_id, from_id, to_id, duration, cost) VALUES
    (NULL, 1, 3, 1, 20),
    (NULL, 1, 4, 1, 14),
    (NULL, 3, 2, 1, 12),
    (NULL, 4, 2, 1, 8);
