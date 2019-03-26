CREATE TABLE IF NOT EXISTS edges (
   edge_id       INT NOT NULL AUTO_INCREMENT,
   from_id       INT NOT NULL,
   to_id         INT NOT NULL,
   duration      INT NOT NULL,
   cost          INT NOT NULL,
   PRIMARY KEY(edge_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS nodes (
    node_id         INT NOT NULL AUTO_INCREMENT,
    name CHAR(20)   NOT NULL,
    PRIMARY KEY(node_id)
) ENGINE=InnoDB;
