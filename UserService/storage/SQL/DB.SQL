CREATE TABLE types(
    id INT NOT NULL PRIMARY KEY,
    type VARCHAR(50)
);

CREATE TABLE users(
    id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email TEXT [] NOT NULL,
    bio VARCHAR(50) NOT NULL,
    phone_number TEXT [] NOT NULL,
    type_id INT NOT NULL,
    status BOOLEAN NOT NULL,
    FOREIGN KEY (type_id)
    REFERENCES types(id)
    ON DELETE CASCADE
);
CREATE TABLE addresses(
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid,
    country VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    district VARCHAR(50) NOT NULL ,
    postal_code BIGINT NOT NULL,
    FOREIGN KEY(user_id)
      REFERENCES users(id)
      ON DELETE CASCADE
);
SELECT u.*, t.Type FROM users u join types t on  t.id = u.TypeID;
SELECT id FROM Users;
SELECT * FROM addresses;
SELECT id, city, district, country, postal_code FROM addresses WHERE user_id = '1de759e7-ab11-446f-9d1d-0333936aed4a';
SELECT id, first_name, last_name, email, bio, phone_number, type_id, status FROM users