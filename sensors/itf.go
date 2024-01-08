package sensors

import "fmt"

type Sensor struct {
	
}
type Device struct {
}


type DHT11 struct {
	
}
// smoke, co, co2, no3
type MQ135 struct {
}

// butane gas leak detection
type MQ7 struct {
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

	return nil
}
