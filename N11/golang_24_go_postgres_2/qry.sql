CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR,
    lastname VARCHAR,
    age INT
);

INSERT INTO "users" (firstname, lastname, age) 
-- VALUES ('John', 'Doe', 20);
VALUES ('Hamidulloh', 'Hamidullayev', 22);


CREATE TABLE "car" (
    id serial PRIMARY KEY,
    user_id int,
    model varchar,
    price numeric,
    year int
);

INSERT INTO "car" (user_id, model, price, year)
    VALUES
(2, 'Tahoe', 100000, 2024),
(2, 'Mercedes', 134543, 2024),
(3, 'Toyota', 13000, 2024),
(2, 'Tesla', 239834, 2020),
(2, 'BMW', 134000, 2024),
(3, 'Gentra', 20000, 2022),
(4, 'Gentra', 20000, 2022),
(10, 'K5', 20000, 2022),
(5, 'Malibu', 2333, 2022),
(10, 'Tracer', 2000, 2022),
(90, 'Spark', 12000, 2022);



SELECT
    model as car_model,
    price as car_price,
    year as mashina_yili
FROM   
    car;


SELECT
    u.id,
    u.firstname,
    u.lastname,
    u.age,
    c.id,
    c.user_id,
    c.model,
    c.price,
    c.year
FROM
    users AS u
INNER JOIN car AS c ON u.id = c.user_id
WHERE u.id = 2
;

SELECT 
    u.first_name,
    Array_Agg(c.model)
FROM 
    car c
GROUP BY 
