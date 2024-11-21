package env

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	l "github.com/peterszarvas94/goat/logger"
)

func Load(variables interface{}, keys ...string) error {
	v := reflect.ValueOf(variables)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		err := fmt.Errorf("Load expects a pointer to a struct, received %v", v.Kind())
		l.Logger.Error(err.Error())
		return err
	}

	structValue := v.Elem()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		if field.Kind() != reflect.String {
			err := fmt.Errorf("Field %s must be of type string", fieldType.Name)
			l.Logger.Error(err.Error())
			return err
		}

		envVarName := strings.ToUpper(fieldType.Name)
		envVarValue, found := os.LookupEnv(envVarName)

		if !found {
			err := fmt.Errorf("Environment variable %s not found", envVarName)
			l.Logger.Error(err.Error())
			return err
		}

		field.SetString(envVarValue)
	}

	l.Logger.Debug("Envs are loaded")

	return nil
}
