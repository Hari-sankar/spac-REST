CREATE TABLE maps (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "width" INTEGER NOT NULL,
    "height" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "thumbnail" TEXT NOT NULL
);
