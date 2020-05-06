CREATE TABLE IF NOT EXISTS project
(
    project_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    project_created_at timestamp without time zone NOT NULL,
    project_status integer NOT NULL,
    project_title text NOT NULL,
    project_description text,
    CONSTRAINT project_pkey PRIMARY KEY (project_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE project
    OWNER to development;

INSERT INTO project(
	project_created_at, project_status, project_title, project_description)
	VALUES ('2020-04-09T08:12:22.811908', 1, 'Test Project 1', 'Description for Test Project 1');

INSERT INTO project(
	project_created_at, project_status, project_title, project_description)
	VALUES ('2020-04-09T08:12:22.811908', 1, 'Test Project 2', 'Description for Test Project 2');

INSERT INTO project(
	project_created_at, project_status, project_title, project_description)
	VALUES ('2020-04-09T08:12:22.811908', 1, 'Test Project 3', 'Description for Test Project 3');

INSERT INTO project(
	project_created_at, project_status, project_title, project_description)
	VALUES ('2020-04-09T08:12:22.811908', 1, 'Test Project 4', 'Description for Test Project 4');

INSERT INTO project(
	project_created_at, project_status, project_title, project_description)
	VALUES ('2020-04-09T08:12:22.811908', 1, 'Test Project 5', 'Description for Test Project 5');