CREATE TABLE IF NOT EXISTS
  "units" (
    "id" INTEGER NOT NULL UNIQUE,
    "name" VARCHAR NOT NULL,
    "created" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
  );

INSERT INTO
  units (name)
VALUES
  ('ml (mililitry)'),
  ('l (litry)'),
  ('łyżeczka (łyżeczki)'),
  ('łyżka (łyżki)'),
  ('szklanka (szklanki)'),
  ('g (gramy)'),
  ('kg (kilogramy)'),
  ('szt. (sztuka/sztuki)'),
  ('plaster (plastry)'),
  ('ząbek (ząbki) [np. czosnku]'),
  ('listek (listki) [np. laurowy, bazylii]'),
  ('garść (garście)'),
  ('szczypta (szczypty)'),
  ('kropla (krople)');
