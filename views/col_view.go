package views

import (
	"fmt"
	"strings"

	"github.com/ideaspaper/mer/entities"
)

type IColView interface {
	Search([]entities.Col)
	DisplayError(error)
	DidYouMean([]string)
	NoResult()
}

type colView struct {
}

func NewColView() IColView {
	return &colView{}
}

func searchHeadword(col *entities.Col) string {
	prsSlice := []string{}
	if len(col.Hwi.Prs) > 0 {
		for _, v := range col.Hwi.Prs {
			prsSlice = append(prsSlice, fmt.Sprintf("\x1b[32m\\\x1b[3m%v\x1b[0m\x1b[32m\\\x1b[0m", v.Mw))
		}
		prs := strings.Join(prsSlice, " | ")
		return fmt.Sprintf("\x1b[1;34m  ► %v\x1b[0m [ 🗣️  %v ]", strings.ToUpper(col.Hwi.Hw), prs)
	}
	return fmt.Sprintf("\x1b[1;34m  ► %v\x1b[0m", strings.ToUpper(col.Hwi.Hw))
}

func searchOffensive(col *entities.Col) string {
	if col.Meta.Offensive {
		return "\x1b[31m offensive\x1b[0m"
	}
	return ""
}

func searchFunctionalLabel(col *entities.Col) string {
	if col.Fl != "" {
		return fmt.Sprintf("\x1b[3;33m %s\x1b[0m", col.Fl)
	}
	return ""
}

func searchCrossReferences(col *entities.Col) string {
	if len(col.Cxs) > 0 {
		crossReferences := []string{}
		for _, v := range col.Cxs {
			sub := ""
			for _, w := range v.Cxtis {
				sub += w.Cxt
			}
			crossReferences = append(crossReferences, fmt.Sprintf("\x1b[3m%s\x1b[0m %s", v.Cxl, sub))
		}

		return fmt.Sprintf("\n    %s", strings.Join(crossReferences, ", "))
	}
	return ""
}

func splitString(str string, size int) []string {
	if len(str) <= size {
		return []string{str}
	}
	return append([]string{string(str[0:size])}, splitString(str[size:], size)...)
}

func searchShortdefs(col *entities.Col) string {
	shortdefsSlice := []string{}
	if len(col.Shortdef) > 0 {
		shortdefsSlice = append(shortdefsSlice, col.Shortdef...)
		for i, v := range shortdefsSlice {
			shortdefsSlice[i] = strings.Join(splitString(v, 80), "\n        ")
		}
		return "\n    \x1b[1mdefinitions\x1b[0m:\n\x1b[31m      ↪ \x1b[0m" + strings.Join(shortdefsSlice, "\x1b[31m\n      ↪ \x1b[0m")
	}
	return ""
}

func (cv *colView) Search(results []entities.Col) {
	fmt.Println("\x1b[1;34m╭─────────╮\x1b[0m")
	fmt.Println("\x1b[1;34m│ RESULTS │\x1b[0m")
	fmt.Println("\x1b[1;34m╰─────────╯\x1b[0m")
	for _, v := range results {
		fmt.Printf("%s%s%s%s%s\n", searchHeadword(&v), searchOffensive(&v), searchFunctionalLabel(&v), searchCrossReferences(&v), searchShortdefs(&v))
	}
}

func (cv *colView) DidYouMean(results []string) {
	fmt.Println("\x1b[1mDid you mean:\x1b[0m")
	displayResult := strings.Join(results, "\x1b[0m \x1b[37m|\x1b[0m \x1b[33m")
	fmt.Printf("\x1b[33m%s\x1b[0m\n", displayResult)
}

func (cv *colView) DisplayError(err error) {
	fmt.Println(err)
}

func (cv *colView) NoResult() {
	fmt.Println("\x1b[1;33m╭───────────╮\x1b[0m")
	fmt.Println("\x1b[1;33m│ NO RESULT │\x1b[0m")
	fmt.Println("\x1b[1;33m╰───────────╯\x1b[0m")
}
