ALTER TABLE users ADD COLUMN role_id INT;
ALTER TABLE users ADD CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(role_id);
