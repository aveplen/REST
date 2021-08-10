CREATE TYPE role_enum AS ENUM ('watcher', 'student', 'admin');
CREATE TYPE gender_enum AS ENUM ('male', 'female');

CREATE TABLE students (
        student_id              BIGSERIAL       NOT NULL        PRIMARY KEY,
        first_name              VARCHAR(255)    NOT NULL,
        second_name             VARCHAR(255)    NOT NULL,
        gender                  gender_enum     NOT NULL,
        group_number            VARCHAR(255)    NOT NULL,
        graduation_year         INT             NOT NULL,
        exam_score              INT             NOT NULL,
        additional_score        SMALLINT        NOT NULL
);

CREATE TABLE users (
        user_id                 BIGSERIAL       NOT NULL        PRIMARY KEY,
        email                   VARCHAR(255)    NOT NULL        UNIQUE,
        encrypted_password      VARCHAR(200)    NOT NULL,
        role                    role_enum       NOT NULL,
        stud_id                 BIGINT,

        CONSTRAINT stud_id_fk FOREIGN KEY (stud_id)
        REFERENCES public.students (student_id) MATCH SIMPLE
        ON DELETE CASCADE
);