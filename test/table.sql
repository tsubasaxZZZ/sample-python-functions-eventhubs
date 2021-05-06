CREATE TABLE fromfunctions (
    id varchar(256),
    process_datetime datetime,
    value_a real,
    value_b real,
    eh_enqueuedtimeutc datetime,
    eh_offset varchar(1024),
    eh_sequenceno bigint,
    sql_insert_datetime datetime,
    func_invocationid varchar(256)
)

--チェック用--
select count(*) from fromfunctions where id like 'test7_%' --order by eh_enqueuedtimeutc
SELECT  * FROM [dbo].[fromfunctions] order by eh_enqueuedtimeutc desc
truncate table fromfunctions
drop table fromfunctions
--
SELECT top(100) * FROM [dbo].[fromfunctions] where id like 'test7_%' order by process_datetime asc
SELECT top(100) * FROM [dbo].[fromfunctions] where id like 'test7_%' order by sql_insert_datetime desc