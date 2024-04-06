PGDMP      -                |            ticketing_db    16.2    16.2 9    I           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            J           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            K           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            L           1262    16721    ticketing_db    DATABASE     n   CREATE DATABASE ticketing_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';
    DROP DATABASE ticketing_db;
                postgres    false            �            1259    16746    cinema_screens    TABLE       CREATE TABLE public.cinema_screens (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    cinema_id bigint NOT NULL
);
 "   DROP TABLE public.cinema_screens;
       public         heap    postgres    false            �            1259    16745    cinema_screens_id_seq    SEQUENCE     ~   CREATE SEQUENCE public.cinema_screens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.cinema_screens_id_seq;
       public          postgres    false    220            M           0    0    cinema_screens_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.cinema_screens_id_seq OWNED BY public.cinema_screens.id;
          public          postgres    false    219            �            1259    16733    cinemas    TABLE     �   CREATE TABLE public.cinemas (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    city_id bigint NOT NULL
);
    DROP TABLE public.cinemas;
       public         heap    postgres    false            �            1259    16732    cinemas_id_seq    SEQUENCE     w   CREATE SEQUENCE public.cinemas_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.cinemas_id_seq;
       public          postgres    false    218            N           0    0    cinemas_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.cinemas_id_seq OWNED BY public.cinemas.id;
          public          postgres    false    217            �            1259    16723    cities    TABLE     	  CREATE TABLE public.cities (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    zipcode character varying(255) NOT NULL
);
    DROP TABLE public.cities;
       public         heap    postgres    false            �            1259    16722    cities_id_seq    SEQUENCE     v   CREATE SEQUENCE public.cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.cities_id_seq;
       public          postgres    false    216            O           0    0    cities_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;
          public          postgres    false    215            �            1259    16781    movie_shows    TABLE     ]  CREATE TABLE public.movie_shows (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    cinema_screen_id bigint NOT NULL,
    movie_id bigint NOT NULL
);
    DROP TABLE public.movie_shows;
       public         heap    postgres    false            �            1259    16780    movie_shows_id_seq    SEQUENCE     {   CREATE SEQUENCE public.movie_shows_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.movie_shows_id_seq;
       public          postgres    false    226            P           0    0    movie_shows_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.movie_shows_id_seq OWNED BY public.movie_shows.id;
          public          postgres    false    225            �            1259    16759    movies    TABLE     ,  CREATE TABLE public.movies (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    duration bigint NOT NULL
);
    DROP TABLE public.movies;
       public         heap    postgres    false            �            1259    16758    movies_id_seq    SEQUENCE     v   CREATE SEQUENCE public.movies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.movies_id_seq;
       public          postgres    false    222            Q           0    0    movies_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;
          public          postgres    false    221            �            1259    16769    users    TABLE     4  CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    16768    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    224            R           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    223            �           2604    16749    cinema_screens id    DEFAULT     v   ALTER TABLE ONLY public.cinema_screens ALTER COLUMN id SET DEFAULT nextval('public.cinema_screens_id_seq'::regclass);
 @   ALTER TABLE public.cinema_screens ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    219    220    220            �           2604    16736 
   cinemas id    DEFAULT     h   ALTER TABLE ONLY public.cinemas ALTER COLUMN id SET DEFAULT nextval('public.cinemas_id_seq'::regclass);
 9   ALTER TABLE public.cinemas ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    217    218            �           2604    16726 	   cities id    DEFAULT     f   ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);
 8   ALTER TABLE public.cities ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            �           2604    16784    movie_shows id    DEFAULT     p   ALTER TABLE ONLY public.movie_shows ALTER COLUMN id SET DEFAULT nextval('public.movie_shows_id_seq'::regclass);
 =   ALTER TABLE public.movie_shows ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    225    226    226            �           2604    16762 	   movies id    DEFAULT     f   ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);
 8   ALTER TABLE public.movies ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    221    222    222            �           2604    16772    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    224    223    224            @          0    16746    cinema_screens 
   TABLE DATA           a   COPY public.cinema_screens (id, created_at, updated_at, deleted_at, name, cinema_id) FROM stdin;
    public          postgres    false    220   <A       >          0    16733    cinemas 
   TABLE DATA           X   COPY public.cinemas (id, created_at, updated_at, deleted_at, name, city_id) FROM stdin;
    public          postgres    false    218   YA       <          0    16723    cities 
   TABLE DATA           W   COPY public.cities (id, created_at, updated_at, deleted_at, name, zipcode) FROM stdin;
    public          postgres    false    216   vA       F          0    16781    movie_shows 
   TABLE DATA              COPY public.movie_shows (id, created_at, updated_at, deleted_at, start_time, end_time, cinema_screen_id, movie_id) FROM stdin;
    public          postgres    false    226   �A       B          0    16759    movies 
   TABLE DATA           f   COPY public.movies (id, created_at, updated_at, deleted_at, title, description, duration) FROM stdin;
    public          postgres    false    222   �A       D          0    16769    users 
   TABLE DATA           ^   COPY public.users (id, created_at, updated_at, deleted_at, name, email, password) FROM stdin;
    public          postgres    false    224   �A       S           0    0    cinema_screens_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.cinema_screens_id_seq', 1, false);
          public          postgres    false    219            T           0    0    cinemas_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.cinemas_id_seq', 1, false);
          public          postgres    false    217            U           0    0    cities_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.cities_id_seq', 1, false);
          public          postgres    false    215            V           0    0    movie_shows_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.movie_shows_id_seq', 1, false);
          public          postgres    false    225            W           0    0    movies_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.movies_id_seq', 1, false);
          public          postgres    false    221            X           0    0    users_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.users_id_seq', 1, false);
          public          postgres    false    223            �           2606    16751 "   cinema_screens cinema_screens_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.cinema_screens
    ADD CONSTRAINT cinema_screens_pkey PRIMARY KEY (id);
 L   ALTER TABLE ONLY public.cinema_screens DROP CONSTRAINT cinema_screens_pkey;
       public            postgres    false    220            �           2606    16738    cinemas cinemas_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT cinemas_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.cinemas DROP CONSTRAINT cinemas_pkey;
       public            postgres    false    218            �           2606    16730    cities cities_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.cities DROP CONSTRAINT cities_pkey;
       public            postgres    false    216            �           2606    16786    movie_shows movie_shows_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.movie_shows
    ADD CONSTRAINT movie_shows_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.movie_shows DROP CONSTRAINT movie_shows_pkey;
       public            postgres    false    226            �           2606    16766    movies movies_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.movies DROP CONSTRAINT movies_pkey;
       public            postgres    false    222            �           2606    16778    users uni_users_email 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT uni_users_email;
       public            postgres    false    224            �           2606    16776    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    224            �           1259    16757    idx_cinema_screens_deleted_at    INDEX     ^   CREATE INDEX idx_cinema_screens_deleted_at ON public.cinema_screens USING btree (deleted_at);
 1   DROP INDEX public.idx_cinema_screens_deleted_at;
       public            postgres    false    220            �           1259    16744    idx_cinemas_deleted_at    INDEX     P   CREATE INDEX idx_cinemas_deleted_at ON public.cinemas USING btree (deleted_at);
 *   DROP INDEX public.idx_cinemas_deleted_at;
       public            postgres    false    218            �           1259    16731    idx_cities_deleted_at    INDEX     N   CREATE INDEX idx_cities_deleted_at ON public.cities USING btree (deleted_at);
 )   DROP INDEX public.idx_cities_deleted_at;
       public            postgres    false    216            �           1259    16797    idx_movie_shows_deleted_at    INDEX     X   CREATE INDEX idx_movie_shows_deleted_at ON public.movie_shows USING btree (deleted_at);
 .   DROP INDEX public.idx_movie_shows_deleted_at;
       public            postgres    false    226            �           1259    16767    idx_movies_deleted_at    INDEX     N   CREATE INDEX idx_movies_deleted_at ON public.movies USING btree (deleted_at);
 )   DROP INDEX public.idx_movies_deleted_at;
       public            postgres    false    222            �           1259    16779    idx_users_deleted_at    INDEX     L   CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);
 (   DROP INDEX public.idx_users_deleted_at;
       public            postgres    false    224            �           2606    16752 '   cinema_screens fk_cinema_screens_cinema    FK CONSTRAINT     �   ALTER TABLE ONLY public.cinema_screens
    ADD CONSTRAINT fk_cinema_screens_cinema FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id);
 Q   ALTER TABLE ONLY public.cinema_screens DROP CONSTRAINT fk_cinema_screens_cinema;
       public          postgres    false    3480    220    218            �           2606    16739    cinemas fk_cinemas_city    FK CONSTRAINT     w   ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT fk_cinemas_city FOREIGN KEY (city_id) REFERENCES public.cities(id);
 A   ALTER TABLE ONLY public.cinemas DROP CONSTRAINT fk_cinemas_city;
       public          postgres    false    3477    216    218            �           2606    16792 (   movie_shows fk_movie_shows_cinema_screen    FK CONSTRAINT     �   ALTER TABLE ONLY public.movie_shows
    ADD CONSTRAINT fk_movie_shows_cinema_screen FOREIGN KEY (cinema_screen_id) REFERENCES public.cinema_screens(id);
 R   ALTER TABLE ONLY public.movie_shows DROP CONSTRAINT fk_movie_shows_cinema_screen;
       public          postgres    false    220    226    3483            �           2606    16787     movie_shows fk_movie_shows_movie    FK CONSTRAINT     �   ALTER TABLE ONLY public.movie_shows
    ADD CONSTRAINT fk_movie_shows_movie FOREIGN KEY (movie_id) REFERENCES public.movies(id);
 J   ALTER TABLE ONLY public.movie_shows DROP CONSTRAINT fk_movie_shows_movie;
       public          postgres    false    3487    226    222            @      x������ � �      >      x������ � �      <      x������ � �      F      x������ � �      B      x������ � �      D      x������ � �     