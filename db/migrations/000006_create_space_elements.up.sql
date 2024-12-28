CREATE TABLE space_elements (
    id VARCHAR(255) PRIMARY KEY,
    element_id VARCHAR(255) NOT NULL,
    space_id VARCHAR(255) NOT NULL,
    x INTEGER NOT NULL,
    y INTEGER NOT NULL,
    FOREIGN KEY (element_id) REFERENCES elements(id),
    FOREIGN KEY (space_id) REFERENCES spaces(id)
);
