CREATE TABLE IF NOT EXISTS answer (
   	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    question_id uuid NOT NULL ,
    body varchar(280) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted bool DEFAULT false NOT NULL,
    CONSTRAINT fk_question
      FOREIGN KEY(question_id)
	  REFERENCES question(id)
);