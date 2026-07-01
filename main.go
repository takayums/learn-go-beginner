package main

import (
	"fmt"

	"github.com/takayums/learn-go-beginner/akses"
)

func main() {
	// 01HelloWorld
	fmt.Println("1. Hello Wolrd")
	HelloworldFunc()

	// 02Variable
	fmt.Println("2. Variable")
	VariableFunc()

	// 03TipeData
	fmt.Println("3. Tipe Data")
	TipeDataFunc()

	// 04Constanta
	fmt.Println("4. Konstanta")
	ConsFunc()

	// 05Operator
	fmt.Println("5. Operator")
	OperasiFunc()

	// 06Condition
	fmt.Println("6. Condition")
	ConditionFunc()

	// 07Loop
	fmt.Println("7. Loop")
	LoopFunc()

	// 08Array
	fmt.Println("8. Array")
	ArrayFunc()

	// 09Slice
	fmt.Println("9. Slice")
	SliceFunc()

	// 10Map
	fmt.Println("10. Map")
	MapFunc()

	// 11TypeDec
	fmt.Println("11. Type Declaration")
	TypedocFunc()

	// 12Switch
	fmt.Println("12. Switch")
	SwitchFunc()

	// 13Func
	fmt.Println("13. Function")
	FuncFunc()

	// 14Pointer
	fmt.Println("14. Ponter")
	Pointer()

	// 15Struct
	fmt.Println("15. Struct")
	StructFunc()

	// 16Method
	fmt.Println("16. Method")
	Method()

	// 17AksesLibrary
	fmt.Println("17. Akses Library")
	akses.LevelAkses()

	// 18Interface
	fmt.Println("18. Interface")
	Interface()
}
