CREATE TABLE client (
    id INT PRIMARY KEY,
    name VARCHAR,
    age INT
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    client_id INT REFERENCES client (id)
);

INSERT INTO client (id, name, age)
    VALUES
(1, 'testname1', 21),
(2, 'testname2', 22),
(3, 'testname3', 23),
(4, 'testname4', 24),
(5, 'testname5', 25),
(6, 'testname6', 26),
(7, 'testname7', 27),
(8, 'testname8', 28),
(9, 'testname9', 29);



INSERT INTO orders (client_id) 
    VALUES
(1),
(3),
(2),
(2),
(1),
(5),
(4),
(4),
(4),
(6),
(7),
(2),
(3),
(2),
(4);


-- SELECT
--     c.name,
--     ARRAY_AGG(o.id)
-- FROM 
--     client AS c
-- INNER JOIN orders AS o ON c.id = o.client_id
-- GROUP BY c.id
-- ;


-- SELECT
--     COUNT(orders.client_id)
-- FROM 
--     orders
-- WHERE orders.client_id = 4
-- ;

SELECT c.name,Array_agg(o.id)
FROM client as c
INNER JOIN orders as o on o.client_id = c.id
WHERE c.age > 25
GROUP BY c.name 
HAVING COUNT(o) > 2;