CREATE TABLE IF NOT EXISTS course (
    id              BIGINT          PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    title           VARCHAR(50)     NOT NULL,
    credits         INTEGER         NOT NULL,
    created_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS course_professor (
    id              BIGINT          PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    professor_id    BIGINT          NOT NULL REFERENCES professor (id) ON DELETE CASCADE,
    course_id       BIGINT          NOT NULL REFERENCES course (id) ON DELETE CASCADE,
    UNIQUE (professor_id, course_id)    
);