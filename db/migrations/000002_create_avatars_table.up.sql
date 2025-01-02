CREATE TABLE "Avatar" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "image_url" TEXT,
    "name" TEXT
);
