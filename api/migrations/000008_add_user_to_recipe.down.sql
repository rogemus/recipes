ALTER TABLE recipes DROP COLUMN user_id;
ALTER TABLE recipes DROP CONSTRAINT recipes_user_id_fkey;
