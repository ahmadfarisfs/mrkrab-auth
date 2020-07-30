CREATE TABLE IF NOT EXISTS roles(
   id  serial PRIMARY  KEY,
   name VARCHAR (255) UNIQUE NOT NULL,
   created_at timestamp  NOT NULL  DEFAULT current_timestamp,
   updated_at timestamp  DEFAULT NULL
);
CREATE INDEX role_index ON roles(
   id
);
CREATE TRIGGER role_updated_at_modtime BEFORE UPDATE ON roles FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();