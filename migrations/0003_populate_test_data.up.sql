BEGIN;

INSERT INTO question
(id, body, deleted)
VALUES
('3687fdc0-6d02-4c84-891e-71f3760fb603', 'Is this the body of a question?', true),
('490b8a6d-0a22-4cff-92fa-557ee3dcf99f', 'What is Lorem Ipsum?', false),
('cb106909-3064-4756-beb5-cd19c03d1e92', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi eu laoreet nibh. Donec eget nunc lacus. Nunc vehicula nunc velit, in fringilla nibh imperdiet et. In hac habitasse platea dictumst. Quisque egestas velit quis odio egestas, sit amet interdum urna sollicitudin. Vestibu', false),
('8d41da29-7965-4de4-807b-5f8d5f1fe416', '', false);

INSERT INTO answer
(question_id, body, deleted)
VALUES
('3687fdc0-6d02-4c84-891e-71f3760fb603', 'Yes, that was the body of a question. Now this is the body of an answer.', false),
('3687fdc0-6d02-4c84-891e-71f3760fb603', 'No offense, but that was a really silly question', false),
('3687fdc0-6d02-4c84-891e-71f3760fb603', 'Something incredibly rude that got deleted', true),

('490b8a6d-0a22-4cff-92fa-557ee3dcf99f', 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five cent', false),
('490b8a6d-0a22-4cff-92fa-557ee3dcf99f', 'basically just garbage text used to show how text will look', false),
('490b8a6d-0a22-4cff-92fa-557ee3dcf99f', 'What they said ^^', false),
('490b8a6d-0a22-4cff-92fa-557ee3dcf99f', 'Something incredibly rude that got deleted', true),

('8d41da29-7965-4de4-807b-5f8d5f1fe416', 'You forgot to type your question', false);

COMMIT;