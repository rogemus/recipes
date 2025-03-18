CREATE TABLE IF NOT EXISTS
  units (
    id bigserial PRIMARY KEY,  
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    version integer NOT NULL DEFAULT 1
  );


CREATE INDEX IF NOT EXISTS units_name_idx ON units USING GIN (to_tsvector('simple', name));
