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