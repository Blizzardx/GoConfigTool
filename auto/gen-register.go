package auto

import (
	"reflect"
	//"fmt"
	"github.com/Blizzardx/GoConfigTool/configManager"
	"fmt"
)

func init(){
//Codec:= codec.MustGetCodec("msgpack")
	type1 :=  reflect.TypeOf((*WorldPlayerInfo)(nil)).Elem()
	configManager.Register(type1)
	//fmt.Print(type1)

	obj := reflect.New(type1).Interface()
	instance := obj.(*WorldPlayerInfo)
	instance.PlayerId = 1000
	fmt.Println(instance)
}
