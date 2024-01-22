package flags

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type color string

const (
	ColorBlack  color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color color, message string) {
	fmt.Println(string(color), message, ColorReset)
}

func Flag() {

	useColor := flag.Bool("color", false, "display colorized output")

	if *useColor {
		colorize(ColorBlue, "Hello, DigitalOcean!")
		return
	}
}

func Address() string {
	host := new(Roud)
	_ = flag.Value(host)

	flag.Var(host, "set_port", "Net address host:port")

	flag.Parse()

	return host.String()
}

type Roud struct {
	Host string
	Port string
}

func (h Roud) String() string {
	return h.Host + ":" + h.Port
}

func (h *Roud) Set(str string) error {
	hp := strings.Split(str, ":")
	if len(hp) != 2 {
		return errors.New("Need address in a form host:port")
	}
	port := hp[1]

	h.Host = hp[0]
	h.Port = port
	return nil
}

func FlagBase() {

	imgFile := flag.String("file", "", "input image file")
	destDir := flag.String("dest", "./output", "destination folder")
	width := flag.Int("w", 1024, "width of the image")
	isThumb := flag.Bool("thumb", false, "create thumb")

	// разбор командной строки
	flag.Parse()
	fmt.Println("Image file:", *imgFile)
	fmt.Println("Destination folder:", *destDir)
	fmt.Println("Width:", *width)
	fmt.Println("Thumbs:", *isThumb)

}
