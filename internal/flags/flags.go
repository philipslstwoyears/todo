package flags

import (
	"errors"
	"flag"
)

type Flags struct {
	InputFile, OutputFile string
	Count                 bool // Флаг -c
	Duplicate             bool // Флаг -d
	Uniq                  bool // Флаг -u
	IgnoreCase            bool // Флаг -i
	NumFields             int  // Флаг -f
	NumChars              int  // Флаг -s
}

func New() (*Flags, error) {
	var f Flags
	flag.BoolVar(&f.Count, "c", false, "подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.")
	flag.BoolVar(&f.Duplicate, "d", false, "вывести только те строки, которые повторились во входных данных.")
	flag.BoolVar(&f.Uniq, "u", false, "вывести только те строки, которые не повторились во входных данных.")
	flag.BoolVar(&f.IgnoreCase, "i", false, "не учитывать регистр букв.")
	flag.IntVar(&f.NumFields, "f", 0, "не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.")
	flag.IntVar(&f.NumChars, "s", 0, "не учитывать первые num_chars символов в строке. При использовании вместе с параметром -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля).")

	flag.Parse()

	f.InputFile = flag.Arg(0)
	f.OutputFile = flag.Arg(1)

	if len(flag.Args()) > 2 || f.Count && f.Duplicate || f.Count && f.Uniq || f.Duplicate && f.Uniq {
		return nil, errors.New("uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}
	return &f, nil
}
