-- Create composite type "address"
CREATE TYPE "address" AS ("street" text, "city" text);
-- Create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "address" "address" NOT NULL, PRIMARY KEY ("id"));