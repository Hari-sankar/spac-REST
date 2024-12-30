CREATE TABLE spaces (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" TEXT NOT NULL,
    "width" INTEGER NOT NULL,
    "height" INTEGER NOT NULL,
    "thumbnail" TEXT,
    "creator_id" UUID NOT NULL,
    CONSTRAINT "fk_creatorId" FOREIGN KEY ("creator_id") REFERENCES users ("id")
);
