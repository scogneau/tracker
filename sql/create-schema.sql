create user user_test with password 'passw0rd' ;
GRANT ALL on tracker_test to user_test;

CREATE TABLE people (id int,nom text,prenom text);
