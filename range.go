package protocol

type Range struct {
	Vertex
	RangeData
	Tag *RangeDeclarationTag `json:"tag,omitempty"`
}

type RangeData struct {
	Start Pos `json:"start"`
	End   Pos `json:"end"`
}

// Formerly known as RangeSymbolTag
// TODO: rename RangeTag
type RangeDeclarationTag struct {
	Type      string     `json:"type"`
	FullRange *RangeData `json:"fullRange,omitempty"`
	SymbolData
}

type Pos struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

func NewRange(id uint64, start, end Pos, tag *RangeDeclarationTag) Range {
	return Range{
		Vertex: Vertex{
			Element: Element{
				ID:   id,
				Type: ElementVertex,
			},
			Label: VertexRange,
		},
		RangeData: RangeData{
			Start: start,
			End:   end,
		},
		Tag: tag,
	}
}
