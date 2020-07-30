CREATE TABLE IF NOT EXISTS permissions(
   id  serial PRIMARY  KEY,
   name VARCHAR (50) UNIQUE NOT NULL,
   route VARCHAR (50) NOT NULL,
   created_at timestamp  NOT NULL  DEFAULT current_timestamp,
   updated_at timestamp  DEFAULT NULL
);
CREATE INDEX permission_index ON permissions(
   id
);



CREATE TRIGGER permission_updated_at_modtime BEFORE UPDATE ON permissions FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();