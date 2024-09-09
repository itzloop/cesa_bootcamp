CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name varchar(255),
    created_at timestamp
);

INSERT INTO users (name, created_at)
SELECT
    'Bob' || i,
    time_hour
FROM 
    generate_series(1, 50000) as i,
    generate_series(
        TIMESTAMP '2020-11-01', 
        TIMESTAMP '2024-06-01', 
        INTERVAL '2 days'
    ) as time_hour;

