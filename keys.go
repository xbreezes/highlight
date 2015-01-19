package highlight

var keywords = map[string]struct{}{
	"defer":            struct{}{},
	"go":               struct{}{},
	"BEGIN":            struct{}{},
	"END":              struct{}{},
	"False":            struct{}{},
	"Infinity":         struct{}{},
	"NaN":              struct{}{},
	"None":             struct{}{},
	"True":             struct{}{},
	"abstract":         struct{}{},
	"alias":            struct{}{},
	"align_union":      struct{}{},
	"alignof":          struct{}{},
	"and":              struct{}{},
	"append":           struct{}{},
	"as":               struct{}{},
	"asm":              struct{}{},
	"assert":           struct{}{},
	"auto":             struct{}{},
	"axiom":            struct{}{},
	"begin":            struct{}{},
	"bool":             struct{}{},
	"boolean":          struct{}{},
	"break":            struct{}{},
	"byte":             struct{}{},
	"caller":           struct{}{},
	"case":             struct{}{},
	"catch":            struct{}{},
	"char":             struct{}{},
	"class":            struct{}{},
	"concept":          struct{}{},
	"concept_map":      struct{}{},
	"const":            struct{}{},
	"const_cast":       struct{}{},
	"constexpr":        struct{}{},
	"continue":         struct{}{},
	"debugger":         struct{}{},
	"decltype":         struct{}{},
	"def":              struct{}{},
	"default":          struct{}{},
	"defined":          struct{}{},
	"del":              struct{}{},
	"delegate":         struct{}{},
	"delete":           struct{}{},
	"die":              struct{}{},
	"do":               struct{}{},
	"double":           struct{}{},
	"dump":             struct{}{},
	"dynamic_cast":     struct{}{},
	"elif":             struct{}{},
	"else":             struct{}{},
	"elsif":            struct{}{},
	"end":              struct{}{},
	"ensure":           struct{}{},
	"enum":             struct{}{},
	"eval":             struct{}{},
	"except":           struct{}{},
	"exec":             struct{}{},
	"exit":             struct{}{},
	"explicit":         struct{}{},
	"export":           struct{}{},
	"extends":          struct{}{},
	"extern":           struct{}{},
	"false":            struct{}{},
	"final":            struct{}{},
	"finally":          struct{}{},
	"float":            struct{}{},
	"float32":          struct{}{},
	"float64":          struct{}{},
	"for":              struct{}{},
	"foreach":          struct{}{},
	"friend":           struct{}{},
	"from":             struct{}{},
	"func":             struct{}{},
	"function":         struct{}{},
	"generic":          struct{}{},
	"get":              struct{}{},
	"global":           struct{}{},
	"goto":             struct{}{},
	"if":               struct{}{},
	"implements":       struct{}{},
	"import":           struct{}{},
	"in":               struct{}{},
	"inline":           struct{}{},
	"instanceof":       struct{}{},
	"int":              struct{}{},
	"int8":             struct{}{},
	"int16":            struct{}{},
	"int32":            struct{}{},
	"int64":            struct{}{},
	"interface":        struct{}{},
	"is":               struct{}{},
	"lambda":           struct{}{},
	"last":             struct{}{},
	"late_check":       struct{}{},
	"local":            struct{}{},
	"long":             struct{}{},
	"make":             struct{}{},
	"map":              struct{}{},
	"module":           struct{}{},
	"mutable":          struct{}{},
	"my":               struct{}{},
	"namespace":        struct{}{},
	"native":           struct{}{},
	"new":              struct{}{},
	"next":             struct{}{},
	"nil":              struct{}{},
	"no":               struct{}{},
	"nonlocal":         struct{}{},
	"not":              struct{}{},
	"null":             struct{}{},
	"nullptr":          struct{}{},
	"operator":         struct{}{},
	"or":               struct{}{},
	"our":              struct{}{},
	"package":          struct{}{},
	"pass":             struct{}{},
	"print":            struct{}{},
	"private":          struct{}{},
	"property":         struct{}{},
	"protected":        struct{}{},
	"public":           struct{}{},
	"raise":            struct{}{},
	"redo":             struct{}{},
	"register":         struct{}{},
	"reinterpret_cast": struct{}{},
	"require":          struct{}{},
	"rescue":           struct{}{},
	"retry":            struct{}{},
	"return":           struct{}{},
	"self":             struct{}{},
	"set":              struct{}{},
	"short":            struct{}{},
	"signed":           struct{}{},
	"sizeof":           struct{}{},
	"static":           struct{}{},
	"static_assert":    struct{}{},
	"static_cast":      struct{}{},
	"strictfp":         struct{}{},
	"struct":           struct{}{},
	"sub":              struct{}{},
	"super":            struct{}{},
	"switch":           struct{}{},
	"synchronized":     struct{}{},
	"template":         struct{}{},
	"then":             struct{}{},
	"this":             struct{}{},
	"throw":            struct{}{},
	"throws":           struct{}{},
	"transient":        struct{}{},
	"true":             struct{}{},
	"try":              struct{}{},
	"type":             struct{}{},
	"typedef":          struct{}{},
	"typeid":           struct{}{},
	"typename":         struct{}{},
	"typeof":           struct{}{},
	"undef":            struct{}{},
	"undefined":        struct{}{},
	"union":            struct{}{},
	"unless":           struct{}{},
	"unsigned":         struct{}{},
	"until":            struct{}{},
	"use":              struct{}{},
	"using":            struct{}{},
	"var":              struct{}{},
	"virtual":          struct{}{},
	"void":             struct{}{},
	"volatile":         struct{}{},
	"wantarray":        struct{}{},
	"when":             struct{}{},
	"where":            struct{}{},
	"while":            struct{}{},
	"with":             struct{}{},
	"yield":            struct{}{},
}
