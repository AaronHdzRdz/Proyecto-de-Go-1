create table users (
    id serial primary key,
    name varchar(50) not null,
    email varchar(100) unique not null,
    created_at timestamptz default now()
);