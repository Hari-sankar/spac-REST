ALTER TABLE spaces ADD CONSTRAINT "fk_creator_spaces" FOREIGN KEY ("creatorId") REFERENCES users ("id");
ALTER TABLE users ADD CONSTRAINT "fk_users_avatar" FOREIGN KEY ("avatarId") REFERENCES avatars ("id");
