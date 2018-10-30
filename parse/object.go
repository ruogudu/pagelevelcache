package parse


type ObjAlias struct {
	Val []byte
}

func NewObject(size int) *ObjAlias {
	return &ObjAlias{Val:make([]byte, size, size)}
}

func (o *ObjAlias) Size() int64 {
	return int64(len(o.Val))
}