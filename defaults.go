package defaults

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// ConversionFunc defines a function used for converting a string to a reflect.Kind
type ConversionFunc func(val reflect.Value, dflt string) error

// This is the global table (populated on init) of conversion functions.
var CONVERSIONS = map[reflect.Kind]ConversionFunc{}

func init() {
	CONVERSIONS = map[reflect.Kind]ConversionFunc{
		reflect.Bool: func(val reflect.Value, dflt string) error {
			if !val.Interface().(bool) {
				b, err := strconv.ParseBool(dflt)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(b))
			}
			return nil
		},
		reflect.Int8: func(val reflect.Value, dflt string) error {
			if val.Interface().(int8) == 0 {
				b, err := strconv.ParseInt(dflt, 10, 8)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(int8(b)))
			}
			return nil
		},
		reflect.Int16: func(val reflect.Value, dflt string) error {
			if val.Interface().(int16) == 0 {
				b, err := strconv.ParseInt(dflt, 10, 16)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(int16(b)))
			}
			return nil
		},
		reflect.Int32: func(val reflect.Value, dflt string) error {
			if val.Interface().(int32) == 0 {
				b, err := strconv.ParseInt(dflt, 10, 32)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(int32(b)))
			}
			return nil
		},
		reflect.Int64: func(val reflect.Value, dflt string) error {
			if val.Interface().(int64) == 0 {
				b, err := strconv.ParseInt(dflt, 10, 64)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(int64(b)))
			}
			return nil
		},
		reflect.Int: func(val reflect.Value, dflt string) error {
			if val.Interface().(int) == 0 {
				b, err := strconv.ParseInt(dflt, 10, 32)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(int(b)))
			}
			return nil
		},
		reflect.Uint8: func(val reflect.Value, dflt string) error {
			if val.Interface().(uint8) == 0 {
				b, err := strconv.ParseUint(dflt, 10, 8)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(uint8(b)))
			}
			return nil
		},
		reflect.Uint16: func(val reflect.Value, dflt string) error {
			if val.Interface().(uint16) == 0 {
				b, err := strconv.ParseUint(dflt, 10, 16)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(uint16(b)))
			}
			return nil
		},
		reflect.Uint32: func(val reflect.Value, dflt string) error {
			if val.Interface().(uint32) == 0 {
				b, err := strconv.ParseUint(dflt, 10, 32)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(uint32(b)))
			}
			return nil
		},
		reflect.Uint64: func(val reflect.Value, dflt string) error {
			if val.Interface().(uint64) == 0 {
				b, err := strconv.ParseUint(dflt, 10, 64)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(uint64(b)))
			}
			return nil
		},
		reflect.Uint: func(val reflect.Value, dflt string) error {
			if val.Interface().(uint) == 0 {
				b, err := strconv.ParseUint(dflt, 10, 32)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(uint(b)))
			}
			return nil
		},
		reflect.Uintptr: func(val reflect.Value, dflt string) error {
			if val.Interface().(uintptr) == 0 {
				b, err := strconv.ParseUint(dflt, 10, 32)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(uintptr(b)))
			}
			return nil
		},
		reflect.Float32: func(val reflect.Value, dflt string) error {
			if val.Interface().(float32) == 0.0 {
				b, err := strconv.ParseFloat(dflt, 32)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(float32(b)))
			}
			return nil
		},
		reflect.Float64: func(val reflect.Value, dflt string) error {
			if val.Interface().(float64) == 0.0 {
				b, err := strconv.ParseFloat(dflt, 64)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(float64(b)))
			}
			return nil
		},
		reflect.Complex64: func(val reflect.Value, dflt string) error {
			if val.Interface().(complex64) == 0+0i {
				b, err := strconv.ParseComplex(dflt, 64)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(complex64(b)))
			}
			return nil
		},
		reflect.Complex128: func(val reflect.Value, dflt string) error {
			if val.Interface().(complex128) == 0+0i {
				b, err := strconv.ParseComplex(dflt, 128)
				if err != nil {
					return err
				}
				val.Set(reflect.ValueOf(complex128(b)))
			}
			return nil
		},
		reflect.String: func(val reflect.Value, dflt string) error {
			if val.Interface().(string) == "" {
				val.Set(reflect.ValueOf(dflt))
			}
			return nil
		},
		reflect.Struct: func(val reflect.Value, dflt string) error {
			return defaultValue(val)
		},
	}
}

func defaultValue(this reflect.Value) error {
	if method := this.MethodByName("Default"); method.IsValid() {
		if !this.Elem().CanSet() {
			return errors.New("cannot set this struct")
		}

		ret := method.Call(nil)
		if !ret[0].IsNil() {
			return ret[0].Interface().(error)
		}

		return nil
	}

	if this.Kind() == reflect.Pointer {
		this = this.Elem()
	}

	if this.Kind() == reflect.Struct {
		for i := 0; i < this.NumField(); i++ {
			field := this.Type().Field(i)
			dflt, ok := field.Tag.Lookup("default")
			isStruct := field.Type.Kind() == reflect.Struct || (field.Type.Kind() == reflect.Pointer && field.Type.Elem().Kind() == reflect.Struct)

			if isStruct && ok {
				return fmt.Errorf("defaults are not allowed on structs such as %q", field.Name)
			}

			if (ok && dflt != "") || isStruct {
				valField := this.Field(i)

				if valField.CanSet() {
					if valField.Kind() == reflect.Pointer {
						if valField.IsNil() && field.Type.Elem().Kind() == reflect.Struct {
							valField.Set(reflect.New(field.Type.Elem()))
						}

						valField = valField.Elem()
					}

					if f, ok := CONVERSIONS[valField.Kind()]; ok {
						if err := f(valField, dflt); err != nil {
							return fmt.Errorf("invalid conversion of default type for %q: %w", field.Name, err)
						}
					} else {
						return fmt.Errorf("default provided for %q but it's type cannot be set", field.Name)
					}
				} else {
					return fmt.Errorf("default provided for %q but it cannot be set", field.Name)
				}
			}
		}
	}

	return nil
}

func Default(this any) error {
	return defaultValue(reflect.ValueOf(this))
}
