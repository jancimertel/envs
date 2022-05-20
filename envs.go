package envs

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const TagName = "envs"

func MustHave(schema interface{}) error {
	val := reflect.ValueOf(schema)

	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	} else {
		return errors.New("unexpected type")
	}

	schemaType := val.Type()
	for i := 0; i < val.NumField(); i++ {
		schemaField := schemaType.Field(i)
		tagValue := schemaField.Tag.Get(TagName)
		fieldName := schemaField.Name

		envName := tagValue
		if envName == "" {
			envName = fieldName
		}

		if passedVal := os.Getenv(envName); passedVal == "" {
			return errors.New(fmt.Sprintf("field %s: empty", fieldName))
		} else {
			elemField := val.Field(i)
			if !elemField.CanSet() {
				return errors.New(fmt.Sprintf("field %s: cannot be set", fieldName))
			}

			var err error
			switch elemField.Kind() {
			case reflect.Bool:
				err = assignBoolOrFail(elemField, passedVal)
			case reflect.String:
				elemField.SetString(passedVal)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				err = assignIntOrFail(elemField, passedVal)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				err = assignUintOrFail(elemField, passedVal)
			case reflect.Float32, reflect.Float64:
				err = assignFloatOrFail(elemField, passedVal)
			default:
				err = errors.New("complex type in schema")
			}

			if err != nil {
				return errors.New(fmt.Sprintf("field %s: %v", fieldName, err))
			}
		}
	}
	return nil
}

func isStruct(i interface{}) bool {
	return reflect.ValueOf(i).Type().Kind() == reflect.Struct
}

func assignBoolOrFail(value reflect.Value, raw string) error {
	parsedValue, err := parseBool(raw)
	if err == nil {
		value.SetBool(parsedValue)
		return nil
	} else {
		return errors.New(fmt.Sprintf("cannot parse bool: %v", err))
	}
}

func assignIntOrFail(value reflect.Value, raw string) error {
	parsedValue, err := strconv.ParseInt(raw, 10, 64)
	if err == nil {
		value.SetInt(parsedValue)
		return nil
	} else {
		return errors.New(fmt.Sprintf("cannot parse int: %v", err))
	}
}

func assignUintOrFail(value reflect.Value, raw string) error {
	parsedValue, err := strconv.ParseUint(raw, 10, 64)
	if err == nil {
		value.SetUint(parsedValue)
		return nil
	} else {
		return errors.New(fmt.Sprintf("cannot parse uint: %v", err))
	}
}

func assignFloatOrFail(value reflect.Value, raw string) error {
	parsedValue, err := strconv.ParseFloat(raw, 10)
	if err == nil {
		value.SetFloat(parsedValue)
		return nil
	} else {
		return errors.New(fmt.Sprintf("cannot parse float: %v", err))
	}
}

func parseBool(in string) (bool, error) {
	trimmed := strings.ToLower(strings.TrimSpace(in))
	if trimmed == "true" || trimmed == "1" {
		return true, nil
	} else if trimmed == "false" || trimmed == "0" {
		return false, nil
	} else {
		return false, errors.New(fmt.Sprintf("cannot parse bool: %v", in))
	}
}
