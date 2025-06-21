INSERT INTO sen.sensor (
    name,
    type,
    sampling_interval,
    alert_threshold_min,
    alert_threshold_max,
    unit,
    error,
    status
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
