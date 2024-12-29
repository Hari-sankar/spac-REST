CREATE TABLE map_elements (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "map_id" UUID NOT NULL,
    "element_id" UUID NOT NULL,
    "x" INTEGER,
    "y" INTEGER,
    CONSTRAINT "fk_mapId" FOREIGN KEY ("map_id") REFERENCES maps ("id"),
    CONSTRAINT "fk_elementId" FOREIGN KEY ("element_id") REFERENCES elements ("id")
);
