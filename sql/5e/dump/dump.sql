--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

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
-- Name: ability_scores; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ability_scores (
    ability_index character varying NOT NULL,
    ability_name character varying,
    full_name character varying,
    description text[],
    url character varying
);


ALTER TABLE public.ability_scores OWNER TO postgres;

--
-- Name: alignments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alignments (
    alignment_index character varying NOT NULL,
    alignment_name character varying,
    abbreviation character varying,
    description text,
    url character varying
);


ALTER TABLE public.alignments OWNER TO postgres;

--
-- Name: classes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.classes (
    index character varying NOT NULL,
    name character varying,
    hit_die integer,
    class_levels character varying,
    multi_classing jsonb,
    subclasses jsonb,
    spellcasting jsonb,
    spells character varying,
    url character varying
);


ALTER TABLE public.classes OWNER TO postgres;

--
-- Name: condition_descriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.condition_descriptions (
    condition_index character varying,
    description text
);


ALTER TABLE public.condition_descriptions OWNER TO postgres;

--
-- Name: conditions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.conditions (
    index character varying NOT NULL,
    name character varying,
    url character varying
);


ALTER TABLE public.conditions OWNER TO postgres;

--
-- Name: damage_type_descriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.damage_type_descriptions (
    damage_type_index character varying,
    description text
);


ALTER TABLE public.damage_type_descriptions OWNER TO postgres;

--
-- Name: damage_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.damage_types (
    index character varying NOT NULL,
    name character varying,
    url character varying
);


ALTER TABLE public.damage_types OWNER TO postgres;

--
-- Name: proficiencies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.proficiencies (
    class_index character varying,
    proficiency_index character varying,
    proficiency_name character varying,
    proficiency_url character varying
);


ALTER TABLE public.proficiencies OWNER TO postgres;

--
-- Name: proficiency_choices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.proficiency_choices (
    class_index character varying,
    description character varying,
    choose_number integer,
    proficiency_type character varying,
    proficiency_options jsonb
);


ALTER TABLE public.proficiency_choices OWNER TO postgres;

--
-- Name: saving_throws; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.saving_throws (
    class_index character varying,
    saving_throw_index character varying,
    saving_throw_name character varying,
    saving_throw_url character varying
);


ALTER TABLE public.saving_throws OWNER TO postgres;

--
-- Name: skills; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.skills (
    skill_index character varying NOT NULL,
    skill_name character varying,
    ability_index character varying,
    url character varying
);


ALTER TABLE public.skills OWNER TO postgres;

--
-- Name: starting_equipment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.starting_equipment (
    class_index character varying,
    equipment_index character varying,
    equipment_name character varying,
    equipment_url character varying,
    quantity integer
);


ALTER TABLE public.starting_equipment OWNER TO postgres;

--
-- Name: starting_equipment_options; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.starting_equipment_options (
    class_index character varying,
    description text,
    choose integer,
    equipment_type character varying,
    equipment_options jsonb
);


ALTER TABLE public.starting_equipment_options OWNER TO postgres;

--
-- Name: ability_scores ability_scores_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ability_scores
    ADD CONSTRAINT ability_scores_pkey PRIMARY KEY (ability_index);


--
-- Name: alignments alignments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alignments
    ADD CONSTRAINT alignments_pkey PRIMARY KEY (alignment_index);


--
-- Name: classes classes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.classes
    ADD CONSTRAINT classes_pkey PRIMARY KEY (index);


--
-- Name: conditions conditions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.conditions
    ADD CONSTRAINT conditions_pkey PRIMARY KEY (index);


--
-- Name: damage_types damage_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.damage_types
    ADD CONSTRAINT damage_types_pkey PRIMARY KEY (index);


--
-- Name: skills skills_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.skills
    ADD CONSTRAINT skills_pkey PRIMARY KEY (skill_index);


--
-- Name: condition_descriptions condition_descriptions_condition_index_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.condition_descriptions
    ADD CONSTRAINT condition_descriptions_condition_index_fkey FOREIGN KEY (condition_index) REFERENCES public.conditions(index);


--
-- Name: damage_type_descriptions damage_type_descriptions_damage_type_index_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.damage_type_descriptions
    ADD CONSTRAINT damage_type_descriptions_damage_type_index_fkey FOREIGN KEY (damage_type_index) REFERENCES public.damage_types(index);


--
-- Name: proficiencies proficiencies_class_index_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.proficiencies
    ADD CONSTRAINT proficiencies_class_index_fkey FOREIGN KEY (class_index) REFERENCES public.classes(index);


--
-- Name: proficiency_choices proficiency_choices_class_index_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.proficiency_choices
    ADD CONSTRAINT proficiency_choices_class_index_fkey FOREIGN KEY (class_index) REFERENCES public.classes(index);


--
-- Name: saving_throws saving_throws_class_index_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.saving_throws
    ADD CONSTRAINT saving_throws_class_index_fkey FOREIGN KEY (class_index) REFERENCES public.classes(index);


--
-- Name: skills skills_ability_index_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.skills
    ADD CONSTRAINT skills_ability_index_fkey FOREIGN KEY (ability_index) REFERENCES public.ability_scores(ability_index);


--
-- Name: starting_equipment starting_equipment_class_index_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starting_equipment
    ADD CONSTRAINT starting_equipment_class_index_fkey FOREIGN KEY (class_index) REFERENCES public.classes(index);


--
-- Name: starting_equipment_options starting_equipment_options_class_index_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starting_equipment_options
    ADD CONSTRAINT starting_equipment_options_class_index_fkey FOREIGN KEY (class_index) REFERENCES public.classes(index);


--
-- PostgreSQL database dump complete
--

