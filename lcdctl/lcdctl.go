package main

import (
	"log"
	"time"

	device "github.com/d2r2/go-hd44780"
	i2c "github.com/d2r2/go-i2c"
)

// Reference: https://github.com/d2r2/go-hd44780
func main() {
	// Create new connection to i2c-bus on 1 line with address 0x27.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2c, err := i2c.NewI2C(0x27, 1)
	if err != nil {
		log.Fatal(err)
	}
	// Free I2C connection on exit
	defer i2c.Close()
	// Construct lcd-device connected via I2C connection
	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	if err != nil {
		log.Fatal(err)
	}
	// Turn on the backlight
	err = lcd.BacklightOn()
	if err != nil {
		log.Fatal(err)
	}
	// Put text on 1 line of lcd-display
	err = lcd.ShowMessage("--=! Let's rock !=--", device.SHOW_LINE_1)
	if err != nil {
		log.Fatal(err)
	}
	// Wait 5 secs
	time.Sleep(5 * time.Second)
	// Output text to 2 line of lcd-screen
	err = lcd.ShowMessage("Welcome to RPi dude!", device.SHOW_LINE_2)
	if err != nil {
		log.Fatal(err)
	}
	// Wait 5 secs
	time.Sleep(5 * time.Second)

	lcd.Clear()

	// Turn off the backlight and exit
	err = lcd.BacklightOff()
	if err != nil {
		log.Fatal(err)
	}
}
