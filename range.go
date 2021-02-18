package protocol

type Range struct {
	Vertex
	RangeData
	Tag *RangeTag `json:"tag,omitempty"`
}

type RangeData struct {
	Start Pos `json:"start"`
	End   Pos `json:"end"`
}

type RangeTag struct {
	Type      string     `json:"type"`
	FullRange *RangeData `json:"fullRange,omitempty"`
	SymbolData
}

type Pos struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

func NewRange(id uint64, start, end Pos, tag *RangeTag) Range {
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
