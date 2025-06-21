UPDATE sen.sensor
SET
    sampling_interval = COALESCE($1, sampling_interval),
    alert_threshold_min = COALESCE($2, alert_threshold_min),
    alert_threshold_max = COALESCE($3, alert_threshold_max),
    unit = COALESCE(NULLIF($4, ''), unit),
    error = COALESCE(NULLIF($5, ''), error),
    status = COALESCE(NULLIF($6, ''), status)
WHERE id = $7;
