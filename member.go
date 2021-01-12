package protocol

type Member struct {
	Edge
	OutV uint64   `json:"outV"`
	InVs []uint64 `json:"inVs"`
}

func NewMember(id, outV uint64, inVs []uint64) Member {
	return Member{
		Edge: Edge{
			Element: Element{
				ID:   id,
				Type: ElementEdge,
			},
			Label: EdgeMember,
		},
		OutV: outV,
		InVs: inVs,
	}
}
