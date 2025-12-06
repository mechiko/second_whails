package templates

import (
	"fmt"
	"html/template"
	"korrectkm/domain"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// functions is a template.FuncMap providing custom date formatting functions for use in HTML templates.
var functions = template.FuncMap{
	"humanDate": func(t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return t.Format("02 Jan 2006 at 15:04")
	},
	"formatDate": func(t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return t.Format("2006-01-02")
	},
	"iterate": func(count int) []int {
		if count <= 0 {
			return []int{}
		}
		numbers := make([]int, count)
		for i := 0; i < count; i++ {
			numbers[i] = i
		}
		return numbers
	},
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
	"mul": func(a, b int) int {
		return a * b
	},
	"min": func(a, b int) int {
		if a < b {
			return a
		}
		return b
	},
	"inc": func(i int) int {
		return i + 1
	},
	"noescape": func(str string) template.HTML {
		return template.HTML(str)
	},
	"base": func(t string) string {
		if t == "" {
			return ""
		}
		return filepath.Base(t)
	},
	"groupByID": func(id int) string {
		m := domain.ProductGroupByID()
		return m[id].Name
	},
	"groupByIDs": func(id int) string {
		return domain.ProductGroupByIDs[id].Name
	},
	"fBalance": func(vol int) string {
		return formatMoney(float64(vol) / 2)
	},
	"ddmmyyhhmmss": func(t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return t.Format("02-01-2006 15:04:05")
	},
}

func formatMoney(value float64) string {
	// Convert to cents to avoid floating point inaccuracies for calculations
	cents := int64(value * 100)

	// Determine sign
	isNegative := cents < 0
	if isNegative {
		cents = -cents
	}

	// Separate dollars and cents
	dollars := cents / 100
	remainingCents := cents % 100

	// Format dollars with thousands separators
	dollarStr := strconv.FormatInt(dollars, 10)
	var parts []string
	for i := len(dollarStr) - 3; i >= 0; i -= 3 {
		if i == 0 {
			parts = append([]string{dollarStr[:3]}, parts...)
		} else {
			parts = append([]string{dollarStr[i : i+3]}, parts...)
		}
	}
	if len(dollarStr)%3 != 0 {
		parts = append([]string{dollarStr[:len(dollarStr)%3]}, parts...)
	}
	formattedDollars := strings.Join(parts, ",")

	// Format cents
	formattedCents := fmt.Sprintf("%02d", remainingCents)

	// Assemble the final string
	if isNegative {
		return fmt.Sprintf("%s.%s", formattedDollars, formattedCents)
	}
	return fmt.Sprintf("%s.%s", formattedDollars, formattedCents)
}
