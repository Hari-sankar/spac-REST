CREATE TABLE users (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "username" TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL,
    "avatarId" UUID,
    "role" "Role" NOT NULL,
    CONSTRAINT "fk_avatarId" FOREIGN KEY ("avatarId") REFERENCES avatars ("id")
);
