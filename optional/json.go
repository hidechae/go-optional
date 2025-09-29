package optional

import (
	"encoding/json"
)

func (o *Option[T]) UnmarshalJSON(b []byte) error {
	var v *T
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	if v != nil {
		o.value = []T{*v}
	} else {
		o.value = nil
	}
	return nil
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToPtr())
}
