package argo

import "fmt"

// Option allows to pass custom options
//
// See the option subpackage
type Option map[string]interface{}

func mergeOptions(options ...Option) Option {
	opt := Option{}
	for _, o := range options {
		for k, v := range o {
			opt[k] = v
		}
	}
	return opt
}

// GetID returns the "id" value if present (empty string otherwise)
func (o Option) GetID() string {
	return o.getString("id")
}

func (o Option) getString(key string) string {
	if v, ok := o[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

// GetPosition returns the "position" value if present
func (o Option) GetPosition() (int, bool) {
	return o.getInt("position")
}

func (o Option) getInt(key string) (int, bool) {
	if v, ok := o[key]; ok {
		if s, ok := v.(int); ok {
			return s, true
		}
	}
	return 0, false
}
