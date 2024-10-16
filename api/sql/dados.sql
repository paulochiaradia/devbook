insert into usuarios (nome, nick, email, senha)
values
("usuario1", "user1", "user1@teste.com", "$2a$10$Cyxau6aM10caUSclePkm3ujRSRuL.537su6nsw8mM1HYxLKEcYJ9e"),
("usuario2", "user2", "user2@teste.com", "$2a$10$Cyxau6aM10caUSclePkm3ujRSRuL.537su6nsw8mM1HYxLKEcYJ9e"),
("usuario3", "user3", "user3@teste.com", "$2a$10$Cyxau6aM10caUSclePkm3ujRSRuL.537su6nsw8mM1HYxLKEcYJ9e");

insert into seguidores(usuario_id, seguidor_id)
values
(1,2),
(3,1),
(1,3);

insert into publicacoes(titulo, conteudo, autor_id)
values
    ("Publicacao do usuario 1", "Esse é o conteúdo da publicacao do usuario 1", 1),
    ("Publicacao do usuario 2", "Esse é o conteúdo da publicacao do usuario 2", 2),
    ("Publicacao do usuario 3", "Esse é o conteúdo da publicacao do usuario 3", 3);