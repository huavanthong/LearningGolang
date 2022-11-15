CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;

CREATE TABLE items(
  id uuid DEFAULT gen_random_uuid() NOT NULL,
  name character varying NOT NULL
)