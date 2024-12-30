CREATE TABLE map_elements (
    id SERIAL PRIMARY KEY,
    map_id INTEGER NOT NULL,
    element_id INTEGER NOT NULL,
    x INTEGER,
    y INTEGER,
    FOREIGN KEY (map_id) REFERENCES maps(id),
    FOREIGN KEY (element_id) REFERENCES elements(id)
);
