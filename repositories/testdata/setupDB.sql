CREATE TABLE IF NOT EXISTS users (
	id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	email VARCHAR(100) NOT NULL UNIQUE,
	password VARCHAR(100) NOT NULL,
	name VARCHAR(100) DEFAULT '',
	image_url VARCHAR(255) DEFAULT '',
	intro TEXT DEFAULT '',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS themes (
	id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS quizzes (
	id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	quest VARCHAR(255) NOT NULL,
	answer VARCHAR(100) NOT NULL,
	descript TEXT,
	theme_id INT UNSIGNED,
	FOREIGN KEY (theme_id) REFERENCES themes(id) ON DELETE CASCADE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- data.goのUserTestDataの値をインサートしている
INSERT INTO users (email, password, name, image_url, intro, created_at) 
VALUES
	('test@test.com', 'hashedPassword', 'testUser', 'testImage.png', 'テスト用自己紹介文です。', NOW());

INSERT INTO themes (name) VALUES ('尾崎豊'), ('歴史');

INSERT INTO quizzes (quest, answer, descript, theme_id) 
VALUES 
	('ozaki_tokyo_dome.png', '1988', '東京ドーム公演の様子である。キャリア唯一であると同時に自身最多の5万6000人を動員した。', 1),
	('ozaki_nihon_senenkan.png', '1985', '日本青年館でのライブ模様である。尾崎にとって東京での初めてのホールコンサートとなった。', 1),
	('honnoj_incident.png', '1582', '本能寺の変の様子である。早朝、明智光秀が謀反を起こし、京都本能寺に滞在する主君・織田信長を襲撃した事件である。', 2);