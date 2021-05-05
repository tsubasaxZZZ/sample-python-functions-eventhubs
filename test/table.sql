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
