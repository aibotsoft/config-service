create or alter proc dbo.uspSaveCurrency @TVP dbo.CurrencyType READONLY as
begin
    set nocount on

    MERGE dbo.Currency AS t
    USING @TVP s
    ON (t.Code = s.Code)

    WHEN MATCHED THEN
        UPDATE
        SET Value                = s.Value,
            UpdatedAt           = sysdatetimeoffset()

    WHEN NOT MATCHED THEN
        INSERT (Code, Value)
        VALUES (s.Code, s.Value);
end