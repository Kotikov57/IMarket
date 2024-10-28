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

COPY public.users (id, first_name, last_name, address, login, password) FROM stdin;
1	Dylan	Doe	Grove street, 1	sosiska12	12345678
3	Andrew	Tate	Pushkin street, 3	cutie_patootie	asdfghjk
4	Nikolay	Lawson	Pupkin street, 4	king_of_world	qazwsxed
5	Karlos	Halapenjo	Albukerke street,5	ricardo_4ever	gachiflex
2	Yayaye	Kokojambo	Pig's Cellar	Chan_ramen_rulez	nyam-nyam
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, product_id, quantity, status, user_id) FROM stdin;
3	3	3	Cancelled	2
2	2	2	Paid	1
4	4	4	Delivered	3
1	1	1	Registered	1
5	1	5	Delivered	4
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, name, price, quantity) FROM stdin;
1	Utyug	1199.99	12
2	Holodilnik	5000	20
3	Chainik	1500.1	15
4	Noutbuk	9999.99	50
5	Myshka	200	200
\.


--
-- PostgreSQL database dump complete
--

