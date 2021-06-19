package main

// #cgo pkg-config: python-3.9-embed
// #include <Python.h>
import "C"

func main() {
	//  Pythonインタプリタ終了
	defer C.Py_Finalize()
	// Pythonインタプリタの初期化
	C.Py_Initialize()
	pyCodeStr := `print("Hello, World!")`
	// GoのstringをCのcharに型変換する（変換しないとPyRun_SimpleStringに型が合ってないよって怒られる）
	// cannot use "print(\"Hello, World!\")" (type string) as type *_Ctype_char in argument to _Cfunc_PyRun_SimpleString
	pyCodeChar := C.CString(pyCodeStr)
	// Pythonコードを文字列として受け取ってインタプリタ上で実行
	C.PyRun_SimpleString(pyCodeChar)
}
