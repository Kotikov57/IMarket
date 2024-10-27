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
-- Data for Name: products; Type: TABLE DATA; Schema: imarket_schema; Owner: postgres
--

COPY imarket_schema.products (id, name, price, quantity) FROM stdin;
1	Utyug	1199.99	12
2	Holodilnik	5000	20
3	Chainik	1500.1	15
4	Noutbuk	9999.99	50
5	Myshka	200	200
6	Klaviatura	100	100
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: imarket_schema; Owner: postgres
--

COPY imarket_schema.orders (id, product_id, quantity, status) FROM stdin;
1	1	1	Registered
2	2	2	Paid
3	3	3	Cancelled
4	4	4	Delivered
5	1	5	Delivered
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, product_id, quantity, status) FROM stdin;
2	10	100000	shipped
1	1	1	pending
4	4	4	cancelled
5	1	5	delivered
3	3	3	cancelled
6	1	10	shipped
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, name, price, quantity) FROM stdin;
3	Chainik	1500.1	15
4	Noutbuk	9999.99	50
5	Myshka	200	200
6	Klaviatura	100	100
2	Sosiska	3	1000000
1	Utyug	1199.99	2
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, password) FROM stdin;
\.


--
-- PostgreSQL database dump complete
--

