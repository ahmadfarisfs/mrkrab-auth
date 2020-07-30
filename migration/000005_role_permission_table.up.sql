CREATE TABLE IF NOT EXISTS role_permissions(
   id  serial PRIMARY  KEY,
   role_id BIGINT NOT NULL,
   permission_id BIGINT NOT NULL,
   created_at timestamp  NOT NULL  DEFAULT current_timestamp,
   updated_at timestamp  DEFAULT NULL
);
CREATE INDEX role_permission_index ON role_permissions(
   role_id,
   permission_id
);
CREATE TRIGGER role_permission_updated_at_modtime BEFORE UPDATE ON role_permissions FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();