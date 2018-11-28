/*
 * Copyright 2017 Crimson AS <info@crimson.no>
 * Author: Robin Burchell <robin.burchell@crimson.no>
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this
 *    list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

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

	r.LayoutChildren(geometryChildren)

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
