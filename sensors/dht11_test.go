package sensors

import "testing"

func TestDHt11PinSetup(t *testing.T) {
	dht := DHT11{}
	dht.PinSetup(1, 2, 3, 4)
	t.Log(dht)
}
