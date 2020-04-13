CREATE TABLE books(
    id int UNIQUE PRIMARY KEY NOT NULL, 
    isbn varchar(20),
    title varchar(20),
    author varchar(20),
    createdAt TIMESTAMP WITH TIME ZONE,
    updatedAt TIMESTAMP WITH TIME ZONE
)