package protocol

// TODO(sqs): does not support inline DocumentSymbol[] in documentSymbolResult, only supports
// range-based ("tag") symbols.

type SymbolData struct {
	Text   string      `json:"text"`
	Detail string      `json:"detail,omitempty"`
	Kind   SymbolKind  `json:"kind"`
	Tags   []SymbolTag `json:"tags,omitempty"`
}

// RangeSymbolTag is a "tag" property on a range vertex that describes the symbol enclosed by the
// range.
type RangeSymbolTag struct {
	Type string `json:"type"`
	SymbolData
	FullRange *RangeData `json:"fullRange,omitempty"`
}

type RangeBasedDocumentSymbol struct {
	ID       uint64                     `json:"id"`
	Children []RangeBasedDocumentSymbol `json:"children,omitempty"`
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

//go:generate go build -o .bin/stringer golang.org/x/tools/cmd/stringer
//go:generate .bin/stringer -type=SymbolTag

type SymbolTag int

const (
	Deprecated SymbolTag = 1

	// TODO(sqs): these are custom extensions, see https://github.com/microsoft/language-server-protocol/issues/98
	Exported   SymbolTag = 100
	Unexported SymbolTag = 101
)

type Symbol struct {
	Vertex
	SymbolData
	Locations []SymbolLocation `json:"locations,omitempty"`
}

type SymbolLocation struct {
	URI       string     `json:"uri"`
	Range     *RangeData `json:"range,omitempty"`
	FullRange RangeData  `json:"fullRange"`
}

func NewSymbol(id uint64, data SymbolData, locations []SymbolLocation) Symbol {
	return Symbol{
		Vertex: Vertex{
			Element: Element{
				ID:   id,
				Type: ElementVertex,
			},
			Label: VertexSymbol,
		},
		SymbolData: data,
		Locations:  locations,
	}
}

type DocumentSymbolResult struct {
	Vertex
	Result []RangeBasedDocumentSymbol `json:"result"`
}

func NewDocumentSymbolResult(id uint64, result []RangeBasedDocumentSymbol) DocumentSymbolResult {
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

type WorkspaceSymbol struct {
	Edge
	OutV uint64   `json:"outV"`
	InVs []uint64 `json:"inVs"`
}

func NewWorkspaceSymbolEdge(id, outV uint64, inVs []uint64) WorkspaceSymbol {
	return WorkspaceSymbol{
		Edge: Edge{
			Element: Element{
				ID:   id,
				Type: ElementEdge,
			},
			Label: EdgeWorkspaceSymbol,
		},
		OutV: outV,
		InVs: inVs,
	}
}
