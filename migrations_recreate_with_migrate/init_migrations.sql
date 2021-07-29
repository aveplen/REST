CREATE TABLE scores (
    score_id            BIGSERIAL       NOT NULL    PRIMARY KEY
,   mathematics         SMALLINT
,   russian             SMALLINT
,   physics             SMALLINT
,   computer_science    SMALLINT
,   literature          SMALLINT
,   social_science      SMALLINT
,   history             SMALLINT
,   biology             SMALLINT
,   geography_science   SMALLINT

,   CONSTRAINT adequate_mathematics_score CHECK (mathematics
        IS NULL OR ((mathematics >= 0) AND (mathematics <= 100)))

,   CONSTRAINT adequate_russian_score CHECK (russian
        IS NULL OR ((russian >= 0) AND (russian <= 100)))

,   CONSTRAINT adequate_physics_score CHECK (physics
        IS NULL OR ((physics >= 0) AND (physics <= 100)))

,   CONSTRAINT adequate_computer_science_s CHECK (computer_science
        IS NULL OR ((computer_science >= 0) AND (computer_science <= 100)))

,   CONSTRAINT adequate_literature_score CHECK (literature
        IS NULL OR ((literature >= 0) AND (literature <= 100)))

,   CONSTRAINT adequate_social_science_sco CHECK (social_science
        IS NULL OR ((social_science >= 0) AND (social_science <= 100)))

,   CONSTRAINT adequate_history_score CHECK (history
        IS NULL OR ((history >= 0) AND (history <= 100)))

,   CONSTRAINT adequate_biology_score CHECK (biology
        IS NULL OR ((biology >= 0) AND (biology <= 100)))

,   CONSTRAINT adequate_geography_science_score CHECK (geography_science
        IS NULL OR ((geography_science >= 0) AND (geography_science <= 100)))
);

CREATE TABLE cities (
    city_id             BIGSERIAL       NOT NULL    PRIMARY KEY
,   city_name           VARCHAR(255)    NOT NULL

,   CONSTRAINT unique_city_name UNIQUE (city_name)
);

CREATE TABLE schools (
    school_id           BIGSERIAL       NOT NULL    PRIMARY KEY
,   school_number       VARCHAR(25)     NOT NULL
,   city_id             BIGINT          NOT NULL
,   geo_address         VARCHAR(255)    NOT NULL

,   CONSTRAINT unique_school_number UNIQUE (school_number)

,   CONSTRAINT unique_geo_address UNIQUE (geo_address)

,   CONSTRAINT city_id_fk FOREIGN KEY (city_id)
        REFERENCES public.cities (city_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);


CREATE TABLE cridentials (
    cridentials_id      BIGSERIAL       NOT NULL    PRIMARY KEY
,   first_name          VARCHAR(255)    NOT NULL
,   second_name         VARCHAR(255)    NOT NULL
,   gender              VARCHAR(6)      NOT NULL
,   date_of_birth       DATE            NOT NULL
-- add NOT NULL for authentication
-- ,   email               VARCHAR(255)
-- add salted password for authentication

-- ,   CONSTRAINT unique_email UNIQUE (email)
);

CREATE TABLE students (
    student_id          BIGSERIAL       NOT NULL    PRIMARY KEY
,   score_id            BIGINT          NOT NULL
,   school_id           BIGINT          NOT NULL
,   cridentials_id       BIGINT          NOT NULL

,   CONSTRAINT score_id_fk FOREIGN KEY (score_id)
        REFERENCES public.scores (score_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE

,   CONSTRAINT school_id_fk FOREIGN KEY (school_id)
        REFERENCES public.schools (school_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT

,   CONSTRAINT cridentials_id_fk FOREIGN KEY (cridentials_id)
        REFERENCES public.cridentials (cridentials_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);