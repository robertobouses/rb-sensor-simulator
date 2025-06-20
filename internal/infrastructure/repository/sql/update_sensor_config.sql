UPDATE sen.sensor
SET sampling_interval = $1,
	alert_threshold = $2,
	unit = $3,
	enabled = $4
WHERE id = $5