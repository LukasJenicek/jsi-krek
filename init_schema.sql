DROP TABLE IF EXISTS "surveys";
DROP TABLE IF EXISTS "survey_questions";
DROP TABLE IF EXISTS "survey_answers";
DROP TABLE IF EXISTS "surveys";
DROP TABLE IF EXISTS "responds";

CREATE TABLE "surveys"
(
    "id" INTEGER PRIMARY KEY
);

CREATE TABLE "survey_questions"
(
    "id"        INTEGER PRIMARY KEY,
    "question"  TEXT    NOT NULL,
    "survey_id" INTEGER NOT NULL,
    FOREIGN KEY ("survey_id") REFERENCES surveys (id)
);

CREATE TABLE "survey_answers"
(
    "id"                 INTEGER PRIMARY KEY, -- works as an alias to rowid
    "answer"             TEXT    NOT NULL,
    "survey_question_id" INTEGER NOT NULL,
    FOREIGN KEY (survey_question_id) REFERENCES survey_questions (id)

);

CREATE TABLE "distributed_surveys"
(
    "id"                 INTEGER PRIMARY KEY,
    "created_at"         TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "finished"           BOOLEAN                     NOT NULL DEFAULT FALSE,
    "survey_question_id" INTEGER                     NOT NULL,
    FOREIGN KEY ("survey_question_id") REFERENCES survey_questions (id)
);

CREATE TABLE "responds"
(
    "id"                    INTEGER PRIMARY KEY,
    "distributed_survey_id" INTEGER                     NOT NULL,
    "created_at"            TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "who"                   VARCHAR                     NOT NULL,
    "answer"                TEXT                        NOT NULL,
    FOREIGN KEY (distributed_survey_id) REFERENCES distributed_surveys (id)
);

INSERT INTO "surveys" (id)
VALUES (1),
       (2),
       (3),
       (4);

INSERT INTO "survey_questions" (id, question, survey_id)
VALUES (1, 'Krek ?', 1),
       (2, 'Krek?', 2),
       (3, 'isKrek?', 3),
       (4, 'krekd?', 4);

INSERT INTO "survey_answers" (id, answer, survey_question_id)
VALUES (1, 'Yes', 1),
       (2, 'No', 1),
       (3, 'Ne', 1),
       (4, 'Megakrek', 2),
       (5, 'Classic krek', 2),
       (6, 'Jednou do roka jsem takhle moc krek', 2),
       (7, 'Dve kafe abych vubec videl krek', 2),
       (8, 'True', 3),
       (9, 'False', 3),
       (10, 'Tu', 4),
       (11, 'Tru', 4),
       (12, 'True', 4);



