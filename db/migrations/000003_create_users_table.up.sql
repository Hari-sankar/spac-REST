CREATE TABLE "User" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "username" TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL,
    "avatarId" UUID REFERENCES "Avatar"("id") ON DELETE SET NULL ON UPDATE CASCADE,
    "role" "Role" NOT NULL
);
