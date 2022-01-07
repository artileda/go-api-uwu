BEGIN;
  CREATE TABLE IF NOT EXISTS articles (
    id bigserial PRIMARY KEY,
    author_id integer NOT NULL,
    title varchar(64) NOT NULL,
    body text NOT NULL,
    created_at timestamp NOT NULL
  );

  ALTER TABLE articles ADD COLUMN article_fts tsvector 
      GENERATED ALWAYS AS 
          (setweight(to_tsvector('english', coalesce(title, '')), 'A') || 
          setweight(to_tsvector('english', coalesce(body, '')), 'B')) STORED;

  CREATE TABLE IF NOT EXISTS authors (
    id bigserial PRIMARY KEY,
    name varchar(64) NOT NULL
  );

  INSERT INTO authors(name) values ('Anya');
  INSERT INTO authors(name) values ('Bela');
  INSERT INTO authors(name) values ('Celia');
COMMIT;