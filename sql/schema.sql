--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: imarket_schema; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA imarket_schema;


ALTER SCHEMA imarket_schema OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: orders; Type: TABLE; Schema: imarket_schema; Owner: postgres
--

CREATE TABLE imarket_schema.orders (
    id integer,
    product_id integer,
    quantity integer,
    status character varying(255)
);


ALTER TABLE imarket_schema.orders OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: imarket_schema; Owner: postgres
--

CREATE TABLE imarket_schema.products (
    id integer NOT NULL,
    name character varying(255),
    price double precision,
    quantity integer
);


ALTER TABLE imarket_schema.products OWNER TO postgres;

--
-- Name: products products_pk; Type: CONSTRAINT; Schema: imarket_schema; Owner: postgres
--

ALTER TABLE ONLY imarket_schema.products
    ADD CONSTRAINT products_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

