CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null, 
    seguidor_id int not null,
    
    FOREIGN KEY(seguidor_id) REFEREnCES usuarios(id) ON DELETE CASCADE,
    FOREIGN KEY(usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    primary key(usuario_id, seguidor_id)
) ENGINE=INNODB;

GRANT ALL ON seguidores TO golang;
GRANT ALL ON usuarios TO golang;