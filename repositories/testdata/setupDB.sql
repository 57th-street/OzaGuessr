create table if not exists users (
	id integer unsigned auto_increment primary key,
	email varchar(100) not null,
	password varchar(100) not null,
	created_at datetime
);

insert into users (email, password, created_at) values
	('test@test.com', '$2a$10$qf.9gh7iLVLkqZ1FcDggReI3RaLK8fpztigA8QVJbW9IdH4hZ9j2W', now());