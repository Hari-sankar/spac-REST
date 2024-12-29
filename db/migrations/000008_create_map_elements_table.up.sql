CREATE TABLE map_elements (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "mapId" UUID NOT NULL,
    "elementId" UUID NOT NULL,
    "x" INTEGER,
    "y" INTEGER,
    CONSTRAINT "fk_mapId" FOREIGN KEY ("mapId") REFERENCES maps ("id"),
    CONSTRAINT "fk_elementId" FOREIGN KEY ("elementId") REFERENCES elements ("id")
);
