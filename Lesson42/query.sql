CREATE TABLE users (
    user_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    birthday TIMESTAMP,
    password VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE courses (
    course_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    title VARCHAR(100),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE lessons (
    lesson_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    course_id UUID REFERENCES courses(course_id),
    title VARCHAR(100),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE enrollments (
    enrollment_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID REFERENCES users(user_id),
    course_id UUID REFERENCES courses(course_id),
    enrollment_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

INSERT INTO users (name, email, birthday, password) VALUES
('John Doe', 'john.doe@example.com', '1990-01-01 00:00:00', 'password123'),
('Jane Smith', 'jane.smith@example.com', '1985-05-15 00:00:00', 'password456'),
('Alice Johnson', 'alice.johnson@example.com', '2000-12-12 00:00:00', 'password789');

INSERT INTO courses (title, description) VALUES
('Basic Spanish', 'Learn basic Spanish'),
('Intermediate French', 'Advance your French skills'),
('Beginner Japanese', 'Introduction to Japanese language');

INSERT INTO lessons (course_id, title, content) VALUES
((SELECT course_id FROM courses WHERE title = 'Basic Spanish'), 'Introduction to Spanish', 'This is the first lesson.'),
((SELECT course_id FROM courses WHERE title = 'Basic Spanish'), 'Spanish Greetings', 'This is the second lesson.'),
((SELECT course_id FROM courses WHERE title = 'Intermediate French'), 'French Basics', 'This is the first French lesson.'),
((SELECT course_id FROM courses WHERE title = 'Intermediate French'), 'French Pronunciation', 'This is the second French lesson.'),
((SELECT course_id FROM courses WHERE title = 'Beginner Japanese'), 'Japanese Alphabets', 'This is the first Japanese lesson.');

INSERT INTO enrollments (user_id, course_id, enrollment_date) VALUES
((SELECT user_id FROM users WHERE email = 'john.doe@example.com'), (SELECT course_id FROM courses WHERE title = 'Basic Spanish'), CURRENT_TIMESTAMP),
((SELECT user_id FROM users WHERE email = 'john.doe@example.com'), (SELECT course_id FROM courses WHERE title = 'Intermediate French'), CURRENT_TIMESTAMP),
((SELECT user_id FROM users WHERE email = 'jane.smith@example.com'), (SELECT course_id FROM courses WHERE title = 'Beginner Japanese'), CURRENT_TIMESTAMP),
((SELECT user_id FROM users WHERE email = 'alice.johnson@example.com'), (SELECT course_id FROM courses WHERE title = 'Basic Spanish'), CURRENT_TIMESTAMP);




