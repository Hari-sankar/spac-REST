CREATE TABLE space_elements (
    id SERIAL PRIMARY KEY,
    element_id INTEGER NOT NULL,
    space_id INTEGER NOT NULL,
    x INTEGER NOT NULL,
    y INTEGER NOT NULL,
    FOREIGN KEY (element_id) REFERENCES elements(id),
    FOREIGN KEY (space_id) REFERENCES spaces(id)
);
