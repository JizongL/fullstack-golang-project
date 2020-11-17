-- Create a `snippets` table.
CREATE DATABASE IF NOT EXISTS snippetbox;
USE snippetbox;
CREATE TABLE IF NOT EXISTS snippets (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, title VARCHAR(100) NOT NULL,
content TEXT NOT NULL,
created DATETIME NOT NULL,
expires DATETIME NOT NULL
);
-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);