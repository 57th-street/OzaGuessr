create table if not exists users (
	id integer unsigned auto_increment primary key,
	email varchar(100) not null,
	password varchar(100) not null,
	name varchar(100) default '',
	image_url varchar(255) default '',
	intro text default '',
	created_at datetime
);
