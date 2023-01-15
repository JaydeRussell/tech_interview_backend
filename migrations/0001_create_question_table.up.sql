CREATE TABLE IF NOT EXISTS question (
   	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    body varchar(280) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted bool DEFAULT false NOT NULL
);