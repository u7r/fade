package fade

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GreenToBlue(x string) string {
	str := ""

	blue := 100

	for _, v := range strings.Split(x, "\n") {
		str += fmt.Sprintf("\033[38;2;0;255;%vm%v\033[0m\n", blue, v)
		if blue != 255 {
			blue += 25
			if blue > 255 {
				blue = 255
			}
		}
	}

	return str
}

func YellowToOrange(x string) string {
	str := ""

	green := 250

	for _, v := range strings.Split(x, "\n") {
		str += fmt.Sprintf("\033[38;2;255;%v;0m%v\033[0m\n", green, v)

		if green != 0 {
			green -= 25
			if green < 0 {
				green = 0
			}
		}
	}

	return str
}

func GreenToYellow(x string) string {
	str := ""

	red := 0

	for _, v := range strings.Split(x, "\n") {
		str += fmt.Sprintf("\033[38;2;%v;255;0m%v\033[0m\n", red, v)

		if red < 200 {
			red += 30
		}
	}

	return str
}

func PurpleToPink(x string) string {
	str := ""

	red := 40

	for _, v := range strings.Split(x, "\n") {
		str += fmt.Sprintf("\033[38;2;%v;0;220m%v\033[0m\n", red, v)
		if red != 255 {
			red += 15
			if red > 255 {
				red = 255
			}
		}
	}

	return str
}

func BlackToWhite(x string) string {
	str := ""

	red := 20
	green := 20
	blue := 20

	for _, v := range strings.Split(x, "\n") {
		str += fmt.Sprintf("\033[38;2;%v;%v;%vm%v\033[0m\n", red, green, blue, string(v))

		if red != 255 && green != 255 && blue != 255 {
			red += 20
			green += 20
			blue += 20
			if red > 255 && green > 255 && blue > 255 {
				red = 255
				green = 255
				blue = 255
			}
		}
	}

	return str
}

func PurpleToBlue(x string) string {
	str := ""

	red := 110

	for _, v := range strings.Split(x, "\n") {
		str += fmt.Sprintf("\033[38;2;%v;0;255m%v\033[0m\n", red, v)

		if red != 0 {
			red -= 10
			if red < 0 {
				red = 0
			}
		}
	}

	return str
}

func BlueToCyan(x string) string {
	str := ""

	green := 10

	for _, v := range strings.Split(x, "\n") {
		str += fmt.Sprintf("\033[38;2;0;%v;255m%v\033[0m\n", green, v)

		if green != 0 {
			green += 15
			if green < 0 {
				green = 0
			}
		}
	}
	return str
}

func RedToOrange(x string) string {
	str := ""

	green := 10

	for _, v := range strings.Split(x, "\n") {
		str += fmt.Sprintf("\033[38;2;255;%v;0m%v\033[0m\n", green, v)

		if green != 0 {
			green += 15
			if green < 0 {
				green = 0
			}
		}
	}
	return str
}

func PinkToRed(x string) string {
	str := ""

	blue := 255

	for _, v := range strings.Split(x, "\n") {
		str += fmt.Sprintf("\033[38;2;255;0;%vm%v\033[0m\n", blue, v)

		if blue != 0 {
			blue -= 20
			if blue < 0 {
				blue = 0
			}
		}
	}

	return str
}

func Randomness(x string) string {
	rand.Seed(time.Now().Unix())

	str := ""

	for _, v := range strings.Split(x, "\n") {
		for _, v1 := range v {
			str += fmt.Sprintf("\033[38;2;%v;%v;%vm%v\033[0m", rand.Intn(255), rand.Intn(255), rand.Intn(255), string(v1))
		}
		str += "\n"
	}

	return str
}
