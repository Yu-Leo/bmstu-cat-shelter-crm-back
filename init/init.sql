DROP TABLE IF EXISTS cats;
CREATE TABLE cats
(
    photo_url                    text,
    nickname                     text        NOT NULL,
    gender                       text        NOT NULL CHECK (gender IN ('male', 'female')),
    age                          integer     NOT NULL CHECK (age >= 0),
    chip_number                  text UNIQUE NOT NULL CHECK (LENGTH(chip_number) == 15) PRIMARY KEY,
    date_of_admission_to_shelter date        NOT NULL
);

DROP TABLE IF EXISTS residents;
CREATE TABLE residents
(
    cat_chip_number text UNIQUE NOT NULL PRIMARY KEY,
    booking         bool,
    aggressiveness  bool,
    vk_album_url    text,
    guardian_id     integer,
    room_number     text,
    FOREIGN KEY (cat_chip_number) REFERENCES cats (chip_number),
    FOREIGN KEY (guardian_id) REFERENCES guardians (guardian_id),
    FOREIGN KEY (room_number) REFERENCES rooms (number)
);

DROP TABLE IF EXISTS people;
CREATE TABLE people
(
    person_id  integer UNIQUE NOT NULL PRIMARY KEY,
    photo_url  text,
    firstname  text           NOT NULL,
    lastname   text           NOT NULL,
    patronymic text,
    phone      text UNIQUE    NOT NULL CHECK (LENGTH(phone) == 11)
);

DROP TABLE IF EXISTS guardians;
CREATE TABLE guardians
(
    guardian_id integer UNIQUE NOT NULL PRIMARY KEY,
    person_id   integer        NOT NULL,
    FOREIGN KEY (person_id) REFERENCES people (person_id)
);

DROP TABLE IF EXISTS rooms;
CREATE TABLE rooms
(
    number text UNIQUE NOT NULL CHECK (LENGTH(number) == 2) PRIMARY KEY,
    status text
);