CREATE TABLE teachers
(
    id serial not null unique,
    name varchar(88) not null,
    email varchar(20) not null unique,
    password_hash varchar(88) not null
);

CREATE TABLE students
(
    id serial not null unique,
    teacher_id int references teachers(id) on delete cascade not null,
    name varchar(88) not null
);

CREATE TABLE transactions
(
    id serial not null unique,
    student_id int references students(id) on delete cascade not null,
    sum int not null
);