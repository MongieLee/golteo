package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
	"log"
	"reflect"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		return err
	}
	return
}

func GenerateSnowId() int64 {
	if node == nil {
		err := Init("2024-01-01", 1)
		if err != nil {
			log.Fatalln("snow init failed")
		}
	}
	return node.Generate().Int64()
}

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	s := reflect.ValueOf(obj).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		result[typeOfT.Field(i).Name] = field.Interface()
	}
	return result
}

const _constSalt = "go-pure"

func GenerateRandomSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(append(salt, []byte(_constSalt)...)), nil
}

func HashPassword(password, salt string) (string, error) {
	combined := fmt.Sprintf("%v%v", password, salt)
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(combined), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
}
