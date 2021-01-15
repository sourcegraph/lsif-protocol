package reader

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	protocol "github.com/sourcegraph/lsif-protocol"
)

func TestUnmarshalElement(t *testing.T) {
	element, err := unmarshalElement(NewInterner(), []byte(`{"id": "47", "type": "vertex", "label": "test"}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling element data: %s", err)
	}

	expectedElement := Element{
		ID:    47,
		Type:  "vertex",
		Label: "test",
	}
	if diff := cmp.Diff(expectedElement, element); diff != "" {
		t.Errorf("unexpected element (-want +got):\n%s", diff)
	}
}

func TestUnmarshalElementNumericIDs(t *testing.T) {
	element, err := unmarshalElement(NewInterner(), []byte(`{"id": 47, "type": "vertex", "label": "test"}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling element data: %s", err)
	}

	expectedElement := Element{
		ID:    47,
		Type:  "vertex",
		Label: "test",
	}
	if diff := cmp.Diff(expectedElement, element); diff != "" {
		t.Errorf("unexpected element (-want +got):\n%s", diff)
	}
}

func TestUnmarshalEdge(t *testing.T) {
	edge, err := unmarshalEdge(NewInterner(), []byte(`{"id": "35", "type": "edge", "label": "item", "outV": "12", "inVs": ["07"], "document": "03"}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling meta data: %s", err)
	}

	expectedEdge := Edge{
		OutV:     12,
		InV:      0,
		InVs:     []int{7},
		Document: 3,
	}
	if diff := cmp.Diff(expectedEdge, edge); diff != "" {
		t.Errorf("unexpected edge (-want +got):\n%s", diff)
	}
}

func TestUnmarshalEdgeNumericIDs(t *testing.T) {
	edge, err := unmarshalEdge(NewInterner(), []byte(`{"id": 35, "type": "edge", "label": "item", "outV": 12, "inVs": [7], "document": 3}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling meta data: %s", err)
	}

	expectedEdge := Edge{
		OutV:     12,
		InV:      0,
		InVs:     []int{7},
		Document: 3,
	}
	if diff := cmp.Diff(expectedEdge, edge); diff != "" {
		t.Errorf("unexpected edge (-want +got):\n%s", diff)
	}
}

func TestUnmarshalMetaData(t *testing.T) {
	metadata, err := unmarshalMetaData(NewInterner(), []byte(`{"id": "01", "type": "vertex", "label": "metaData", "version": "0.4.3", "projectRoot": "file:///test"}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling meta data: %s", err)
	}

	expectedMetadata := MetaData{
		Version:     "0.4.3",
		ProjectRoot: "file:///test",
	}
	if diff := cmp.Diff(expectedMetadata, metadata); diff != "" {
		t.Errorf("unexpected metadata (-want +got):\n%s", diff)
	}
}

func TestUnmarshalDocument(t *testing.T) {
	uri, err := unmarshalDocument(NewInterner(), []byte(`{"id": "02", "type": "vertex", "label": "document", "uri": "file:///test/root/foo.go"}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling document data: %s", err)
	}

	if diff := cmp.Diff("file:///test/root/foo.go", uri); diff != "" {
		t.Errorf("unexpected uri (-want +got):\n%s", diff)
	}
}

func TestUnmarshalRange(t *testing.T) {
	r, err := unmarshalRange(nil, []byte(`{"id": "04", "type": "vertex", "label": "range", "start": {"line": 1, "character": 2}, "end": {"line": 3, "character": 4}, "tag": {"type": "definition", "text": "foo", "kind": 11, "fullRange": {"start": {"line": 1, "character": 0}, "end": {"line": 3, "character": 7}}, "detail": "bar"}}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling range data: %s", err)
	}

	expectedRange := Range{
		RangeData: protocol.RangeData{
			Start: protocol.Pos{Line: 1, Character: 2},
			End:   protocol.Pos{Line: 3, Character: 4},
		},
		Tag: &protocol.RangeSymbolTag{
			Type: "definition",
			SymbolData: protocol.SymbolData{
				Text:   "foo",
				Detail: "bar",
				Kind:   11,
			},
			FullRange: &protocol.RangeData{
				Start: protocol.Pos{Line: 1, Character: 0},
				End:   protocol.Pos{Line: 3, Character: 7},
			},
		},
	}
	if diff := cmp.Diff(expectedRange, r); diff != "" {
		t.Errorf("unexpected range (-want +got):\n%s", diff)
	}
}

func TestUnmarshalHover(t *testing.T) {
	testCases := []struct {
		contents      string
		expectedHover string
	}{
		{
			contents:      `"text"`,
			expectedHover: "text",
		},
		{
			contents:      `[{"kind": "markdown", "value": "text"}]`,
			expectedHover: "text",
		},
		{
			contents:      `[{"language": "go", "value": "text"}]`,
			expectedHover: "```go\ntext\n```",
		},
		{
			contents:      `[{"language": "go", "value": "text"}, {"language": "python", "value": "pext"}]`,
			expectedHover: "```go\ntext\n```\n\n---\n\n```python\npext\n```",
		},
	}

	for _, testCase := range testCases {
		name := fmt.Sprintf("contents=%s", testCase.contents)

		t.Run(name, func(t *testing.T) {
			hover, err := unmarshalHover(NewInterner(), []byte(fmt.Sprintf(`{"id": "16", "type": "vertex", "label": "hoverResult", "result": {"contents": %s}}`, testCase.contents)))
			if err != nil {
				t.Fatalf("unexpected error unmarshalling hover data: %s", err)
			}

			if diff := cmp.Diff(testCase.expectedHover, hover); diff != "" {
				t.Errorf("unexpected hover text (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUnmarshalMoniker(t *testing.T) {
	moniker, err := unmarshalMoniker(NewInterner(), []byte(`{"id": "18", "type": "vertex", "label": "moniker", "kind": "import", "scheme": "scheme A", "identifier": "ident A"}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling moniker data: %s", err)
	}

	expectedMoniker := Moniker{
		Kind:       "import",
		Scheme:     "scheme A",
		Identifier: "ident A",
	}
	if diff := cmp.Diff(expectedMoniker, moniker); diff != "" {
		t.Errorf("unexpected moniker (-want +got):\n%s", diff)
	}
}

func TestUnmarshalPackageInformation(t *testing.T) {
	packageInformation, err := unmarshalPackageInformation(NewInterner(), []byte(`{"id": "22", "type": "vertex", "label": "packageInformation", "name": "pkg A", "version": "v0.1.0"}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling package information data: %s", err)
	}

	expectedPackageInformation := PackageInformation{
		Name:    "pkg A",
		Version: "v0.1.0",
	}
	if diff := cmp.Diff(expectedPackageInformation, packageInformation); diff != "" {
		t.Errorf("unexpected package information (-want +got):\n%s", diff)
	}
}

func TestUnmarshalDiagnosticResult(t *testing.T) {
	diagnosticResult, err := unmarshalDiagnosticResult(NewInterner(), []byte(`{"id": 18, "type": "vertex", "label": "diagnosticResult", "result": [{"severity": 1, "code": 2322, "source": "eslint", "message": "Type '10' is not assignable to type 'string'.", "range": {"start": {"line": 1, "character": 5}, "end": {"line": 1, "character": 6}}}]}`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling diagnostic result data: %s", err)
	}

	expectedDiagnosticResult := []Diagnostic{
		{
			Severity:       1,
			Code:           "2322",
			Message:        "Type '10' is not assignable to type 'string'.",
			Source:         "eslint",
			StartLine:      1,
			StartCharacter: 5,
			EndLine:        1,
			EndCharacter:   6,
		},
	}
	if diff := cmp.Diff(expectedDiagnosticResult, diagnosticResult); diff != "" {
		t.Errorf("unexpected diagnostic result (-want +got):\n%s", diff)
	}
}

func TestUnmarshalDocumentSymbolResult(t *testing.T) {
	interner := NewInterner()

	t.Run("range-based", func(t *testing.T) {
		result, err := UnmarshalDocumentSymbolResult(interner, []byte(`{"id": 39, "type": "vertex", "label": "documentSymbolResult", "result": [{"id": 7, "children": [{"id": "12"}]}, {"id": 8}]}`))
		if err != nil {
			t.Fatalf("unexpected error unmarshalling document symbol result data: %s", err)
		}

		expected := []protocol.RangeBasedDocumentSymbol{
			{ID: 7, Children: []protocol.RangeBasedDocumentSymbol{{ID: 12}}},
			{ID: 8},
		}
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Errorf("unexpected document symbol result (-want +got):\n%s", diff)
		}
	})
}

func TestUnmarshalSymbol(t *testing.T) {
	symbol, err := UnmarshalSymbol(NewInterner(), []byte(`{
 "id": 35,
 "type": "vertex",
 "label": "symbol",
 "text": "writer",
 "detail": "github.com\/sourcegraph\/lsif-protocol\/writer",
 "kind": 4,
 "tags": [
  100
 ],
 "locations": [
  {
   "uri": "file:\/\/\/test",
   "range": {
    "start": {
     "line": 0,
     "character": 8
    },
    "end": {
     "line": 0,
     "character": 14
    }
   },
   "fullRange": {
    "start": {
     "line": 0,
     "character": 0
    },
    "end": {
     "line": 81,
     "character": 2
    }
   }
  }
 ]
}
`))
	if err != nil {
		t.Fatalf("unexpected error unmarshalling symbol: %s", err)
	}

	expectedSymbol := Symbol{
		Text:   "writer",
		Detail: "github.com/sourcegraph/lsif-protocol/writer",
		Kind:   4,
		Tags:   []int{100},
		Locations: []Location{
			{
				URI: "file:///test",
				Range: &protocol.RangeData{
					Start: protocol.Pos{
						Line:      0,
						Character: 8,
					},
					End: protocol.Pos{
						Line:      0,
						Character: 14,
					},
				},
				FullRange: protocol.RangeData{
					Start: protocol.Pos{
						Line:      0,
						Character: 0,
					},
					End: protocol.Pos{
						Line:      81,
						Character: 2,
					},
				},
			},
		},
	}
	if diff := cmp.Diff(expectedSymbol, symbol); diff != "" {
		t.Errorf("unexpected symbol (-want +got):\n%s", diff)
	}
}
