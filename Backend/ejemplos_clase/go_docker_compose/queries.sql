CREATE TABLE IF NOT EXISTS comentarios(
  id SERIAL PRIMARY KEY,
  time TIMESTAMP,
  comment TEXT,
  reactions INT
);

INSERT INTO comentarios (time, comment, reactions) 
VALUES
('2023-10-14 07:33:12', 'Estaban pasando bueno, no? xD', 10),
('2024-04-29 23:59:59', 'Felicitaciones, se ven muy bien juntos', 5),
('2022-12-30 12:11:45', 'Happy New Year!', 20),
('2024-04-29 23:59:59', 'You look great!', 2) RETURNING id;
