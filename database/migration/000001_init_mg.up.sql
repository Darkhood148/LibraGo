CREATE DATABASE librago IF NOT EXISTS librago;
USE librago;

CREATE TABLE users (
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    fullName VARCHAR(255) NOT NULL,
    hash VARCHAR(255) NOT NULL,
    isAdmin BOOLEAN DEFAULT FALSE
);
CREATE TABLE books (
    bookID INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author VARCHAR(255),
    copiesAvailable SMALLINT UNSIGNED
);
CREATE TABLE checkouts (
    checkoutID INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    ofBook INT UNSIGNED,
    byUser VARCHAR(255),
    status ENUM('pending', 'issued', 'checkinDenied', 'checkinPending'),
    issueTime DATETIME,
    FOREIGN KEY (ofBook) REFERENCES books(bookID),
    FOREIGN KEY (byUser) REFERENCES users(username)
);
