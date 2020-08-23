package learn

import (
	"fmt"
	"reflect"
	"testing"
)

type Emploee struct {
	EmploeeId string
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

func (e *Emploee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Emploee{"1", "Mike", 30}
	t.Logf("Name:value(%[1]v),Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("json"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age", e)
}

func CheckType(v interface{}) {
	switch t := reflect.TypeOf(v); t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknow", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func TestWithChannelTime(t *testing.T) {
	withChannelTime()
}
