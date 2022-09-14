package influx

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
)

const (
	flushCount = 4
)

var (
	pointCount int

	influxdbClient influxdb2.Client
	bucket         string
	org            string
)

func InitInflux2() {
	influxdbClient = influxdb2.NewClient(viper.GetString("influxDB.url"), viper.GetString("influxDB.token"))
	bucket = viper.GetString("influxDB.bucket")
	org = viper.GetString("influxDB.org")
}

func CloseInflux2() {
	// always close client at the end
	if pointCount != 0 {
		writeAPI := influxdbClient.WriteAPI(org, bucket)
		writeAPI.Flush()
	}
	influxdbClient.Close()
}

func WriteRecord(measurement string, tags map[string]string, fields map[string]interface{}) {
	WriteRecordWithTime(measurement, tags, fields, time.Now())
}

func WriteRecordWithTime(measurement string, tags map[string]string, fields map[string]interface{}, t time.Time) {
	// write point asynchronously
	writeAPI := influxdbClient.WriteAPI(org, bucket)
	p := influxdb2.NewPoint(measurement, tags, fields, t)
	writeAPI.WritePoint(p)

	// Flush writes
	if pointCount++; (pointCount % flushCount) == 0 {
		writeAPI.Flush()
		pointCount = 0
	}
}
