package struct_nostruct

// memory:
// "0"のスタックポインタを変数の値として保存する
//	[
//		...
//		0, value1, value2, ...
//		...
//	]

// define(str):
//	struct_name_1 param_name_1:type1 param_name_2:type2=default_value_1
// struct_table:
//	[struct_name_1, ...]
// param_table:
//	[[param_name_1, param_name_2], ...]
// count_of_param:
//	[2, ...]

func DefineStruct(vm *VM) {}
