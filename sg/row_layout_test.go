package sg

import "testing"

type rowPositionerTest struct {
	geometries         []*Vec4
	expectedGeometries []*Vec4
}

func TestRowPositioning(t *testing.T) {
	rt := rowPositionerTest{
		geometries: []*Vec4{
			&Vec4{1234, 5678, 101, 201},
			&Vec4{1234, 5678, 102, 202},
		},
		expectedGeometries: []*Vec4{
			&Vec4{0, 5678, 101, 201},
			&Vec4{101, 5678, 102, 202},
		},
	}
	runRowTest(t, rt)
}

func TestRowPositioningWithNilChild(t *testing.T) {
	rt := rowPositionerTest{
		geometries: []*Vec4{
			&Vec4{1234, 5678, 101, 201},
			nil,
			&Vec4{1234, 5678, 102, 202},
		},
		expectedGeometries: []*Vec4{
			&Vec4{0, 5678, 101, 201},
			nil,
			&Vec4{101, 5678, 102, 202},
		},
	}
	runRowTest(t, rt)
}

func runRowTest(t *testing.T, testData rowPositionerTest) {
	if len(testData.geometries) != len(testData.expectedGeometries) {
		t.Fatalf("Bad test data")
	}

	geometryChildren := []Geometryable{}
	nodeChildren := []Node{}
	for _, geom := range testData.geometries {
		if geom == nil {
			geometryChildren = append(geometryChildren, nil)
			nodeChildren = append(nodeChildren, nil)
		} else {
			r := &RectangleNode{
				X:      geom.X,
				Y:      geom.Y,
				Width:  geom.Z,
				Height: geom.W,
			}
			geometryChildren = append(geometryChildren, r)
			nodeChildren = append(nodeChildren, r)
		}
	}
	r := Row{
		Children: nodeChildren,
	}

	r.PositionChildren(geometryChildren)

	for idx, expectedGeometry := range testData.expectedGeometries {
		if expectedGeometry == nil {
			continue
		}
		expectedPosition := Vec2{expectedGeometry.X, expectedGeometry.Y}
		expectedSize := Vec2{expectedGeometry.Z, expectedGeometry.W}
		if geometryChildren[idx].Position() != expectedPosition {
			t.Fatalf("Child %d was put in position %s, expected %s", idx, geometryChildren[idx].Position(), expectedPosition)
		}
		if geometryChildren[idx].Size() != expectedSize {
			t.Fatalf("Child %d was sized %s, expected %s", idx, geometryChildren[idx].Size(), expectedSize)
		}
	}
}
