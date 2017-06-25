-- Over our ../147/7instructions database:

-- Alter use "agent" so it can write/read
ALTER USER agent WITH SUPERUSER;

-- Add books table
CREATE TABLE books (
  ISBN CHAR(14) PRIMARY KEY NOT NULL,
  title varchar(255) NOT NULL,
  author varchar(255) NOT NULL,
  price decimal(5,2) NOT NULL
);


-- Insert some recors to be readed
INSERT INTO books(isbn, title, author, price) VALUES
('B1', 'The Republic', 'Plato', 2.50),
('B2', 'Atlas Shrugged', 'Ayn Rand', 3.00),
('B3', 'Anatomy of the state', 'Murray Rothbard', 2.75),
('B4', 'Historia de la conquista de Nueva Espana', 'Bartolome de las Casas', 3.50),
('B5', 'The Virtue of Selfishness', 'Ayn Rand', 2.25),
('B6', 'Basic economics', 'Thomas Sowell', 3.00),
('B7', 'Democracy in America', 'Alexis de Tocqueville', 4.50);
