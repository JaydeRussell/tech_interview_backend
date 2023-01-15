BEGIN;

DELETE FROM question
where id in 
('3687fdc0-6d02-4c84-891e-71f3760fb603','490b8a6d-0a22-4cff-92fa-557ee3dcf99f','cb106909-3064-4756-beb5-cd19c03d1e92','8d41da29-7965-4de4-807b-5f8d5f1fe416');

DELETE FROM answer
where question_id in
('3687fdc0-6d02-4c84-891e-71f3760fb603','490b8a6d-0a22-4cff-92fa-557ee3dcf99f','cb106909-3064-4756-beb5-cd19c03d1e92','8d41da29-7965-4de4-807b-5f8d5f1fe416');

COMMIT;