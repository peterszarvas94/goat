package env

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/peterszarvas94/goat/pkg/logger"

	_ "github.com/joho/godotenv/autoload"
)

func Load(variables any, keys ...string) error {
	v := reflect.ValueOf(variables)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		err := fmt.Errorf("Load expects a pointer to a struct, received %v", v.Kind())
		logger.Error(err.Error())
		return err
	}

	structValue := v.Elem()

	for i := range structValue.NumField() {
		field := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		if field.Kind() != reflect.String {
			err := fmt.Errorf("Field %s must be of type string", fieldType.Name)
			logger.Error(err.Error())
			return err
		}

		envVarName := strings.ToUpper(fieldType.Name)
		envVarValue, found := os.LookupEnv(envVarName)

		if !found {
			err := fmt.Errorf("Environment variable %s not found", envVarName)
			logger.Error(err.Error())
			return err
		}

		field.SetString(envVarValue)
	}

	logger.Debug("Envs are loaded")

	return nil
}
