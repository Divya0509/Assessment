PGDMP  6    3                |            movie_db    15.7    16.3     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                        0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    24576    movie_db    DATABASE     {   CREATE DATABASE movie_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_India.1252';
    DROP DATABASE movie_db;
                postgres    false                        2615    24577    sample    SCHEMA        CREATE SCHEMA sample;
    DROP SCHEMA sample;
                postgres    false            �            1259    24597    movies    TABLE     �  CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    year integer,
    ratings integer,
    genre character varying(100),
    CONSTRAINT movies_ratings_check CHECK (((ratings >= 0) AND (ratings <= 5))),
    CONSTRAINT movies_year_check CHECK (((year >= 1900) AND ((year)::numeric <= EXTRACT(year FROM CURRENT_DATE))))
);
    DROP TABLE public.movies;
       public         heap    postgres    false            �            1259    24596    movies_id_seq    SEQUENCE     �   CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.movies_id_seq;
       public          postgres    false    216                       0    0    movies_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;
          public          postgres    false    215            f           2604    24600 	   movies id    DEFAULT     f   ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);
 8   ALTER TABLE public.movies ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            �          0    24597    movies 
   TABLE DATA           N   COPY public.movies (id, title, description, year, ratings, genre) FROM stdin;
    public          postgres    false    216   G                  0    0    movies_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.movies_id_seq', 9, true);
          public          postgres    false    215            j           2606    24606    movies movies_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.movies DROP CONSTRAINT movies_pkey;
       public            postgres    false    216            l           2606    24608    movies movies_title_key 
   CONSTRAINT     S   ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_title_key UNIQUE (title);
 A   ALTER TABLE ONLY public.movies DROP CONSTRAINT movies_title_key;
       public            postgres    false    216            �   �  x�]RKn�0]ӧ��B�	��(��t3�F�))>W�Ћ�Qjڠ���h�>3��}�m
Bd�F�<�ih��zeT�T1�l�<��|�g�*�]���;�w%�:�B��G�l�i.�Y���De�̓�Q�?�LM1�t"|	E�;����F���}��Q��4~<�S;���h�P�z���̹�4��2�E�-]LK�����d��}�r�jA:E=]����b�*�{�WϽ{k|�ru_�:�^}�{�!�O[���j�����O��0ހ
�/��l�IX=�C[ �`�������{�O�ٵ�R�X2�lJ�E(�7)�i�����=2-W�ܷ�"�,�X��i�L�1ej� A2$҄�h.��_Lr�v�Lt�tFNP���i�KF�|��K�,P�l�;#0��{���f��Ѷ��     