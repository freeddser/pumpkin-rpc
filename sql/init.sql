

CREATE TABLE public."customer"
(
    id bigint NOT NULL,
    name character varying(200) COLLATE pg_catalog."default",
    email character varying(200) COLLATE pg_catalog."default",
    phone character varying(200) COLLATE pg_catalog."default",
    CONSTRAINT customer_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public."customer"
    OWNER to postgres;