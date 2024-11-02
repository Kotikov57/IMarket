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
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, first_name, last_name, email, password) FROM stdin;
1	Dylan	Doe	sosiska12@gmail.com	12345678
3	Andrew	Tate	cutie_patootie@gmail.com	asdfghjk
2	Yayaye	Kokojambo	Chan_ramen_rulez@gmail.com	nyam-nyam
4	Nikolay	Lawson	king_of_world@gmail.com	qazwsxed
5	Glebchik	Dobryakov	sosiska14	bebebebababa
6	Chichi	Sobachka	gavgav@gmail.com	daypojrat
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, product_id, quantity, status, user_id, address) FROM stdin;
4	4	4	Delivered	3	Albukerke street,3
1	1	1	Registered	1	Pig's Cellar
5	1	5	Delivered	4	Pushkin street,5
3	3	3	Cancelled	2	Groove street,1
2	2	2	Paid	1	Peshkov streer,2
7	1	10	shipped	5	Malaha-Chapalaha street
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, name, price, quantity) FROM stdin;
2	Holodilnik	5000	20
3	Chainik	1500.1	15
4	Noutbuk	9999.99	50
5	Myshka	200	200
1	Utyug	1199.99	2
\.


--
-- PostgreSQL database dump complete
--

