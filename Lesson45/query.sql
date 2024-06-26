CREATE TABLE books (
    book_id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year_published INT NOT NULL
);

CREATE TABLE borrows (
    borrow_id UUID PRIMARY KEY,
    book_id UUID REFERENCES books(book_id),
    user_id VARCHAR(255) NOT NULL,
    borrowed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO books (book_id, title, author, year_published) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'The Catcher in the Rye', 'J.D. Salinger', 1951),
('550e8400-e29b-41d4-a716-446655440001', 'To Kill a Mockingbird', 'Harper Lee', 1960),
('550e8400-e29b-41d4-a716-446655440002', '1984', 'George Orwell', 1949),
('550e8400-e29b-41d4-a716-446655440003', 'Pride and Prejudice', 'Jane Austen', 1813),
('550e8400-e29b-41d4-a716-446655440004', 'The Great Gatsby', 'F. Scott Fitzgerald', 1925);

INSERT INTO borrows (borrow_id, book_id, user_id) VALUES
('550e8400-e29b-41d4-a716-446655440100', '550e8400-e29b-41d4-a716-446655440000', 'user-001'),
('550e8400-e29b-41d4-a716-446655440101', '550e8400-e29b-41d4-a716-446655440002', 'user-002'),
('550e8400-e29b-41d4-a716-446655440102', '550e8400-e29b-41d4-a716-446655440001', 'user-003');