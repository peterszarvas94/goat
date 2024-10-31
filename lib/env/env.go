package env

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func Load(variables interface{}, keys ...string) error {
	v := reflect.ValueOf(variables)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Load expects a pointer to a struct, received %v", v.Kind())
	}

	structValue := v.Elem()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		if field.Kind() != reflect.String {
			return fmt.Errorf("Field %s must be of type string", fieldType.Name)
		}

		envVarName := strings.ToUpper(fieldType.Name)
		envVarValue, found := os.LookupEnv(envVarName)

		if !found {
			return fmt.Errorf("Environment variable %s not found", envVarName)
		}

		field.SetString(envVarValue)
	}

	return nil
}
