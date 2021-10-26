--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Debian 13.3-1.pgdg100+1)
-- Dumped by pg_dump version 13.4 (Ubuntu 13.4-4.pgdg20.04+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: daily_tradings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.daily_tradings (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    symbol text,
    date timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
    trade_volume bigint DEFAULT 0,
    trade_value bigint DEFAULT 0,
    opening_price numeric DEFAULT 0.000000,
    highest_price numeric DEFAULT 0.000000,
    lowest_price numeric DEFAULT 0.000000,
    closing_price numeric DEFAULT 0.000000,
    change numeric DEFAULT 0.000000,
    transaction bigint DEFAULT 0,
    dividend_yield numeric DEFAULT 0.000000,
    dividend_year text,
    pe_radio numeric DEFAULT 0.000000,
    pb_radio numeric DEFAULT 0.000000,
    fiscal_year_quarter text
);


ALTER TABLE public.daily_tradings OWNER TO postgres;

--
-- Name: daily_tradings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.daily_tradings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.daily_tradings_id_seq OWNER TO postgres;

--
-- Name: daily_tradings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.daily_tradings_id_seq OWNED BY public.daily_tradings.id;


--
-- Name: highlights; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.highlights (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    date timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
    trade_volume numeric DEFAULT 0.000000,
    trade_value numeric DEFAULT 0.000000,
    transaction numeric DEFAULT 0.000000,
    taiex numeric DEFAULT 0.000000,
    change numeric DEFAULT 0.000000
);


ALTER TABLE public.highlights OWNER TO postgres;

--
-- Name: highlights_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.highlights_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.highlights_id_seq OWNER TO postgres;

--
-- Name: highlights_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.highlights_id_seq OWNED BY public.highlights.id;


--
-- Name: stock_infos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stock_infos (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    symbol text,
    name text,
    market_type text,
    industry text,
    listing_date timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone
);


ALTER TABLE public.stock_infos OWNER TO postgres;

--
-- Name: stock_infos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.stock_infos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.stock_infos_id_seq OWNER TO postgres;

--
-- Name: stock_infos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.stock_infos_id_seq OWNED BY public.stock_infos.id;


--
-- Name: taiexes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.taiexes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    date timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
    opening_index numeric(10,2) DEFAULT 0.000000,
    closing_index numeric(10,2) DEFAULT 0.000000,
    lowest_index numeric(10,2) DEFAULT 0.000000,
    highest_index numeric(10,2) DEFAULT 0.000000
);


ALTER TABLE public.taiexes OWNER TO postgres;

--
-- Name: taiexes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.taiexes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.taiexes_id_seq OWNER TO postgres;

--
-- Name: taiexes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.taiexes_id_seq OWNED BY public.taiexes.id;


--
-- Name: daily_tradings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.daily_tradings ALTER COLUMN id SET DEFAULT nextval('public.daily_tradings_id_seq'::regclass);


--
-- Name: highlights id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.highlights ALTER COLUMN id SET DEFAULT nextval('public.highlights_id_seq'::regclass);


--
-- Name: stock_infos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_infos ALTER COLUMN id SET DEFAULT nextval('public.stock_infos_id_seq'::regclass);


--
-- Name: taiexes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.taiexes ALTER COLUMN id SET DEFAULT nextval('public.taiexes_id_seq'::regclass);


--
-- Data for Name: daily_tradings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.daily_tradings (id, created_at, updated_at, deleted_at, symbol, date, trade_volume, trade_value, opening_price, highest_price, lowest_price, closing_price, change, transaction, dividend_yield, dividend_year, pe_radio, pb_radio, fiscal_year_quarter) FROM stdin;
\.


--
-- Data for Name: highlights; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.highlights (id, created_at, updated_at, deleted_at, date, trade_volume, trade_value, transaction, taiex, change) FROM stdin;
\.


--
-- Data for Name: stock_infos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.stock_infos (id, created_at, updated_at, deleted_at, symbol, name, market_type, industry, listing_date) FROM stdin;
\.


--
-- Data for Name: taiexes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.taiexes (id, created_at, updated_at, deleted_at, date, opening_index, closing_index, lowest_index, highest_index) FROM stdin;
\.


--
-- Name: daily_tradings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.daily_tradings_id_seq', 110604, true);


--
-- Name: highlights_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.highlights_id_seq', 3388, true);


--
-- Name: stock_infos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.stock_infos_id_seq', 956, true);


--
-- Name: taiexes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.taiexes_id_seq', 2210, true);


--
-- Name: daily_tradings daily_tradings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.daily_tradings
    ADD CONSTRAINT daily_tradings_pkey PRIMARY KEY (id);


--
-- Name: highlights highlights_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.highlights
    ADD CONSTRAINT highlights_pkey PRIMARY KEY (id);


--
-- Name: stock_infos stock_infos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_infos
    ADD CONSTRAINT stock_infos_pkey PRIMARY KEY (id);


--
-- Name: taiexes taiexes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.taiexes
    ADD CONSTRAINT taiexes_pkey PRIMARY KEY (id);


--
-- Name: idx_daily_tradings_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_daily_tradings_deleted_at ON public.daily_tradings USING btree (deleted_at);


--
-- Name: idx_highlights_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_highlights_deleted_at ON public.highlights USING btree (deleted_at);


--
-- Name: idx_stock_infos_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_stock_infos_deleted_at ON public.stock_infos USING btree (deleted_at);


--
-- Name: idx_taiexes_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_taiexes_deleted_at ON public.taiexes USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

