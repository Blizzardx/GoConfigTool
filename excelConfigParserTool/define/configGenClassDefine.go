package define

type TemplateInfo struct {
	PackageName       string
	ImportPackageList []*ImportPkgInfo
	ClassList         []*ClassInfo
}
type ImportPkgInfo struct {
	ImportPackage string
}
type ClassInfo struct {
	ClassName string
	FieldList []*FieldInfo
}
type FieldInfo struct {
	FieldType  string
	FieldName  string
	FieldIndex int16
	IsList     bool
	IsMap      bool
	MapKeyType string
}
