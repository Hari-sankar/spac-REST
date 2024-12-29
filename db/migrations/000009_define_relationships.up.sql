ALTER TABLE spaces ADD CONSTRAINT "fk_creator_spaces" FOREIGN KEY ("creator_id") REFERENCES users ("id");
ALTER TABLE users ADD CONSTRAINT "fk_users_avatar" FOREIGN KEY ("avatar_id") REFERENCES avatars ("id");
