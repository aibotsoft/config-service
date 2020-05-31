create table dbo.Currency
(
    Code      varchar(6)                                    not null,
    Value     decimal(19, 9)                                 not null,
    UpdatedAt datetimeoffset(2) default sysdatetimeoffset() not null,
    constraint PK_Currency primary key (Code),
)
create type dbo.CurrencyType as table
(
    Code  varchar(6) not null,
    Value decimal(19, 9),
    primary key (Code)
)

insert into dbo.Currency (Code, Value)
values ('USD', 1),
       ('EUR', 0.93)