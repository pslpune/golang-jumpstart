package sensors

import "fmt"

/*
	Java (explicit) - If an class implements an interface it has to implement all the functions stated by the interface.
	C ++ - -++--
	you would have to write the interface definitions ahead of the class definitions



	go (implicit) - if the struct has the method that is stated by the interface, it is said to implement this interface
	interface definitions can come after the class definitions
*/

type Sensor interface {
	PinSetup(pwr, gnd, trigg, op int) error
}

type Power interface {
	PowerOn() error
	PowerOff()
}

type Device struct {
}

type DHT11 struct {
	PWRpin, GNDpin, OPpin int
}

func (dht *DHT11) PowerOn() error {
	return nil
}

func (dht *DHT11) PowerOff() {

}

// this is called monkey-patching
func (dht *DHT11) PinSetup(pwr, gnd, trigg, op int) error {
	dht.PWRpin = pwr
	dht.GNDpin = gnd
	dht.OPpin = op
	return nil
}

// smoke, co, co2, no3
type MQ135 struct {
	PWRpin, GNDpin, OPpin, Triggpin int
}

func (mq135 *MQ135) PowerOn() error {
	return nil
}

func (mq135 *MQ135) PowerOff() {
	
}

func (mq135 *MQ135) PinSetup(pwr, gnd, trigg, op int) error {
	mq135.PWRpin = pwr
	mq135.GNDpin = gnd
	mq135.OPpin = op
	mq135.Triggpin = trigg
	return nil
}

// butane gas leak detection
type MQ7 struct {
	PWRpin, GNDpin, OPpin, Triggpin int
}

func (mq7 *MQ7) PinSetup(pwr, gnd, trigg, op int) error {
	mq7.PWRpin = pwr
	mq7.GNDpin = gnd
	mq7.OPpin = op
	mq7.Triggpin = trigg
	return nil
}

type LM35 struct {
	PWRpin, GNDpin, OPpin int
}

func (lm *LM35) PinSetup(pwr, gnd, trigg, op int) error {
	lm.PWRpin = pwr
	lm.GNDpin = gnd
	lm.OPpin = op
	return nil
}

// func PinSetupDHT(dht *DHT11, pwr, gnd, trigg, op int) error {
// 	return nil
// }
// func PinSetupMQ135(mq *MQ135, pwr, gnd, trigg, op int) error {
// 	return nil
// }

func PinSetup(sensor interface{}, pwr, gnd, trigg, op int) error {
	_, ok := sensor.(*DHT11)
	if ok {
		// Do all that is required for DHT11
		sensor.(*DHT11).PWRpin = pwr
		sensor.(*DHT11).GNDpin = gnd
		sensor.(*DHT11).OPpin = op
	}
	_, ok = sensor.(*MQ135)
	if ok {
		// do all what it takes to setup the pins for mq135
		sensor.(*MQ135).PWRpin = pwr
		sensor.(*MQ135).GNDpin = gnd
		sensor.(*MQ135).Triggpin = trigg
		sensor.(*MQ135).OPpin = op

	}
	_, ok = sensor.(*MQ7)
	if ok {
		// do all what it takes to setup the pins for mq135
		sensor.(*MQ7).PWRpin = pwr
		sensor.(*MQ7).GNDpin = gnd
		sensor.(*MQ7).Triggpin = trigg
		sensor.(*MQ7).OPpin = op

	}
	return nil
}

type SetOfPinsForCalibrate struct {
	Pwr, Gnd, Trigg, Op int
}

// interface{} - does NOT mean instantiation, nor does it mean declaration, its a predefined data type regarded as grandfather of all types
// its an universal datatype , grandfather of types if you may
// O - open closed principle - open for extension closed for modification
func Calibrate(typeOfSensor string, pins SetOfPinsForCalibrate) (Sensor, error) {
	var sensor Sensor
	if typeOfSensor == "DHT11" {
		sensor = &DHT11{}
		// PinSetupDHT(dht, 1, 6, 15, 36)
		// PinSetup(dht, 1, 6, 15, 36)
		// dht.PinSetup(1, 6, 15, 36)
		// return sensor, nil
	} else if typeOfSensor == "MQ135" {
		sensor = &MQ135{}
		// PinSetupMQ135(mq135, 1, 6, 15, 36)
		// PinSetup(mq135, 2, 9, 16, 38)
		// mq135.PinSetup(2, 9, 16, 38)
		// return mq135, nil
	} else if typeOfSensor == "MQ7" {
		sensor = &MQ7{}
		// PinSetupMQ135(mq135, 1, 6, 15, 36)
		// PinSetup(mq7, 2, 9, 16, 38)
		// mq7.PinSetup(2, 9, 16, 38)
		// return mq7, nil
	}
	sensor.PinSetup(pins.Pwr, pins.Gnd, pins.Trigg, pins.Op)
	return sensor, nil
}

// AnyObject
func Setup() error {
	sensor, err := Calibrate("DHT11", SetOfPinsForCalibrate{Pwr: 1, Gnd: 6, Trigg: 15, Op: 36})
	if err != nil || sensor == nil {
		return fmt.Errorf("failed setup , try again with other params")
	}
	// always remember what is the object underlying - &DHT11{}
	sensor.(Power).PowerOn()
	sensor.(Power).PowerOff()

	sensor, err = Calibrate("MQ135", SetOfPinsForCalibrate{Pwr: 2, Gnd: 9, Trigg: 17, Op: 39})
	if err != nil || sensor == nil {
		return fmt.Errorf("failed setup , try again with other params")
	}
	sensor.(Power).PowerOn()
	sensor.(Power).PowerOff()

	sensor, err = Calibrate("MQ7", SetOfPinsForCalibrate{Pwr: 2, Gnd: 9, Trigg: 17, Op: 39})
	if err != nil || sensor == nil {
		return fmt.Errorf("failed setup , try again with other params")
	}
	power, ok :=sensor.(Power)
	if !ok || power == nil {
		fmt.Println("mq7 sensor does not implement the power interface")
	}
	sensor.(Power).PowerOn()
	return nil
}
