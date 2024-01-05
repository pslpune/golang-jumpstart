package sensors

import "fmt"

type Sensor interface {
	
}
type Device struct {
}

// DHT11 is a type of Sensor
// DHT11 contains all the properties sensor
// dht11 - temp & RH
type DHT11 struct {
	Callout   string
	SnsrField Sensor // Anonymous field specification
	Dev       Device // Anon field specification
}

// smoke, co, co2, no3
type MQ135 struct {
}

// butane gas leak detection
type MQ7 struct {
}
type Motherboard struct {
	string // anon string field
}

// interface{} - does NOT mean instantiation, nor does it mean declaration, its a predefined data type regarded as grandfather of all types
// its an universal datatype , grandfather of types if you may
// O - open closed principle - open for extension closed for modification
func Calibrate(typeOfSensor string, calibration float32) (interface{}, error) {
	if typeOfSensor == "DHT11" {
		return &DHT11{}, nil
	} else if typeOfSensor == "MQ135" {
		return &MQ135{}, nil
	} else if typeOfSensor == "MQ7" {
		return &MQ7{}, nil
	}
	return nil, nil
}

// AnyObject
func Setup() error {
	itf, err := Calibrate("DHT11", 9.0)
	if err != nil || itf == nil {
		return fmt.Errorf("failed setup , try again with other params")
	}

	ptrDHT11 := itf.(*DHT11) //downcasting
	fmt.Println(ptrDHT11.SnsrField.ID)
	fmt.Println(ptrDHT11.Callout)
}
