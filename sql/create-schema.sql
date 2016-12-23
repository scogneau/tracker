CREATE SEQUENCE people_id_seq;

CREATE TABLE people (
  id int CONSTRAINT people_primary_key PRIMARY KEY DEFAULT nextval('people_id_seq'),
  nom text NOT NULL,
  prenom text NOT NULL);
