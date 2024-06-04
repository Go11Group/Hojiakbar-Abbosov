CREATE TABLE Employees (
    employee_id INT PRIMARY KEY,
    name VARCHAR(50),
    department_id INT,
    age INT
);

CREATE TABLE Departments (
    department_id INT PRIMARY KEY,
    department_name VARCHAR(50)
);

INSERT INTO Employees (employee_id, name, department_id, age) VALUES
(1, 'Alice', 1, 30),
(2, 'Bob', 2, 25),
(3, 'Charlie', 1, 35),
(4, 'David', 3, 40),
(5, 'Eve', 2, 45);

INSERT INTO Departments (department_id, department_name) VALUES
(1, 'HR'),
(2, 'Engineering'),
(3, 'Marketing');
