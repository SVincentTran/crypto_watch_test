-- Table: public.users

-- DROP TABLE IF EXISTS public.users;

CREATE SEQUENCE IF NOT EXISTS public.users_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.users_id_seq
    OWNER TO postgres;

CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL DEFAULT nextval('public.users_id_seq'::regclass),
    user_id text COLLATE pg_catalog."default" NOT NULL,
    balance bigint DEFAULT 10000,
    eth_amount real DEFAULT 0,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;