--
-- PostgreSQL database dump
--

-- Dumped from database version 13.6 (Ubuntu 13.6-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.6

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
-- Name: event_participants; Type: TABLE; Schema: public; Owner: slvjtpvoyxjyff
--

CREATE TABLE public.event_participants (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    event_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.event_participants OWNER TO slvjtpvoyxjyff;

--
-- Name: events; Type: TABLE; Schema: public; Owner: slvjtpvoyxjyff
--

CREATE TABLE public.events (
    id uuid NOT NULL,
    title character varying(255) DEFAULT ''::character varying NOT NULL,
    description character varying(255) DEFAULT ''::character varying NOT NULL,
    gift_id uuid NOT NULL,
    user_id uuid NOT NULL,
    date date NOT NULL,
    sent boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.events OWNER TO slvjtpvoyxjyff;

--
-- Name: gifts; Type: TABLE; Schema: public; Owner: slvjtpvoyxjyff
--

CREATE TABLE public.gifts (
    id uuid NOT NULL,
    code uuid NOT NULL,
    title character varying(255) DEFAULT ''::character varying NOT NULL,
    video character varying(255) NOT NULL,
    video_url text NOT NULL,
    qr text NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.gifts OWNER TO slvjtpvoyxjyff;

--
-- Name: participants; Type: TABLE; Schema: public; Owner: slvjtpvoyxjyff
--

CREATE TABLE public.participants (
    id uuid NOT NULL,
    first_name character varying(255) DEFAULT ''::character varying NOT NULL,
    last_name character varying(255) NOT NULL,
    correo character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.participants OWNER TO slvjtpvoyxjyff;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: slvjtpvoyxjyff
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO slvjtpvoyxjyff;

--
-- Name: users; Type: TABLE; Schema: public; Owner: slvjtpvoyxjyff
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    first_name character varying(255) DEFAULT ''::character varying NOT NULL,
    last_name character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    phone_number character varying(255) NOT NULL,
    phone_extension character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO slvjtpvoyxjyff;

--
-- Name: event_participants event_participants_pkey; Type: CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.event_participants
    ADD CONSTRAINT event_participants_pkey PRIMARY KEY (id);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: gifts gifts_pkey; Type: CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.gifts
    ADD CONSTRAINT gifts_pkey PRIMARY KEY (id);


--
-- Name: participants participants_pkey; Type: CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.participants
    ADD CONSTRAINT participants_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: slvjtpvoyxjyff
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: event_participants event_participants_id_fk_event_id; Type: FK CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.event_participants
    ADD CONSTRAINT event_participants_id_fk_event_id FOREIGN KEY (event_id) REFERENCES public.events(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_participants event_participants_id_fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.event_participants
    ADD CONSTRAINT event_participants_id_fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: events events_id_fk_gift_id; Type: FK CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_id_fk_gift_id FOREIGN KEY (gift_id) REFERENCES public.gifts(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: events events_id_fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_id_fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: gifts gifts_id_fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: slvjtpvoyxjyff
--

ALTER TABLE ONLY public.gifts
    ADD CONSTRAINT gifts_id_fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: slvjtpvoyxjyff
--

REVOKE ALL ON SCHEMA public FROM postgres;
REVOKE ALL ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO slvjtpvoyxjyff;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- Name: LANGUAGE plpgsql; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON LANGUAGE plpgsql TO slvjtpvoyxjyff;


--
-- PostgreSQL database dump complete
--

