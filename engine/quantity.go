package engine

import (
	"fmt"
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
)

// TODO
// Slice() ->  EvalTo(dst)

// OutputQuantity represents a space-dependent quantity,
// that can be saved, like M, B_eff or alpha.
type outputField interface {
	Slice() (q *data.Slice, recycle bool) // get quantity data (GPU or CPU), indicate need to recycle
	NComp() int                           // Number of components (1: scalar, 3: vector, ...)
	Name() string                         // Human-readable identifier, e.g. "m", "alpha"
	Unit() string                         // Unit, e.g. "A/m" or "".
	Mesh() *data.Mesh                     // Usually the global mesh, unless this quantity has a special size.
	average() []float64
}

func VectorFunc(name, unit string, f func(dst *data.Slice)) VectorField {
	return AsVectorField(&callbackOutput{info{3, name, unit}, f})
}

func ScalarFunc(name, unit string, f func(dst *data.Slice)) ScalarField {
	return AsScalarField(&callbackOutput{info{1, name, unit}, f})
}

type callbackOutput struct {
	info
	call func(*data.Slice)
}

func (c *callbackOutput) Mesh() *data.Mesh   { return Mesh() }
func (c *callbackOutput) average() []float64 { return qAverageUniverse(c) }

// Calculates and returns the quantity.
// recycle is true: slice needs to be recycled.
func (q *callbackOutput) Slice() (s *data.Slice, recycle bool) {
	buf := cuda.Buffer(q.NComp(), q.Mesh().Size())
	cuda.Zero(buf)
	q.call(buf)
	return buf, true
}

// ScalarField is a Quantity guaranteed to have 1 component.
// Provides convenience methods particular to scalars.
type ScalarField struct {
	outputField
}

// AsScalarField promotes a quantity to a ScalarField,
// enabling convenience methods particular to scalars.
func AsScalarField(q outputField) ScalarField {
	if q.NComp() != 1 {
		panic(fmt.Errorf("ScalarField(%v): need 1 component, have: %v", q.Name(), q.NComp()))
	}
	return ScalarField{q}
}

func (s ScalarField) Average() float64         { return s.outputField.average()[0] }
func (s ScalarField) Region(r int) ScalarField { return AsScalarField(inRegion(s.outputField, r)) }

// VectorField is a Quantity guaranteed to have 3 components.
// Provides convenience methods particular to vectors.
type VectorField struct {
	outputField
}

// AsVectorField promotes a quantity to a VectorField,
// enabling convenience methods particular to vectors.
func AsVectorField(q outputField) VectorField {
	if q.NComp() != 3 {
		panic(fmt.Errorf("VectorField(%v): need 3 components, have: %v", q.Name(), q.NComp()))
	}
	return VectorField{q}
}

func (v VectorField) Average() data.Vector     { return unslice(v.outputField.average()) }
func (v VectorField) Region(r int) VectorField { return AsVectorField(inRegion(v.outputField, r)) }
func (v VectorField) Comp(c int) ScalarField   { return AsScalarField(Comp(v.outputField, c)) }
