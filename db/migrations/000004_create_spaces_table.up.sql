CREATE TABLE spaces (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" TEXT NOT NULL,
    "width" INTEGER NOT NULL,
    "height" INTEGER NOT NULL,
    "thumbnail" TEXT,
    "creatorId" UUID NOT NULL,
    CONSTRAINT "fk_creatorId" FOREIGN KEY ("creatorId") REFERENCES users ("id")
);
