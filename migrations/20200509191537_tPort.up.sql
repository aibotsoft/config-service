create table dbo.Port
(
    Id          tinyint identity not null,
    ServiceName varchar(100)     not null,
    GrpcPort    int              not null,
    constraint PK_Port primary key (Id),
)