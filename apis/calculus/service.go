package calculus

import (
	"math"
	"randomapis/utilities"
	"regexp"
	"strconv"

	"github.com/Knetic/govaluate"
)

var functions = map[string]govaluate.ExpressionFunction{
	"sin": func(args ...interface{}) (interface{}, error) {
		return math.Sin(args[0].(float64)), nil
	},
	"cos": func(args ...interface{}) (interface{}, error) {
		return math.Cos(args[0].(float64)), nil
	},
	"tan": func(args ...interface{}) (interface{}, error) {
		return math.Tan(args[0].(float64)), nil
	},
	"exp": func(args ...interface{}) (interface{}, error) {
		return math.Exp(args[0].(float64)), nil
	},
	"log": func(args ...interface{}) (interface{}, error) {
		return math.Log(args[0].(float64)), nil
	},
	"ln": func(args ...interface{}) (interface{}, error) {
		return math.Log(args[0].(float64)), nil
	},
	"sqrt": func(args ...interface{}) (interface{}, error) {
		return math.Sqrt(args[0].(float64)), nil
	},
	"abs": func(args ...interface{}) (interface{}, error) {
		return math.Abs(args[0].(float64)), nil
	},
	"pow": func(args ...interface{}) (interface{}, error) {
		return math.Pow(args[0].(float64), args[1].(float64)), nil
	},
	"floor": func(args ...interface{}) (interface{}, error) {
		return math.Floor(args[0].(float64)), nil
	},
	"ceil": func(args ...interface{}) (interface{}, error) {
		return math.Ceil(args[0].(float64)), nil
	},
}

func insertMultiplication(expr string) string {
	// Insert '*' between number and variable (e.g. 4x -> 4*x)
	re := regexp.MustCompile(`(\d)([a-zA-Z])`)
	return re.ReplaceAllString(expr, `$1*$2`)
}

func replaceAbsolute(expr string) string {
	// Replace |expr| with abs(expr)
	re := regexp.MustCompile(`\|([^\|]+)\|`)
	return re.ReplaceAllString(expr, "abs($1)")
}

func replaceExponent(expr string) string {
	// Replace occurrences of something like x**2 or (x+1)**3
	re := regexp.MustCompile(`(\([^\)]+\)|[a-zA-Z0-9_.]+)\s*\*\*\s*(\([^\)]+\)|[a-zA-Z0-9_.]+)`)
	return re.ReplaceAllString(expr, "pow($1, $2)")
}

func parseFunction(expr string) (func(float64) float64, error) {
	expr = insertMultiplication(expr) // optional: handle 2x → 2*x
	expr = replaceExponent(expr)      // optional: handle x**2 → pow(x, 2)
	expr = replaceAbsolute(expr)      // optional: handle |x| → abs(x)
	parsed, err := govaluate.NewEvaluableExpressionWithFunctions(expr, functions)
	if err != nil {
		return nil, err
	}
	return func(x float64) float64 {
		result, err := parsed.Evaluate(map[string]interface{}{"x": x})
		if err != nil {
			return 0
		}
		return result.(float64)
	}, nil
}
func derivative(f func(float64) float64, x float64) float64 {
	h := 1e-5
	d1 := (f(x+h) - f(x-h)) / (2 * h)
	h /= 2
	d2 := (f(x+h) - f(x-h)) / (2 * h)
	return utilities.RoundTo5Decimals((4*d2 - d1) / 3) // Richardson extrapolation
}

func integral(f func(float64) float64, a, b float64, n int) float64 {
	if n%2 != 0 {
		n++
	}
	h := (b - a) / float64(n)
	sum := f(a) + f(b)

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}

	return utilities.RoundTo5Decimals((h / 3) * sum)
}

func limit(f func(float64) float64, x float64) float64 {
	h := 1e-5
	epsilon := 1e-8
	var prev, curr float64

	prev = f(x + h)
	for i := 0; i < 100; i++ {
		h /= 10
		curr = f(x + h)
		if math.Abs(curr-prev) < epsilon {
			return utilities.RoundTo5Decimals(curr)
		}
		prev = curr
	}
	return utilities.RoundTo5Decimals(curr)
}

func evaluate2DFunction(f func(float64) float64, upper_bound, lower_bound float64) [][]float64 {
	var points [][]float64
	xs := []float64{}
	ys := []float64{}
	for x := lower_bound; x <= upper_bound; x += 0.05 {
		xs = append(xs, x)
		ys = append(ys, f(x))
	}
	points = append(points, xs)
	points = append(points, ys)
	return points
}

func evaluate2DFunctionPoint(f func(float64) float64, x float64) []float64 {
	var points []float64
	points = append(points, x)
	points = append(points, f(x))
	return points
}

func tangentLine(f func(float64) float64, x, upper_bound, lower_bound float64) [][]float64 {
	m := derivative(f, x)
	b := f(x) - m*x
	expr := strconv.FormatFloat(m, 'f', -1, 64) + "x + " + strconv.FormatFloat(b, 'f', -1, 64)
	tangentFunc, err := parseFunction(expr)
	if err != nil {
		return nil
	}
	points := evaluate2DFunction(tangentFunc, upper_bound, lower_bound)
	return points
}

func lineFrom2Points(x1, y1, x2, y2, upper_bound, lower_bound float64) [][]float64 {
	if x1 == x2 {
		return nil
	}
	m := (y2 - y1) / (x2 - x1)
	b := y1 - m*x1
	expr := strconv.FormatFloat(m, 'f', -1, 64) + "x + " + strconv.FormatFloat(b, 'f', -1, 64)
	f, err := parseFunction(expr)
	if err != nil {
		return nil
	}
	points := evaluate2DFunction(f, upper_bound, lower_bound)
	return points
}
