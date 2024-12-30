CREATE TABLE elements (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "width" INTEGER NOT NULL,
    "height" INTEGER NOT NULL,
    "static" BOOLEAN NOT NULL,
    "imageUrl" TEXT NOT NULL
);
