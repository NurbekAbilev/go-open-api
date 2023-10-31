alter table employees
    add constraint fk_employees_positions foreign key (position_id) references positions (id)