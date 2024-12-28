CREATE TABLE map_elements (
    id VARCHAR(255) PRIMARY KEY,
    map_id VARCHAR(255) NOT NULL,
    element_id VARCHAR(255),
    x INTEGER,
    y INTEGER,
    FOREIGN KEY (map_id) REFERENCES maps(id),
    FOREIGN KEY (element_id) REFERENCES elements(id)
);
