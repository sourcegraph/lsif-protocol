package protocol

// RangeSymbolTag is a "tag" property on a range vertex that describes the symbol enclosed by the
// range.
type RangeSymbolTag struct {
	Type      string     `json:"type"`
	Text      string     `json:"text"`
	Kind      string     `json:"kind"`
	FullRange *RangeData `json:"fullRange,omitempty"`
	Detail    string     `json:"detail,omitempty"`
}

type RangeBasedDocumentSymbol struct {
	ID       int64                      `json:"id"`
	Children []RangeBasedDocumentSymbol `json:"children,omitempty"`
}

type DocumentSymbol struct {
	Name           string           `json:"name"`
	Detail         string           `json:"detail,omitempty"`
	Kind           SymbolKind       `json:"kind"`
	Tags           []SymbolTag      `json:"tags,omitempty"`
	Range          RangeData        `json:"range"`
	SelectionRange RangeData        `json:"selectionRange"`
	Children       []DocumentSymbol `json:"children,omitempty"`
}

type SymbolKind int

const (
	File          SymbolKind = 1
	Module        SymbolKind = 2
	Namespace     SymbolKind = 3
	Package       SymbolKind = 4
	Class         SymbolKind = 5
	Method        SymbolKind = 6
	Property      SymbolKind = 7
	Field         SymbolKind = 8
	Constructor   SymbolKind = 9
	Enum          SymbolKind = 10
	Interface     SymbolKind = 11
	Function      SymbolKind = 12
	Variable      SymbolKind = 13
	Constant      SymbolKind = 14
	String        SymbolKind = 15
	Number        SymbolKind = 16
	Boolean       SymbolKind = 17
	Array         SymbolKind = 18
	Object        SymbolKind = 19
	Key           SymbolKind = 20
	Null          SymbolKind = 21
	EnumMember    SymbolKind = 22
	Struct        SymbolKind = 23
	Event         SymbolKind = 24
	Operator      SymbolKind = 25
	TypeParameter SymbolKind = 26
)

type SymbolTag int

const (
	Deprecated SymbolTag = 1
)

type DocumentSymbolResult struct {
	Vertex
	// TODO(sqs): make type-safe (either []RangeBasedDocumentSymbol or []DocumentSymbol)
	Result interface{} `json:"result"`
}

func NewDocumentSymbolResult(id uint64, result interface{}) DocumentSymbolResult {
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

type TextDocumentDocumentSymbol struct {
	Edge
	OutV uint64 `json:"outV"`
	InV  uint64 `json:"inV"`
}

func NewTextDocumentDocumentSymbolEdge(id, outV, inV uint64) TextDocumentDocumentSymbol {
	return TextDocumentDocumentSymbol{
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
