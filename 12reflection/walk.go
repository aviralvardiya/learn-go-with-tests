package reflection

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}

func Walk(x interface{}, fn func(input string)) {
	value := getValue(x)

	// // if(value.Kind()==reflect.Pointer){
	// // 	value=value.Elem()
	// // }

	// // if(value.Kind()==reflect.Slice){
	// // 	for i:=0;i<value.Len();i++{
	// // 		Walk(value.Index(i).Interface(),fn)
	// // 	}
	// // 	return
	// // }

	// // for i := 0; i < value.NumField(); i++ {
	// // 	field := value.Field(i)
	// // 	// if(field.Kind()==reflect.String){
	// // 	// 	fn(field.String())
	// // 	// }
	// // 	// if(field.Kind()==reflect.Struct){
	// // 	// 	Walk(field.Interface(),fn)
	// // 	// }

	// // 	switch field.Kind() {
	// // 	case reflect.String:
	// // 		fn(field.String())
	// // 	case reflect.Struct:
	// // 		Walk(field.Interface(), fn)
	// // 	}
	// // }

	walkValue:= func(value reflect.Value){
		Walk(value.Interface(),fn)
	}

	switch value.Kind(){
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i:=0;i<value.NumField();i++{
			walkValue(value.Field(i))
		}
	case reflect.Slice,reflect.Array:
		for i:=0;i<value.Len();i++{
			walkValue(value.Index(i))
		}
	case reflect.Map:
		for _,key:=range value.MapKeys(){
			walkValue(value.MapIndex(key))
		}
	case reflect.Chan:
		for v,ok:=value.Recv();ok;v,ok=value.Recv(){
			walkValue(v)
		}
	case reflect.Func:
		valFuncResult:=value.Call(nil)
		for _,res:=range valFuncResult{
			walkValue(res)
		}
	}

	// numberOfValues := 0
	// var getField func(int) reflect.Value

	// switch value.Kind() {
	// case reflect.String:
	// 	fn(value.String())
	// case reflect.Struct:
	// 	numberOfValues = value.NumField()
	// 	getField = value.Field
	// case reflect.Slice,reflect.Array:
	// 	numberOfValues = value.Len()
	// 	getField = value.Index
	// }

	// for i:=0;i<numberOfValues;i++{
	// 	Walk(getField(i).Interface(),fn)
	// }

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
