-- Create "permissions" table
CREATE TABLE "permissions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "key" text NOT NULL,
  "description" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_permissions_key" to table: "permissions"
CREATE UNIQUE INDEX "idx_permissions_key" ON "permissions" ("key");
-- Create "roles" table
CREATE TABLE "roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "description" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_roles_name" to table: "roles"
CREATE UNIQUE INDEX "idx_roles_name" ON "roles" ("name");
-- Create "role_permissions" table
CREATE TABLE "role_permissions" (
  "role_id" uuid NOT NULL,
  "permission_id" uuid NOT NULL,
  PRIMARY KEY ("role_id", "permission_id"),
  CONSTRAINT "fk_role_permissions_permission" FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_role_permissions_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "email" text NOT NULL,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "phone" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create "user_credentials" table
CREATE TABLE "user_credentials" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "type" character varying(20) NOT NULL,
  "identifier" text NOT NULL,
  "secret_hash" text NOT NULL,
  "verified_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_credentials" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_credentials_user_id" to table: "user_credentials"
CREATE INDEX "idx_user_credentials_user_id" ON "user_credentials" ("user_id");
-- Create "user_roles" table
CREATE TABLE "user_roles" (
  "user_id" uuid NOT NULL,
  "role_id" uuid NOT NULL,
  "scope_type" character varying(50) NULL,
  "scope_id" uuid NULL,
  PRIMARY KEY ("user_id", "role_id"),
  CONSTRAINT "fk_user_roles_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_user_roles_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "user_sessions" table
CREATE TABLE "user_sessions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "refresh_token_hash" text NOT NULL,
  "user_agent" text NOT NULL,
  "ip_address" text NOT NULL,
  "expires_at" timestamptz NULL,
  "revoked_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_sessions" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_sessions_refresh_token_hash" to table: "user_sessions"
CREATE UNIQUE INDEX "idx_user_sessions_refresh_token_hash" ON "user_sessions" ("refresh_token_hash");
-- Create index "idx_user_sessions_user_id" to table: "user_sessions"
CREATE INDEX "idx_user_sessions_user_id" ON "user_sessions" ("user_id");
