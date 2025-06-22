UPDATE sen.sensor
SET
    sampling_interval     = $1,
    alert_threshold_min   = $2,
    alert_threshold_max   = $3,
    unit                  = $4,
    status                = $5
WHERE id = $6;
