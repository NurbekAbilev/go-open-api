create table employees (
    id serial primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    position_id int references positions(id)
);
