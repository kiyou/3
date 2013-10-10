package script

// Here be dragons

import "reflect"

type ScalarFunction interface {
	Expr
	Float() float64
	Cnst() bool
}

// converts float64 to ScalarFunction
type scalFn struct{ in Expr }

func (c *scalFn) Eval() interface{}  { return c }
func (c *scalFn) Type() reflect.Type { return ScalarFunction_t }
func (c *scalFn) Float() float64     { return c.in.Eval().(float64) }
func (c *scalFn) Cnst() bool         { return Cnst(c.in) }
func (c *scalFn) Child() []Expr      { return []Expr{c.in} }

type VectorFunction interface {
	Expr
	Float3() [3]float64
	Cnst() bool
}

// converts [3]float64 to VectorFunction
type vecFn struct{ in Expr }

func (c *vecFn) Eval() interface{}  { return c }
func (c *vecFn) Type() reflect.Type { return VectorFunction_t }
func (c *vecFn) Float3() [3]float64 { return c.in.Eval().([3]float64) }
func (c *vecFn) Cnst() bool         { return Cnst(c.in) } // always false, need more precise implementation
func (c *vecFn) Child() []Expr      { return []Expr{c.in} }
