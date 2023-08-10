create table if not exists users (
	id integer unsigned auto_increment primary key,
	email varchar(100) not null,
	password varchar(100) not null,
	name varchar(100) default '',
	image_url varchar(255) default '',
	intro text default '',
	created_at datetime
);

-- data.goのUserTestDataの値をインサートしている
insert into users (email, password, name, image_url, intro, created_at) values
	('test@test.com', 'hashedPassword', 'testUser', 'testImage.png', 'テスト用自己紹介文です。', now());