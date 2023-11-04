create table employees (
    id serial primary key,
    "login" varchar(255) not null,
    "password" varchar(255) not null,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    position_id int references positions(id)
);
