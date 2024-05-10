package gotots

func mapTypes(field *Field) {
	switch field.typ {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
		field.typ = "number"
	case "string":
		field.typ = "string"
	case "bool":
		field.typ = "boolean"
	default:
		field.typ = "any"
	}

}
