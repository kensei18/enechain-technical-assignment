create table
  admin_users (
    id uuid primary key default gen_random_uuid (),
    name text not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz
  )
