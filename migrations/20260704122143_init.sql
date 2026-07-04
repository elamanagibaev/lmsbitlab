-- +goose Up
CREATE TABLE courses (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         description TEXT,
                         created_at TIMESTAMP NOT NULL DEFAULT now(),
                         updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE chapters (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          "order" INTEGER NOT NULL,
                          course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
                          created_at TIMESTAMP NOT NULL DEFAULT now(),
                          updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX idx_chapters_course_id ON chapters(course_id);

CREATE TABLE lessons (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         description TEXT,
                         content TEXT,
                         "order" INTEGER NOT NULL,
                         chapter_id INTEGER NOT NULL REFERENCES chapters(id) ON DELETE CASCADE,
                         created_at TIMESTAMP NOT NULL DEFAULT now(),
                         updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX idx_lessons_chapter_id ON lessons(chapter_id);

-- +goose Down
DROP TABLE lessons;
DROP TABLE chapters;
DROP TABLE courses;