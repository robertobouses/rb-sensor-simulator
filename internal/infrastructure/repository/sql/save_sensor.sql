INSERT INTO sen.sensor (
    name,
    type,
    sampling_interval,
    alert_threshold,
    unit,
    enabled
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
