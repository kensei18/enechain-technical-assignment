create table
  companies (
    id uuid primary key default gen_random_uuid (),
    name text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz
  );

create table
  users (
    id uuid primary key default gen_random_uuid (),
    company_id uuid references companies on update no action on delete cascade not null,
    name text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz
  );

create type task_status as enum ('Todo', 'OnGoing', 'Done');

create table
  tasks (
    id uuid primary key default gen_random_uuid (),
    author_id uuid references companies on update no action on delete cascade not null,
    title text not null,
    description text not null,
    status task_status not null default 'Todo',
    is_private boolean not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz
  );

create table
  task_assignees (
    id uuid primary key default gen_random_uuid (),
    task_id uuid references tasks on update no action on delete cascade not null,
    assignee_id uuid references users on update no action on delete cascade not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz
  );
