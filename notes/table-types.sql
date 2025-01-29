CREATE TABLE all_data_types
(
    -- Scalar types (postgresToGo)
    int2_col          SMALLINT,
    int4_col          INTEGER,
    int8_col          BIGINT,
    numeric_col       NUMERIC,
    decimal_col       DECIMAL,
    serial_col        SERIAL,
    bigserial_col     BIGSERIAL,
    smallserial_col   SMALLSERIAL,
    float4_col        REAL,
    float8_col        DOUBLE PRECISION,

    varchar_col       VARCHAR(255),
    char_col          CHAR(10),
    text_col          TEXT,
    bpchar_col        CHAR(50), -- Blank-padded char
    name_col          NAME,

    bool_col          BOOLEAN,

    date_col          DATE,
    timestamp_col     TIMESTAMP,
    timestamptz_col   TIMESTAMPTZ,
    time_col          TIME,
    timetz_col        TIMETZ,
    interval_col      INTERVAL,

    json_col          JSON,
    jsonb_col         JSONB,

    uuid_col          UUID,

    xml_col           XML,
    tsvector_col      TSVECTOR,
    tsquery_col       TSQUERY,
    oid_col           OID,
    xid_col           XID,
    cid_col           CID,
    regclass_col      REGCLASS,
    regproc_col       REGPROC,
    regtype_col       REGTYPE,

    -- Array types (arrayTypes)
    bytea_col         BYTEA,
    int2_array_col    SMALLINT[],
    int4_array_col    INTEGER[],
    int8_array_col    BIGINT[],
    numeric_array_col NUMERIC[],
    text_array_col    TEXT[],
    uuid_array_col    UUID[],
    bool_array_col    BOOLEAN[],
    varchar_array_col VARCHAR[],
    date_array_col    DATE[],
    float4_array_col  REAL[],
    float8_array_col  DOUBLE PRECISION[]
);
