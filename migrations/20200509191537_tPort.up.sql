create table dbo.Port
(
    Id          tinyint identity not null,
    ServiceName varchar(100)     not null,
    GrpcPort    int              not null,
    AccountId   int,
    constraint PK_Port primary key (Id),
    constraint UI_GrpcPort unique (GrpcPort),
)

insert into dbo.Port (ServiceName, GrpcPort)
values ('forted-service', 50051),
       ('surebet-service', 50052),
       ('sbo-service', 50053),
       ('pin-service', 50054),
       ('config-service', 50055),
       ('daf-service', 50056)
