create table if not exists public.notes
(
    id serial primary key,
    text varchar(255)
);

create user :DB_USR with encrypted password :'DB_PWD';
grant all privileges on database :DB_NAME to :DB_USR;
grant usage on schema public to :DB_USR;
grant all on all tables in schema public to :DB_USR;
grant all privileges on all sequences in schema public to :DB_USR;
