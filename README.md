sudo -u postgres psql
\l - list all DBs
\c - use db
\dt - tables info
CREATE TABLE IF NOT EXISTS "commit_type" (
   "type_id" SERIAL PRIMARY KEY,
   "type_name" VARCHAR(256) NOT NULL CHECK (LENGTH(type_name) > 0)
);

CREATE TABLE IF NOT EXISTS "commits" (
    "c_id" SERIAL PRIMARY KEY,
    "type_id" INTEGER REFERENCES commit_type (type_id),
    "commit_text" VARCHAR NOT NULL CHECK (LENGTH(type_text) > 0)
);


ALTER TABLE commits 
RENAME COLUMN type_text TO commit_text;

INSERT INTO commit_type (type_name) VALUES ('gachi');
INSERT INTO commit_type (type_name) VALUES ('uwu');

INSERT INTO commits (type_id, commit_text)
VALUES 
(1, 'fucking slave was merged'),
(1, 'master of this gym fixed fucking disgusting bug'),
(2, 'OwO another UWUshing very nice bugfix meow!!!!');