package calculus

import (
	"math"
	"regexp"

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
	re := regexp.MustCompile(`([a-zA-Z0-9_.()]+)\s*\*\*\s*([a-zA-Z0-9_.()]+)`)
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
	// Calculate the derivative using the limit definition
	h := 1e-10
	return (f(x+h) - f(x-h)) / (2 * h)
}
func integral(f func(float64) float64, a, b float64, n int) float64 {
	// Calculate the integral using the trapezoidal rule
	h := (b - a) / float64(n)
	integral := 0.0
	for i := 0; i < n; i++ {
		x1 := a + float64(i)*h
		x2 := a + float64(i+1)*h
		integral += (f(x1) + f(x2)) * h / 2
	}
	return integral
}
func limit(f func(float64) float64, x float64) float64 {
	h := 1e-5       // start with a small step
	epsilon := 1e-8 // tolerance for convergence
	var prev, curr float64

	prev = f(x + h)
	for i := 0; i < 100; i++ {
		h /= 10
		curr = f(x + h)
		if math.Abs(curr-prev) < epsilon {
			return curr
		}
		prev = curr
	}

	// fallback if it didn't converge nicely
	return curr
}
