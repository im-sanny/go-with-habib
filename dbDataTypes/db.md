# Numeric Types:

- SMALLINT → 16-bit signed integer → Range: -32,768 to +32,767
- INTEGER / INT → 32-bit signed integer → Range: -2,147,483,648 to +2,147,483,647 (most commonly used)
- BIGINT → 64-bit signed integer → Range: -9.2 quintillion to +9.2 quintillion (use when INT might overflow)
- SERIAL → Auto-incrementing 32-bit integer (actually INTEGER + auto-created sequence) → Values: 1 to 2,147,483,647
- BIGSERIAL → Auto-incrementing 64-bit integer → Values: 1 to 9.2 quintillion (use for very large tables)
- SMALLSERIAL → Auto-incrementing 16-bit integer → 1 to 32,767 (rarely used)

# Floating-Point / Decimal Types:

- REAL → 32-bit (4-byte) floating point → ~6 decimal digits of precision
- DOUBLE PRECISION → 64-bit (8-byte) floating point → ~15 decimal digits of precision
  Warning: Floating-point numbers can have rounding errors → not good for money!
- NUMERIC(p, s) or DECIMAL(p, s) → Exact decimal numbers (no rounding error)
  Example: NUMERIC(10,2) → up to 10 digits total, 2 after decimal (e.g., 12345678.90)
  Best for money, financial data, or when exact precision is needed.

# Character / String Types:

- CHAR(n) → Fixed-length string, always uses n characters (padded with spaces)
  Example: CHAR(5) with value 'ab' → stored as 'ab '
  Rarely used → wastes space, not recommended
- VARCHAR(n) → Variable-length string with max limit of n characters
  Only uses as much space as needed + small overhead
  Recommended when you have a reasonable max length
- TEXT → Unlimited length string (up to 1 GB in PostgreSQL)
  Best choice when you don't know the length or want no limit
  Performance is almost same as VARCHAR(n) in PostgreSQL → use TEXT freely!

# Boolean:

- BOOLEAN → Stores true, false, or null
  Accepts: true/false, t/f, yes/no, 1/0

# Date & Time Types:

- DATE → Stores only date → e.g., 2025-09-20
- TIME [WITHOUT TIME ZONE] → Stores time of day → e.g., 13:50:42 (no timezone info)
- TIMESTAMPTZ → (Recommended) Timestamp with time zone → e.g., 2025-11-25 02:23:00+06
  Always stores in UTC internally, converts to local timezone on display
- TIMESTAMP WITHOUT TIME ZONE → Stores date + time but no timezone info
  Can cause confusion across timezones → avoid unless you have a reason
- INTERVAL → Stores a duration → e.g., 3 days, 2 hours 30 minutes

# Other Useful Types:

- UUID → For universally unique IDs → e.g., a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11
- JSON / JSONB → Store JSON data
- JSON → text storage
- JSONB → binary, faster to query, supports indexing → preferred
- ARRAY → Store arrays → e.g., TEXT[] for list of strings
- MONEY → Exists but not recommended (better to use NUMERIC for currency)
