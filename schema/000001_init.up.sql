CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);