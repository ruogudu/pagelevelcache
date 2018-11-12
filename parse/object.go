package parse


type ObjAlias struct {
	size int64
}

func NewObject(size int) *ObjAlias {
	return &ObjAlias{size: int64(size)}
}

func (o *ObjAlias) Size() int64 {
	return o.size
}

func NewNilOBject() *ObjAlias {
	return &ObjAlias{size: -1}
}

func IsNilObject(v interface{}) bool {
	o, ok := v.(*ObjAlias)

	if ok {
		return o.size < 0
	}

	return false
}