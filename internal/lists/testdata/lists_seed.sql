-- Reset table to known state
DELETE FROM lists;

-- Insert predictable test rows
INSERT INTO lists (id, name) VALUES (1, 'todo');
INSERT INTO lists (id, name) VALUES (2, 'work');
INSERT INTO lists (id, name) VALUES (3, 'personal');

