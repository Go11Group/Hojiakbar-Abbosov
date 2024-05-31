CREATE TABLE users
(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    car_id uuid NOT NULL REFERENCES car(id)
);

CREATE TABLE car 
(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    year int,
    brand VARCHAR NOT NULL
);


INSERT INTO car(name,year,brand)
VALUES('Cobalt',2020,'Ravon');

INSERT INTO users(name,car_id)
VALUES('Abdulloh','4047c266-e340-4e1d-bc32-86c65086171b');