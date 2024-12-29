CREATE TABLE space_elements (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "element_id" UUID NOT NULL,
    "space_id" UUID NOT NULL,
    "x" INTEGER NOT NULL,
    "y" INTEGER NOT NULL,
    CONSTRAINT "fk_elementId" FOREIGN KEY ("element_id") REFERENCES elements ("id"),
    CONSTRAINT "fk_spaceId" FOREIGN KEY ("space_id") REFERENCES spaces ("id")
);
