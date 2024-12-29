CREATE TABLE "avatars" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "image_url" TEXT,
    "name" TEXT
);
