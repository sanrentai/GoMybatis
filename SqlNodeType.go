package GoMybatis

type SqlNodeType int

const (
	NArg    SqlNodeType = iota
	NString             //string 节点
	NIf
	NTrim
	NForEach
	NChoose
	NOtherwise
	NWhen
)

func (it SqlNodeType) ToString() string {
	switch it {
	case NString:
		return "NString"
	case NIf:
		return "NIf"
	case NTrim:
		return "NTrim"
	case NForEach:
		return "NForEach"
	case NChoose:
		return "NChoose"
	case NOtherwise:
		return "NOtherwise"
	case NWhen:
		return "NWhen"
	}
	return "Unknow"
}
