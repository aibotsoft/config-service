create table dbo.Currency
(
    Id    tinyint identity not null,
    Code  varchar(6)       not null,
    Value decimal(9, 5)    not null,
    constraint PK_Currency primary key (Id),
)

insert into dbo.Currency (Code, Value) values ('USD', 1), ('EUR', 0.93)