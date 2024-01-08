package sensors

import "testing"

func TestDHt11PinSetup(t *testing.T) {
	dht := &DHT11{}
	var sensor Sensor = dht
	dht.PinSetup(3, 1, 23, 39)
	t.Log(dht)
	sensor.PinSetup(3, 1, 34, 32)
	t.Log(sensor)
}
