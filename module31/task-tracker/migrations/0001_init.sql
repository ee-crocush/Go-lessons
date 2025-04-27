-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
COMMENT ON TABLE users IS 'Пользователи';
COMMENT ON COLUMN users.id IS 'Идентификатор пользователя';
COMMENT ON COLUMN users.name IS 'Имя пользователя';
-- Таблица задач
CREATE TABLE IF NOT EXISTS tasks
(
    id          SERIAL PRIMARY KEY,
    title       TEXT    NOT NULL,
    content     TEXT    NOT NULL,
    author_id   INTEGER NOT NULL,
    assigned_id INTEGER NOT NULL,
    opened      BIGINT,
    closed      BIGINT,
    CONSTRAINT tasks_fk_author_id
        FOREIGN KEY (author_id)
            REFERENCES users (id)
            ON DELETE CASCADE,
    CONSTRAINT tasks_fk_assigned_id
        FOREIGN KEY (assigned_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);
COMMENT ON TABLE tasks IS 'Задачи';
COMMENT ON COLUMN tasks.id IS 'Идентификатор задачи';
COMMENT ON COLUMN tasks.title IS 'Название задачи';
COMMENT ON COLUMN tasks.content IS 'Описание задачи';
COMMENT ON COLUMN tasks.author_id IS 'Автор задачи';
COMMENT ON COLUMN tasks.assigned_id IS 'Ответственный за выполнение задачи';
COMMENT ON COLUMN tasks.opened IS 'Дата открытия задачи';
COMMENT ON COLUMN tasks.closed IS 'Дата закрытия задачи';
-- Таблица Лейблы
CREATE TABLE IF NOT EXISTS labels
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
COMMENT ON TABLE labels IS 'Лейблы';
COMMENT ON COLUMN labels.id IS 'Идентификатор лейбла';
COMMENT ON COLUMN labels.name IS 'Наименование лейбла';
-- Таблица Привязки задач к лейблам
CREATE TABLE IF NOT EXISTS tasks_labels
(
    task_id  INTEGER NOT NULL,
    label_id INTEGER NOT NULL,
    CONSTRAINT pk_tasks_label PRIMARY KEY (task_id, label_id),
    CONSTRAINT fk_tasks_label_task_id
        FOREIGN KEY (task_id)
            REFERENCES tasks (id)
            ON DELETE CASCADE,
    CONSTRAINT fk_tasks_label_label_id
        FOREIGN KEY (label_id)
            REFERENCES labels
            ON DELETE CASCADE
);
COMMENT ON TABLE tasks_labels IS 'Привязка задач к лейблам';
COMMENT ON COLUMN tasks_labels.task_id IS 'Идентификатор задачи';
COMMENT ON COLUMN tasks_labels.label_id IS 'Идентификатор лейбла';