-- public.usuarios_seq definition

-- DROP SEQUENCE public.usuarios_seq;

CREATE SEQUENCE public.usuarios_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;


-- public.usuarios definition

-- Drop table

-- DROP TABLE public.usuarios;

CREATE TABLE public.usuarios (
	id serial4 NOT NULL,
	nome varchar(50) NOT NULL,
	nick varchar(50) NOT NULL,
	email varchar(50) NOT NULL,
	senha varchar(20) NOT NULL,
	criadoem timestamp default current_timestamp,
	CONSTRAINT usuarios_pk PRIMARY KEY (id),
	CONSTRAINT usuarios_unique UNIQUE (nick)
);