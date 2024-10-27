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

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS '';


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
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id integer,
    product_id integer,
    quantity integer,
    status character varying(255)
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id integer,
    name character varying(255),
    price double precision,
    quantity integer
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255),
    password character varying(255)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: products products_pk; Type: CONSTRAINT; Schema: imarket_schema; Owner: postgres
--

ALTER TABLE ONLY imarket_schema.products
    ADD CONSTRAINT products_pk PRIMARY KEY (id);


--
-- Name: products unique_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT unique_id UNIQUE (id);


--
-- Name: orders unique_orders_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT unique_orders_id UNIQUE (id);


--
-- Name: products unique_products_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT unique_products_id UNIQUE (id);


--
-- Name: products unique_products_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT unique_products_name UNIQUE (name);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: orders orders_orders__fk; Type: FK CONSTRAINT; Schema: imarket_schema; Owner: postgres
--

ALTER TABLE ONLY imarket_schema.orders
    ADD CONSTRAINT orders_orders__fk FOREIGN KEY (product_id) REFERENCES imarket_schema.products(id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;


--
-- PostgreSQL database dump complete
--

