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

CREATE TABLE user_cars
(
    user_id uuid NOT NULL,
    car_id uuid NOT NULL,
    PRIMARY KEY (user_id, car_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (car_id) REFERENCES car(id) ON DELETE CASCADE
);

INSERT INTO users (name,car_id) VALUES ('Alisher','024c14e4-9ea6-4093-8017-383df35a77f6');
INSERT INTO users (name,car_id) VALUES ('Bobur','fbc4bf77-a832-41da-bd38-0f18fbaac628');

INSERT INTO car (name, year, brand) VALUES ('Gentra', 2020, 'Chevrolet');
INSERT INTO car (name, year, brand) VALUES ('Nexia 3', 2019, 'Ravon');



INSERT INTO user_cars (user_id, car_id) VALUES ('b2afa29e-96bc-41e8-8495-79636e6047af', 'fbc4bf77-a832-41da-bd38-0f18fbaac628');
INSERT INTO user_cars (user_id, car_id) VALUES ('b5a6566c-61af-455b-836d-2049a899c99c', '024c14e4-9ea6-4093-8017-383df35a77f6');
INSERT INTO user_cars (user_id, car_id) VALUES ('b2afa29e-96bc-41e8-8495-79636e6047af', '024c14e4-9ea6-4093-8017-383df35a77f6');
