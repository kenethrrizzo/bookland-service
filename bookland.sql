drop database if exists bookland;

create database bookland;
use bookland;

-- Schemas creation
create table User
(
    Id          int auto_increment primary key,
    Name        varchar(20)  not null,
    Surname     varchar(20)  not null,
    Email       varchar(25)  not null unique,
    Username    varchar(20)  not null,
    Password    varchar(100) not null,
    DateOfBirth date,
    CreatedAt   datetime default now()
);

create table Genre
(
    Code        varchar(3) primary key,
    Description varchar(20)
);

create table Book
(
    Id        int auto_increment primary key,
    Name      varchar(100)   not null,
    Author    varchar(50)    not null,
    CoverPage varchar(200),
    Synopsis  text,
    Price     decimal(15, 2) not null,
    CreatedAt datetime default now(),
    UpdatedAt datetime,
    Status    char     default 'A' -- A: Active; I: Inactive
);


create table BookGenre
(
    BookId    int,
    GenreCode varchar(3)
);

alter table BookGenre
    add foreign key (BookId) references Book (Id);

alter table BookGenre
    add foreign key (GenreCode) references Genre (Code);

-- Dummy data
insert into Genre (Code, Description)
values ('TER', 'TERROR');
insert into Genre (Code, Description)
values ('COM', 'COMEDIA');
insert into Genre (Code, Description)
values ('MIS', 'MISTERIO');
insert into Genre (Code, Description)
values ('POL', 'POLICIAL');
insert into Genre (Code, Description)
values ('DRA', 'DRAMA');

insert into Book (Id, Name, Author, CoverPage, Synopsis, Price)
values (1, 'La ladrona de libros', 'Markus Zusak',
        'https://imusic.b-cdn.net/images/item/original/075/9788499088075.jpg',
        'En la Alemania previa a la II Guerra Mundial, Liesel y su hermano son enviados a vivir con una familia de acogida, pero el niño fallece. Para ella, el poder de las palabras y la imaginación se convierte en una forma de escapar de lo que le rodea.',
        16.99);

insert into Book (Id, Name, Author, CoverPage, Synopsis, Price)
values (2, 'IT', 'Stephen King',
        'https://images.cdn1.buscalibre.com/fit-in/360x360/ec/c6/ecc6925af7478dd66fce402ea5e3dda0.jpg',
        'Varios niños de una pequeña ciudad del estado de Maine se alían para combatir a una entidad diabólica que adopta la forma de un payaso y desde hace mucho tiempo emerge cada 27 años para saciarse de sangre infantil.',
        26.99);

insert into Book (Id, Name, Author, CoverPage, Synopsis, Price)
values (3, 'La llamada de Cthulu', 'H.P Lovecraft',
        'https://planetadelibrosec2.cdnstatics.com/usuaris/libros/fotos/315/original/314487_portada_la-llamada-de-cthulhu_alvaro-robledo_201908232355.jpg',
        'Hay un culto oscuro más antiguo que la humanidad y que la Tierra. Sus integrantes esperan que las estrellas estén alineadas para que el Gran Cthulhu regrese desde las profundidades del océano y reclame su reino sobre el universo.',
        35.50);

insert into BookGenre (BookId, GenreCode)
values (1, 'DRA');
insert into BookGenre (BookId, GenreCode)
values (2, 'TER');
insert into BookGenre (BookId, GenreCode)
values (3, 'TER');
insert into BookGenre (BookId, GenreCode)
values (1, 'MIS');
insert into BookGenre (BookId, GenreCode)
values (2, 'MIS');
insert into BookGenre (BookId, GenreCode)
values (3, 'MIS');