CREATE INDEX users_company_id_key ON users (company_id);

CREATE INDEX tasks_company_id_key ON tasks (company_id);

CREATE INDEX tasks_author_id_key ON tasks (author_id);

CREATE INDEX task_assignees_task_id_key ON task_assignees (task_id);
