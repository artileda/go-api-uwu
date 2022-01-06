CREATE TABLE articles (
  id bigserial PRIMARY KEY,
  author_id integer NOT NULL,
  title varchar(64) NOT NULL,
  body varchar(64) NOT NULL,
  created_at timestamp NOT NULL
);

CREATE TABLE "authors" (
  id bigserial PRIMARY KEY,
  name varchar(64) NOT NULL
);