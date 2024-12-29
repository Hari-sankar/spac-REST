CREATE TABLE space_elements (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "elementId" UUID NOT NULL,
    "spaceId" UUID NOT NULL,
    "x" INTEGER NOT NULL,
    "y" INTEGER NOT NULL,
    CONSTRAINT "fk_elementId" FOREIGN KEY ("elementId") REFERENCES elements ("id"),
    CONSTRAINT "fk_spaceId" FOREIGN KEY ("spaceId") REFERENCES spaces ("id")
);
