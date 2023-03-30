CREATE DATABASE IF NOT EXISTS musicbot;

USE musicbot;

DROP TABLE IF EXISTS musicas_playlist;
DROP TABLE IF EXISTS playlist;
DROP TABLE IF EXISTS musicas;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id INT auto_increment PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL,
    senha VARCHAR(100) NOT NULL,
    criadoEM TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
)ENGINE = INNODB;


CREATE TABLE musicas(
    nome VARCHAR(400) NOT NULL PRIMARY KEY,
    caminho VARCHAR(400) NOT NULL UNIQUE,
    duracao VARCHAR(15),
    adicionadoPor INT NOT NULL,
    genero VARCHAR(20) NOT NULL,
    FOREIGN KEY (adicionadoPor) REFERENCES usuarios(id)
)ENGINE = INNODB;


CREATE TABLE playlist(
    id INT auto_increment PRIMARY KEY,
    nome VARCHAR(50),
    usuario INT,
    criadoEM TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    descricao VARCHAR(250),
    FOREIGN KEY (usuario) REFERENCES usuarios(id)
    
)ENGINE = INNODB;


CREATE TABLE musicas_playlist(
    id INT auto_increment PRIMARY KEY,
    nome VARCHAR(400) NOT NULL,
    adicionadoPlaylist INT NOT NULL,
    posicao INT NOT NULL,
    criadoEM TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    playlist INT,
    FOREIGN KEY (nome) REFERENCES musicas(nome),
    FOREIGN KEY (adicionadoPlaylist) REFERENCES usuarios(id),
    FOREIGN KEY (playlist) REFERENCES playlist(id)
)