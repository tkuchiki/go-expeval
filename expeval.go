package expeval

import (
	"errors"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
)

var (
	unknownInt64Err   = errors.New("unknown int64 value")
	unknownUint64Err  = errors.New("unknown uint64 value")
	unknownFloat64Err = errors.New("unknown float64 value")
	invalidValueErr   = errors.New("invalid value")
)

// A ExpEval is expression evaluation
type ExpEval struct {
	name string
	path string
}

// NewExpEval returns the *expeval.ExpEval
func NewExpEval(name, path string) *ExpEval {
	return &ExpEval{
		name: name,
		path: path,
	}
}

func (e *ExpEval) setValues(pkg *types.Package, values map[string]interface{}) {
	for key, val := range values {
		r := reflect.ValueOf(val)
		switch r.Kind() {
		case reflect.Bool:
			v := r.Bool()
			pkg.Scope().Insert(types.NewConst(token.NoPos, pkg, key, types.Typ[types.Bool], constant.MakeBool(v)))
		case reflect.String:
			v := r.String()
			pkg.Scope().Insert(types.NewConst(token.NoPos, pkg, key, types.Typ[types.String], constant.MakeString(v)))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := r.Int()
			pkg.Scope().Insert(types.NewConst(token.NoPos, pkg, key, types.Typ[types.Int64], constant.MakeInt64(v)))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := r.Uint()
			pkg.Scope().Insert(types.NewConst(token.NoPos, pkg, key, types.Typ[types.Uint64], constant.MakeUint64(v)))
		case reflect.Float32, reflect.Float64:
			v := r.Float()
			pkg.Scope().Insert(types.NewConst(token.NoPos, pkg, key, types.Typ[types.Float64], constant.MakeFloat64(v)))
		}
	}
}

// Eval returns the evaluated value(interface{})
func (e *ExpEval) Eval(expr string, values map[string]interface{}) (interface{}, error) {
	pkg := types.NewPackage(e.name, e.path)
	e.setValues(pkg, values)

	var i interface{}
	tav, err := types.Eval(token.NewFileSet(), pkg, token.NoPos, expr)
	if err != nil {
		return i, err
	}

	if tav.Type == nil {
		return i, invalidValueErr
	}

	var b bool
	switch tav.Type.(*types.Basic).Kind() {
	case types.Bool, types.UntypedBool:
		i = constant.BoolVal(tav.Value)
	case types.String, types.UntypedString:
		i = constant.StringVal(tav.Value)
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64, types.UntypedInt:
		i, b = constant.Int64Val(tav.Value)

		if !b {
			return i, unknownInt64Err
		}
	case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
		i, b = constant.Uint64Val(tav.Value)

		if !b {
			return i, unknownUint64Err
		}
	case types.Float32, types.Float64, types.UntypedFloat:
		i, b = constant.Float64Val(tav.Value)

		if !b {
			return i, unknownFloat64Err
		}
	}

	return i, nil
}

// EvalToBool returns the evaluated value(bool)
func (e *ExpEval) EvalToBool(expr string, values map[string]interface{}) (bool, error) {
	pkg := types.NewPackage(e.name, e.path)
	e.setValues(pkg, values)

	tav, err := types.Eval(token.NewFileSet(), pkg, token.NoPos, expr)
	if err != nil {
		return false, err
	}

	return constant.BoolVal(tav.Value), nil
}

// EvalToString returns the evaluated value(string)
func (e *ExpEval) EvalToString(expr string, values map[string]interface{}) (string, error) {
	pkg := types.NewPackage(e.name, e.path)
	e.setValues(pkg, values)

	tav, err := types.Eval(token.NewFileSet(), pkg, token.NoPos, expr)
	if err != nil {
		return "", err
	}

	return constant.StringVal(tav.Value), nil
}

// EvalToInt64 returns the evaluated value(int64)
func (e *ExpEval) EvalToInt64(expr string, values map[string]interface{}) (int64, error) {
	pkg := types.NewPackage(e.name, e.path)
	e.setValues(pkg, values)

	var i int64
	tav, err := types.Eval(token.NewFileSet(), pkg, token.NoPos, expr)
	if err != nil {
		return i, err
	}

	i, b := constant.Int64Val(tav.Value)
	if !b {
		return i, unknownInt64Err
	}

	return i, nil
}

// EvalToUint64 returns the evaluated value(uint64)
func (e *ExpEval) EvalToUint64(expr string, values map[string]interface{}) (uint64, error) {
	pkg := types.NewPackage(e.name, e.path)
	e.setValues(pkg, values)

	var i uint64
	tav, err := types.Eval(token.NewFileSet(), pkg, token.NoPos, expr)
	if err != nil {
		return i, err
	}

	i, b := constant.Uint64Val(tav.Value)
	if !b {
		return i, unknownUint64Err
	}

	return i, nil
}

// EvalToFloat64 returns the evaluated value(float64)
func (e *ExpEval) EvalToFloat64(expr string, values map[string]interface{}) (float64, error) {
	pkg := types.NewPackage(e.name, e.path)
	e.setValues(pkg, values)

	var f float64
	tav, err := types.Eval(token.NewFileSet(), pkg, token.NoPos, expr)
	if err != nil {
		return f, err
	}

	f, b := constant.Float64Val(tav.Value)
	if !b {
		return f, unknownFloat64Err
	}

	return f, nil
}
