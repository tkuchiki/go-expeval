package expeval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExpEval(t *testing.T) {
	assert := assert.New(t)

	e1 := &ExpEval{
		name: "expeval",
		path: "expeval",
	}

	e2 := NewExpEval("expeval", "expeval")

	assert.Equal(e1, e2, "not equal")
}

func TestEval(t *testing.T) {
	assert := assert.New(t)

	eval := NewExpEval("expeval", "expeval")

	values := map[string]interface{}{
		"foo": 1,
	}
	i, err := eval.Eval(`foo + 1`, values)
	i64 := i.(int64)

	assert.Equal(i64, int64(2), "not equal")
	assert.Nil(err)

	values = map[string]interface{}{
		"foo": 1.2,
	}
	i, err = eval.Eval(`foo + 1.3`, values)
	f64 := i.(float64)

	assert.Equal(f64, float64(2.5), "not equal")
	assert.Nil(err)

	values = map[string]interface{}{
		"foo": 1,
	}
	i, err = eval.Eval(`foo == 1`, values)
	b := i.(bool)

	assert.Equal(b, true, "not equal")
	assert.Nil(err)

	i, err = eval.Eval(`foo != 1`, values)
	b = i.(bool)

	assert.Equal(b, false, "not equal")
	assert.Nil(err)

	values = map[string]interface{}{
		"foo": "foo",
	}
	i, err = eval.Eval(`foo + "bar"`, values)
	s := i.(string)

	assert.Equal(s, "foobar", "not equal")
	assert.Nil(err)

	values = map[string]interface{}{
		"foo": "bar",
	}
	i, err = eval.Eval(`foo + 1`, values)

	var emptyInterface interface{}

	assert.Equal(emptyInterface, i, "not equal")
	assert.Equal(err, invalidValueErr, "invalid error")

	i, err = eval.Eval(`1 / 0`, values)

	assert.Equal(emptyInterface, i, "not equal")
	assert.Equal(err, invalidValueErr, "invalid error")
}

func TestEvalToBool(t *testing.T) {
	assert := assert.New(t)

	eval := NewExpEval("expeval", "expeval")

	values := map[string]interface{}{
		"foo": 1,
	}
	b, err := eval.EvalToBool(`foo == 1`, values)

	assert.Equal(b, true, "not equal")
	assert.Nil(err)

	values = map[string]interface{}{
		"foo": 1,
	}
	b, err = eval.EvalToBool(`foo + 1`, values)

	assert.Equal(b, false, "not equal")
	assert.Equal(err, invalidValueErr, "invalid error")
}

func TestEvalToString(t *testing.T) {
	assert := assert.New(t)

	eval := NewExpEval("expeval", "expeval")

	values := map[string]interface{}{
		"foo": "foo",
	}
	s, err := eval.EvalToString(`foo + "bar"`, values)

	assert.Equal(s, "foobar", "not equal")
	assert.Nil(err)

	values = map[string]interface{}{
		"foo": "foo",
	}
	s, err = eval.EvalToString(`foo / 1`, values)

	assert.Equal(s, "", "not equal")
	assert.Equal(err, invalidValueErr, "invalid error")
}

func TestEvalToInt64(t *testing.T) {
	assert := assert.New(t)

	eval := NewExpEval("expeval", "expeval")

	values := map[string]interface{}{
		"foo": 1,
	}
	i64, err := eval.EvalToInt64(`foo + 1`, values)

	assert.Equal(i64, int64(2), "not equal")
	assert.Nil(err)

	values = map[string]interface{}{
		"foo": "foo",
	}
	i64, err = eval.EvalToInt64(`foo / 1`, values)

	assert.Equal(i64, int64(0), "not equal")
	assert.Equal(err, invalidValueErr, "invalid error")
}

func TestEvalToUint64(t *testing.T) {
	assert := assert.New(t)

	eval := NewExpEval("expeval", "expeval")

	values := map[string]interface{}{
		"foo": uint64(1),
	}
	ui64, err := eval.EvalToUint64(`foo + 1`, values)

	assert.Equal(ui64, uint64(2), "not equal")
	assert.Nil(err)

	values = map[string]interface{}{
		"foo": "foo",
	}
	ui64, err = eval.EvalToUint64(`foo / 1`, values)

	assert.Equal(ui64, uint64(0), "not equal")
	assert.Equal(err, invalidValueErr, "invalid error")
}

func TestEvalToFloat64(t *testing.T) {
	assert := assert.New(t)

	eval := NewExpEval("expeval", "expeval")

	values := map[string]interface{}{
		"foo": 1.2,
	}
	f64, err := eval.EvalToFloat64(`foo + 1.3`, values)

	assert.Equal(f64, float64(2.5), "not equal")
	assert.Nil(err)

	values = map[string]interface{}{
		"foo": "foo",
	}
	f64, err = eval.EvalToFloat64(`foo / 1`, values)

	assert.Equal(f64, float64(0), "not equal")
	assert.Equal(err, invalidValueErr, "invalid error")
}
