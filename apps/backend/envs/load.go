package envs

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Load() {
	t := reflect.TypeOf(serverEnvs{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envName := field.Tag.Get("env")
		switch field.Type.Kind() {
		case reflect.String:
			reflect.ValueOf(&envs).Elem().Field(i).SetString(getEnvStr(envName))
		case reflect.Int:
			reflect.ValueOf(&envs).Elem().Field(i).SetInt(int64(getEnvInt(envName)))
		default:
			panic(fmt.Sprintf("Unsupported type %s", field.Type.Kind()))
		}
		fmt.Println("Loaded ->", envName)
	}
}
func getEnvStr(envName string) string {
	osEnv := os.Getenv(envName)
	if osEnv == "" {
		panic(fmt.Sprintf("Environment variable %s not set", envName))
	}
	return osEnv
}

func getEnvInt(envName string) int {
	osEnv := getEnvStr(envName)
	intEnv, err := strconv.Atoi(osEnv)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s is not an integer", envName))
	}
	return intEnv
}
