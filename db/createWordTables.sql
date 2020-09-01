CREATE TABLE IF NOT EXISTS public.words(
    id serial PRIMARY KEY,
    word varchar(50) NOT NULL,
    is_palindrom boolean NOT NULL
);