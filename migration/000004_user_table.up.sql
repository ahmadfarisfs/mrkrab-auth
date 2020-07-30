CREATE TABLE IF NOT EXISTS users(
   id  serial PRIMARY  KEY,
   email VARCHAR (255) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   role_id BIGINT DEFAULT NULL,
   created_at timestamp  NOT NULL  DEFAULT current_timestamp,
   updated_at timestamp  DEFAULT NULL
);
CREATE INDEX user_index ON users(
   id,
   email,
   role_id
);
CREATE TRIGGER user_updated_at_modtime BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();