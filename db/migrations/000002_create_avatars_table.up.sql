CREATE TABLE "avatars" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "imageUrl" TEXT,
    "name" TEXT
);
