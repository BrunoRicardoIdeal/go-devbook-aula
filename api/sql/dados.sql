insert into usuarios (nome, nick, email, senha) values
('Usuario1', 'usuario_1', 'usuario_1@gmail.com', '$2a$10$Jm2CfQV2jPstStCVHoH4de5TJUbgSwqO6RBtSBdojLXX1H6DC2HE.'),
('Usuario2', 'usuario_2', 'usuario_2@gmail.com', '$2a$10$Jm2CfQV2jPstStCVHoH4de5TJUbgSwqO6RBtSBdojLXX1H6DC2HE.'),
('Usuario3', 'usuario_3', 'usuario_3@gmail.com', '$2a$10$Jm2CfQV2jPstStCVHoH4de5TJUbgSwqO6RBtSBdojLXX1H6DC2HE.');

insert into seguidores values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id) values
('Pub usu 1', 'Texto pub 1', 1),
('Pub usu 2', 'Texto pub 2', 2),
('Pub usu 3', 'Texto pub 3', 3);