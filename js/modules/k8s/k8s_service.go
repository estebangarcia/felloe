package k8s

import (
	"encoding/json"
	"github.com/dop251/goja"
	v12 "k8s.io/api/core/v1"
)

type Service struct {
	*v12.Service
}

func serviceConstructor(rt *goja.Runtime, mod *goja.Object) func(call goja.ConstructorCall) *goja.Object {
	return func(call goja.ConstructorCall) *goja.Object {
		return rt.ToValue(NewService(call.Arguments...)).ToObject(rt)
	}
}

func NewService(args ...goja.Value) *v12.Service {
	return newService()
}

func (s *Service) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Service) String() string {
	return s.String()
}