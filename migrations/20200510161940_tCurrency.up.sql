create table dbo.Currency
(
    Code  varchar(6)    not null,
    Value decimal(9, 5) not null,
    constraint PK_Currency primary key (Code),
)
create type dbo.CurrencyType as table
(
    Code  varchar(6) not null,
    Value decimal(9, 5),
    primary key (Code)
)

insert into dbo.Currency (Code, Value)
values ('USD', 1),
       ('EUR', 0.93)