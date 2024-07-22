// Code generated by gtrace. DO NOT EDIT.

package trace

// coordinationComposeOptions is a holder of options
type coordinationComposeOptions struct {
	panicCallback func(e interface{})
}

// CoordinationOption specified Coordination compose option
type CoordinationComposeOption func(o *coordinationComposeOptions)

// WithCoordinationPanicCallback specified behavior on panic
func WithCoordinationPanicCallback(cb func(e interface{})) CoordinationComposeOption {
	return func(o *coordinationComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new Coordination which has functional fields composed both from t and x.
func (t *Coordination) Compose(x *Coordination, opts ...CoordinationComposeOption) *Coordination {
	var ret Coordination
	return &ret
}
