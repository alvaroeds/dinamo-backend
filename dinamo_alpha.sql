-- =============================================================================
-- Diagram Name: DNAMO_MR
-- Created on: 16/11/2019 19:42:06
-- Diagram Version: 
-- =============================================================================


DROP TABLE IF EXISTS "user_dinamo" CASCADE;
CREATE TABLE "user_dinamo" (
	"id" SERIAL NOT NULL,
	"email" varchar(40) NOT NULL,
	"nickname" varchar(20),
	"slug" varchar(100),
	"phone_number" varchar(20),
	"name" varchar(40),
	"lastname" varchar(40),
	"photo" varchar(100),
	"password" varchar,
	"created_at" timestamp with time zone,
	"updated_at" timestamp with time zone,
	"deleted_at" timestamp with time zone,
	PRIMARY KEY("id"),
	CONSTRAINT "user_email_unique" UNIQUE("email"),
	CONSTRAINT "user_nickname_unique" UNIQUE("nickname"),
	CONSTRAINT "user_slug_unique" UNIQUE("slug")
);

DROP TABLE IF EXISTS "folder" CASCADE;
CREATE TABLE "folder" (
	"id" SERIAL NOT NULL,
	"id_user" int4 NOT NULL,
	"id_room" int4,
	"name" varchar(50) NOT NULL,
	"teacher" varchar(30),
	"lastuse" date,
	"flag" bool NOT NULL,
	"created_at" timestamp with time zone,
	"updated_at" timestamp with time zone,
	"deleted_at" timestamp with time zone,
	PRIMARY KEY("id")
);

DROP TABLE IF EXISTS "title" CASCADE;
CREATE TABLE "title" (
	"id" SERIAL NOT NULL,
	"id_folder" int4 NOT NULL,
	"date_client" date NOT NULL,
	"timeline" bool NOT NULL,
	"name" varchar(50) NOT NULL,
	"lastuse" date,
	"created_at" timestamp with time zone,
	"updated_at" timestamp with time zone,
	"deleted_at" timestamp with time zone,
	PRIMARY KEY("id")
);

DROP TABLE IF EXISTS "image" CASCADE;
CREATE TABLE "image" (
	"id" SERIAL NOT NULL,
	"id_title" int4 NOT NULL,
	"route" varchar(300) NOT NULL,
	"note" varchar(300),
	"position" int4,
	"created_at" timestamp with time zone NOT NULL,
	"updated_at" timestamp with time zone,
	"deleted_at" timestamp with time zone,
	PRIMARY KEY("id")
);

DROP TABLE IF EXISTS "day" CASCADE;
CREATE TABLE "day" (
	"id" SERIAL NOT NULL,
	"name" varchar(20) NOT NULL,
	PRIMARY KEY("id")
);

DROP TABLE IF EXISTS "calendar" CASCADE;
CREATE TABLE "calendar" (
	"id" int4 NOT NULL,
	"id_folder" int4 NOT NULL,
	"id_day" int4 NOT NULL,
	"classroom" varchar(20),
	"start_time" time NOT NULL,
	"end_time" time NOT NULL,
	"flag" bool NOT NULL,
	PRIMARY KEY("id")
);

DROP TABLE IF EXISTS "room" CASCADE;
CREATE TABLE "room" (
	"id" SERIAL NOT NULL,
	"cod_room" varchar(15),
	"name" varchar(50),
	"id_inbox" int4,
	"created_at" timestamp with time zone,
	"updated_at" timestamp with time zone,
	"deleted_at" timestamp with time zone,
	PRIMARY KEY("id")
);

DROP TABLE IF EXISTS "contact" CASCADE;
CREATE TABLE "contact" (
	"id" SERIAL NOT NULL,
	"id_user" int4,
	"id_contact" int4,
	"id_inbox" int4,
	PRIMARY KEY("id")
);

DROP TABLE IF EXISTS "message" CASCADE;
CREATE TABLE "message" (
	"id" SERIAL NOT NULL,
	"id_inbox" int4,
	"contend" text,
	"created_at" timestamp with time zone,
	"updated_at" timestamp with time zone,
	"deleted_at" timestamp with time zone,
	PRIMARY KEY("id")
);

DROP TABLE IF EXISTS "inbox" CASCADE;
CREATE TABLE "inbox" (
	"id" SERIAL NOT NULL,
	"id_image" int4,
	PRIMARY KEY("id")
);

DROP TABLE IF EXISTS "menber" CASCADE;
CREATE TABLE "menber" (
	"id_room" int4,
	"id_user" int4 NOT NULL
);


ALTER TABLE "folder" ADD CONSTRAINT "fk_user_folder" FOREIGN KEY ("id_user")
	REFERENCES "user_dinamo"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "folder" ADD CONSTRAINT "Ref_carpeta_to_sala" FOREIGN KEY ("id_room")
	REFERENCES "room"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "title" ADD CONSTRAINT "fk_folder_title" FOREIGN KEY ("id_folder")
	REFERENCES "folder"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "image" ADD CONSTRAINT "fk_title_image" FOREIGN KEY ("id_title")
	REFERENCES "title"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "calendar" ADD CONSTRAINT "fk_day_calendar" FOREIGN KEY ("id_day")
	REFERENCES "day"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "calendar" ADD CONSTRAINT "fk_folder_calendar" FOREIGN KEY ("id_folder")
	REFERENCES "folder"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "room" ADD CONSTRAINT "Ref_room_to_inbox" FOREIGN KEY ("id_inbox")
	REFERENCES "inbox"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "contact" ADD CONSTRAINT "fk_contact_user" FOREIGN KEY ("id_user")
	REFERENCES "user_dinamo"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "contact" ADD CONSTRAINT "Ref_contactos_to_inbox" FOREIGN KEY ("id_inbox")
	REFERENCES "inbox"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "contact" ADD CONSTRAINT "fk_contact_user_contact" FOREIGN KEY ("id_contact")
	REFERENCES "user_dinamo"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "message" ADD CONSTRAINT "Ref_mensaje_to_inbox" FOREIGN KEY ("id_inbox")
	REFERENCES "inbox"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "inbox" ADD CONSTRAINT "Ref_inbox_to_image" FOREIGN KEY ("id_image")
	REFERENCES "image"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "menber" ADD CONSTRAINT "Ref_menber_to_user" FOREIGN KEY ("id_user")
	REFERENCES "user_dinamo"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

ALTER TABLE "menber" ADD CONSTRAINT "Ref_menber_to_room" FOREIGN KEY ("id_room")
	REFERENCES "room"("id")
	MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	NOT DEFERRABLE;

INSERT INTO user_dinamo (email, name, lastname, password, phone_number) VALUES('prueba@dinamo.com', 'Alvaro', 'Diaz',
'655e786674d9d3e77bc05ed1de37b4b6bc89f788829f9f3c679e7687b410c89b',
'987654321');


INSERT INTO day (name) VALUES('lunes');
INSERT INTO day (name) VALUES('martes');
INSERT INTO day (name) VALUES('miercoles');
INSERT INTO day (name) VALUES('jueves');
INSERT INTO day (name) VALUES('viernes');
INSERT INTO day (name) VALUES('sabado');
INSERT INTO day (name) VALUES('domingo');
