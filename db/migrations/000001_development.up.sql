CREATE SCHEMA IF NOT EXISTS development;

-- CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS development.operators
(
    id                  uuid DEFAULT gen_random_uuid() NOT NULL ,
    created_at          timestamp                      NOT NULL,
    modified_at         timestamp                      NOT NULL,
    first_name          varchar(55)                    NOT NULL,
    last_name           varchar(55)                    NOT NULL,
    middle_name         varchar(55)                    NOT NULL,
    city                varchar(55)                    NOT NULL,
    country_code_number varchar(6)                     NOT NULL,
    phone_number        text                           NOT NULL,
    email               varchar(55) UNIQUE             NOT NULL,
    password            text                           NOT NULL
);

CREATE TYPE public.project_type AS ENUM ('in', 'out', 'auto');

CREATE TABLE IF NOT EXISTS development.projects
(
    id           uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at   timestamp                      NOT NULL,
    modified_at  timestamp                      NOT NULL,
    project_name varchar(100) UNIQUE            NOT NULL,
    project_type project_type                   NOT NULL
);

-- CREATE TABLE IF NOT EXISTS development.project_operators
-- (
--     id          uuid DEFAULT gen_random_uuid() NOT NULL,
--     project_id  uuid                           NOT NULL,
--     operator_id uuid                           NOT NULL,
--     created_at  timestamp                      NOT NULL,
--     modified_at timestamp                      NOT NULL
-- );


ALTER TABLE development.operators
    ADD PRIMARY KEY (id) ;
ALTER TABLE development.projects
    ADD PRIMARY KEY (id);

CREATE INDEX ON development.operators (id);
CREATE INDEX ON development.projects (id);

-- ALTER TABLE development.project_operators
--     ADD CONSTRAINT fk_project_operators_project_id
--         FOREIGN KEY (project_id)
--             REFERENCES development.projects (id);

-- ALTER TABLE development.project_operators
--     ADD CONSTRAINT fk_project_operators_operator_id
--         FOREIGN KEY (operator_id)
--             REFERENCES development.operators (id);

ALTER TABLE development.operators
    ADD CONSTRAINT phone_number_chk
        CHECK (phone_number ~ '^\d{10}$');

CREATE UNIQUE INDEX uq_country_code_phone_number_operators
    ON development.operators (country_code_number, phone_number);


CREATE TABLE development.project_operators
(
    project_id  uuid REFERENCES development.projects (id) ON UPDATE CASCADE ON DELETE CASCADE,
    operator_id uuid REFERENCES development.operators (id) ON UPDATE CASCADE  ON DELETE CASCADE,
    CONSTRAINT project_operators_pkey PRIMARY KEY (project_id, operator_id)
);

-- CREATE INDEX ON development.project_operators (project_id, operator_id);

CREATE UNIQUE INDEX uq_project_operators_ids
    ON development.project_operators (project_id, operator_id);



