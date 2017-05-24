# go-expeval
Expression evaluation for golang

# expeval
--

```go
import "go-expeval"
```


## Usage

#### type ExpEval

```go
type ExpEval struct {
}
```

A ExpEval is expression evaluation

#### func  NewExpEval

```go
func NewExpEval(name, path string) *ExpEval
```
NewExpEval returns the *expeval.ExpEval

#### func (*ExpEval) Eval

```go
func (e *ExpEval) Eval(expr string, values map[string]interface{}) (interface{}, error)
```
Eval returns the evaluated value(interface{})

#### func (*ExpEval) EvalToBool

```go
func (e *ExpEval) EvalToBool(expr string, values map[string]interface{}) (bool, error)
```
EvalToBool returns the evaluated value(bool)

#### func (*ExpEval) EvalToFloat64

```go
func (e *ExpEval) EvalToFloat64(expr string, values map[string]interface{}) (float64, error)
```
EvalToFloat64 returns the evaluated value(float64)

#### func (*ExpEval) EvalToInt64

```go
func (e *ExpEval) EvalToInt64(expr string, values map[string]interface{}) (int64, error)
```
EvalToInt64 returns the evaluated value(int64)

#### func (*ExpEval) EvalToString

```go
func (e *ExpEval) EvalToString(expr string, values map[string]interface{}) (string, error)
```
EvalToString returns the evaluated value(string)

#### func (*ExpEval) EvalToUint64

```go
func (e *ExpEval) EvalToUint64(expr string, values map[string]interface{}) (uint64, error)
```
EvalToUint64 returns the evaluated value(uint64)


## Examples

```go
package main

import (
	"fmt"
	"github.com/tkuchiki/go-expeval"
	"log"
)

func main() {
	// same types.NewPackage("main", "main")
	eval := expeval.NewExpEval("main", "main")

	values := map[string]interface{}{
		"foo": 5,
	}
	tv, err := eval.Eval(`foo < 10 && foo > 1`, values)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tv.(bool))

	values = map[string]interface{}{
		"foo": 10,
	}
	b, err := eval.EvalToBool(`foo < 10 && foo > 1`, values)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)
}
```
