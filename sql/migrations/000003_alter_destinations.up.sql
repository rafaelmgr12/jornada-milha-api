START TRANSACTION;

-- Removendo a coluna Photo
ALTER TABLE destinations DROP COLUMN photo;

-- Adicionando a coluna Foto 1
ALTER TABLE destinations ADD COLUMN photo1 TEXT NOT NULL;

-- Adicionando a coluna Foto 2
ALTER TABLE destinations ADD COLUMN photo2 TEXT NOT NULL;

-- Adicionando a coluna Meta com limite de 160 caracteres
ALTER TABLE destinations ADD COLUMN meta VARCHAR(160) NOT NULL;

-- Adicionando a coluna Texto descritivo, que é opcional
ALTER TABLE destinations ADD COLUMN text_description TEXT NOT NULL;

COMMIT;
