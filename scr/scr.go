package scr

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

type Buff struct {
	Wbus uint8
	Abus uint8
	Bbus uint8
	Obus uint8
	Rbus uint8
	Ibus uint8
	Cbus uint16

	PC   uint8
	Acc  uint8
	Alu  uint8
	BReg uint8
	OReg uint8
	Dis  uint8
	Mar  uint8
	Ram  [16]uint8
	Ir   uint8
	Cntl uint16
	Mic  []uint16
	Cnt  uint16
}

type Screen struct {
	Name string
	Wbus string
	Abus string
	Bbus string
	Obus string
	Rbus string
	Ibus string
	Cbus string
	PC   string

	Acc  string
	Alu  string
	BReg string
	OReg string
	Dis  string
	Mar  string
	Ram  string
	Ir   string
	Cntl string
	Cnt  string
}

func NewScreen() *Screen {
	fmt.Print(ClearScreen)

	return &Screen{
		Name: Blue + "SAP-1" + Reset,
	}
}

var Ticker = -1
var RamTicker = -1
var MicroTicker = 0

func (display *Screen) Tick(buff Buff) {

	if Ticker%2 == 1 {
		MicroTicker++
		MicroTicker = MicroTicker % 6
	}

	if Ticker%12 == 0 {
		RamTicker++
	}

	update := display.ParseDbus(buff)
	display.UpdateData(update)
	layout := display.GetLayout()
	template := display.FillRam(RamTicker, buff, layout)
	template = display.FillMicro(MicroTicker, buff, template)
	template = display.CreateTemplate(template)
	template = display.ColorSignals(buff, template)
	template = display.ColorLines(buff, template)
	template = display.SevenSegment(buff, template)

	fmt.Print(CursorHide)
	fmt.Print(CursorHome)
	display.Display(template)
	fmt.Printf("\x1b[%d;%dH", 38, 10)
	fmt.Print(CursorShow)

	Ticker++
}

func (display *Screen) SevenSegment(buff Buff, str string) string {

	aNum := display.Convert(int(buff.Dis))

	var newStr string
	newStr = strings.Replace(str, " _  _  _ ", Bold+Red+aNum[0]+Reset, 1)
	newStr = strings.Replace(newStr, " _||_ |_ ", Bold+Red+aNum[1]+Reset, 1)
	newStr = strings.Replace(newStr, "|_  _| _|", Bold+Red+aNum[2]+Reset, 1)
	return newStr
}

func (display *Screen) ParseDbus(buff Buff) map[string]string {

	upd := map[string]string{
		"Wbus": Green + fmt.Sprintf("%08b", buff.Wbus) + Reset,
		"Abus": Green + fmt.Sprintf("%08b", buff.Abus) + Reset,
		"Bbus": Green + fmt.Sprintf("%08b", buff.Bbus) + Reset,
		"Obus": Green + fmt.Sprintf("%08b", buff.Obus) + Reset,
		"Rbus": Green + fmt.Sprintf("%04b", buff.Rbus) + Reset,
		"Ibus": Green + fmt.Sprintf("%04b", buff.Ibus) + Reset,
		"Cbus": Green + fmt.Sprintf("%012b", buff.Cbus) + Reset,
		"PC":   Yellow + fmt.Sprintf("%04b", buff.PC) + Reset,
		"Acc":  Yellow + fmt.Sprintf("%08b", buff.Acc) + Reset,
		"Alu":  Yellow + fmt.Sprintf("%08b", buff.Alu) + Reset,
		"BReg": Yellow + fmt.Sprintf("%08b", buff.BReg) + Reset,
		"OReg": Yellow + fmt.Sprintf("%08b", buff.OReg) + Reset,
		"Mar":  Yellow + fmt.Sprintf("%04b", buff.Mar) + Reset,
		"Ram":  Yellow + fmt.Sprintf("%08b", (buff.Ram[buff.Mar])) + Reset,
		"Ir":   Yellow + fmt.Sprintf("%08b", buff.Ir) + Reset,
		"Cntl": Yellow + fmt.Sprintf("%012b", buff.Cntl) + Reset,
		"Cnt":  Yellow + fmt.Sprintf("%02d", buff.Cnt) + Reset,
	}

	return upd
}

func (display *Screen) UpdateData(newData map[string]string) {
	displayValue := reflect.ValueOf(display).Elem()
	typeOfDisplay := displayValue.Type()

	for key, value := range newData {
		for i := 0; i < displayValue.NumField(); i++ {
			if typeOfDisplay.Field(i).Name == key {
				field := displayValue.Field(i)
				if field.CanSet() {
					switch field.Kind() {
					case reflect.String:
						field.SetString(value)
					default:
						fmt.Printf("Unsupported field type %s in the struct.\n", field.Kind())
					}
				}
			}
		}
	}
}

func (display *Screen) FillRam(active int, buff Buff, str string) string {
	newStr := str

	for i := 0; i < 16; i++ {
		decorate := Green
		if i == active {
			decorate = BgGreen + Black
		}
		newStr = strings.Replace(newStr, fmt.Sprintf("RAM%02d 00", i), decorate+fmt.Sprintf("%08b", buff.Ram[i])+Reset, 1)
	}

	return newStr
}

func (display *Screen) FillMicro(active int, buff Buff, str string) string {
	newStr := str

	count := len(buff.Mic)

	for i := 0; i < count; i++ {
		decorate := Cyan
		if i == active {
			decorate = BgCyan + Black
		}
		newStr = strings.Replace(newStr, fmt.Sprintf("MC%02d 0000000", i), decorate+fmt.Sprintf("%012b", buff.Mic[i])+Reset, 1)
	}

	for i := count; i < 6; i++ {
		newStr = strings.Replace(newStr, fmt.Sprintf("MC%02d 0000000", i), "            "+Reset, 1)
	}

	return newStr
}

func (display *Screen) ColorSignals(buff Buff, str string) string {
	signalMapEq := map[uint16]string{
		1: "L̅o",
		2: "L̅b",
		//32:  "ɸ─ L̅a ──────┐",
		32:  "L̅a",
		64:  "E̅i",
		128: "L̅i",
		256: "C̅E",
		512: "L̅m",
	}

	signalMapNeq := map[uint16]string{
		4:    "Eu",
		8:    "Su",
		16:   "Ea",
		1024: "Ep",
		2048: "Cp",
	}

	var replacements []string
	for mask, signal := range signalMapEq {
		if (buff.Cntl & mask) == 0 {
			replacements = append(replacements, signal, Red+signal+Reset)
		}
	}

	for mask, signal := range signalMapNeq {
		if (buff.Cntl & mask) != 0 {
			replacements = append(replacements, signal, Red+signal+Reset)
		}
	}

	replacer := strings.NewReplacer(replacements...)
	return replacer.Replace(str)
}

func (display *Screen) ColorLines(buff Buff, str string) string {
	newStr := str

	// var LLi = "└───────────────────────────────────────────────────┘"

	// newStr := strings.Replace(newStr, LLi, Red+LLi+Reset, 1)

	return newStr
}

func (display *Screen) Display(Temp string) {

	tmpl, err := template.New("Screen").Parse(Temp)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, display)
	if err != nil {
		panic(err)
	}

	result := buf.String()
	println(result)
}

func (display *Screen) CreateTemplate(str string) string {
	replacer := strings.NewReplacer(
		"SAP-1", "{{.Name}}",
		"PC 0", "{{.PC}}",
		"DBUS 000", "{{.Wbus}}",
		"ABUS 000", "{{.Abus}}",
		"BBUS 000", "{{.Bbus}}",
		"OBUS 000", "{{.Obus}}",
		"RBUS", "{{.Rbus}}",
		"IBUS", "{{.Ibus}}",
		"CBUS 0000000", "{{.Cbus}}",
		"ACC 0000", "{{.Acc}}",
		"ALU 0000", "{{.Alu}}",
		"REG 0000", "{{.BReg}}",
		"OUT 0000", "{{.OReg}}",
		"DIS 0000", "{{.Dis}}",
		"MAR0", "{{.Mar}}",
		"RAM 0000", "{{.Ram}}",
		"INS 0000", "{{.Ir}}",
		"CON 00000000", "{{.Cntl}}",
		"XX", "{{.Cnt}}",
	)
	return replacer.Replace(str)
}

func (display *Screen) GetLayout() string { return SCR }
