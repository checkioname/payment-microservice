-- BANCO DE DADOS PAGAMENTO

CREATE DATABASE IF NOT EXISTS pagamentos;

CREATE TABLE IF NOT EXISTS pagamentos.db (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  pedido_id INTEGER NOT NULL UNIQUE, -- Geralmente 1 pagamento por pedido
  data_processamento TEXT NOT NULL,
  status TEXT NOT NULL, -- Ex: 'aprovado', 'rejeitado', 'pendente'
  metodo TEXT,
  FOREIGN KEY (pedido_id) REFERENCES pedidos (id)
);

CREATE DATABASE IF NOT EXISTS produtos;

-- sql_create_produtos_table =
CREATE TABLE IF NOT EXISTS produtos.db(
id INTEGER PRIMARY KEY,
nome TEXT NOT NULL,
preco REAL NOT NULL,
quantidade_estoque INTEGER NOT NULL
);

CREATE DATABASE IF NOT EXISTS clientes

-- sql_create_clientes_table =
CREATE TABLE IF NOT EXISTS clientes.db (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

CREATE DATABASE IF NOT EXISTS pedidos

-- sql_create_pedidos_table =
CREATE TABLE IF NOT EXISTS pedidos.db (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cliente_id INTEGER NOT NULL,
    data_criacao TEXT NOT NULL,
    status TEXT NOT NULL, -- Ex: 'pendente', 'pagamento_aprovado', 'em_separacao', 'nf_emitida', 'enviado', 'cancelado'
    valor_total REAL,
    FOREIGN KEY (cliente_id) REFERENCES clientes (id)
);

CREATE DATABASE IF NOT EXISTS itens_pedido

-- sql_create_itens_pedido_table =
CREATE TABLE IF NOT EXISTS itens_pedido.db (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    pedido_id INTEGER NOT NULL,
    produto_id INTEGER NOT NULL,
    quantidade INTEGER NOT NULL,
    preco_unitario REAL NOT NULL,
    FOREIGN KEY (pedido_id) REFERENCES pedidos (id),
    FOREIGN KEY (produto_id) REFERENCES produtos (id)
-- sql_create_notas_fiscais_table =
);

CREATE DATABASE IF NOT EXISTS notas_fiscais

CREATE TABLE IF NOT EXISTS notas_fiscais.db (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    pedido_id INTEGER NOT NULL UNIQUE,
    numero TEXT NOT NULL UNIQUE,
    data_emissao TEXT NOT NULL,
    chave_acesso TEXT,
    FOREIGN KEY (pedido_id) REFERENCES pedidos (id)
);

CREATE DATABASE IF NOT EXISTS envios
-- sql_create_envios_table
CREATE TABLE IF NOT EXISTS envios.db (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    pedido_id INTEGER NOT NULL UNIQUE,
    nota_fiscal_id INTEGER NOT NULL UNIQUE,
    data_despacho TEXT,
    codigo_rastreamento TEXT,
    status TEXT NOT NULL, -- Ex: 'aguardando_envio', 'enviado', 'entregue'
    FOREIGN KEY (pedido_id) REFERENCES pedidos (id),
    FOREIGN KEY (nota_fiscal_id) REFERENCES notas_fiscais (id)
);