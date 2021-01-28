package protocol

type RangeBasedDocumentSymbol struct {
	ID       uint64                      `json:"id"`
	Children []*RangeBasedDocumentSymbol `json:"children,omitempty"`
}

type DocumentSymbolResult struct {
	Vertex
	Result []*RangeBasedDocumentSymbol `json:"result"` // TODO: lsp.DocumentSymbol
}

func NewDocumentSymbolResult(id uint64, result []*RangeBasedDocumentSymbol) DocumentSymbolResult {
	return DocumentSymbolResult{
		Vertex: Vertex{
			Element: Element{
				ID:   id,
				Type: ElementVertex,
			},
			Label: VertexDocumentSymbolResult,
		},
		Result: result,
	}
}

type DocumentSymbolEdge struct {
	Edge
	InV  uint64 `json:"inV"`
	OutV uint64 `json:"outV"`
}

func NewDocumentSymbolEdge(id, outV, inV uint64) DocumentSymbolEdge {
	return DocumentSymbolEdge{
		Edge: Edge{
			Element: Element{
				ID:   id,
				Type: ElementEdge,
			},
			Label: EdgeTextDocumentDocumentSymbol,
		},
		OutV: outV,
		InV:  inV,
	}
}
