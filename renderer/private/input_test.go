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

package private

import (
	"testing"

	"github.com/CrimsonAS/goggle/sg"
)

type TouchTestNode struct {
	X, Y, W, H float32
	Enters     []sg.Vec2
	Leaves     []sg.Vec2
	Moves      []sg.Vec2
}

func (this *TouchTestNode) PointerEnter(tp sg.Vec2) {
	this.Enters = append(this.Enters, tp)
}
func (this *TouchTestNode) PointerLeave(tp sg.Vec2) {
	this.Leaves = append(this.Enters, tp)
}
func (this *TouchTestNode) PointerMoved(tp sg.Vec2) {
	this.Moves = append(this.Moves, tp)
}

func TestTouchTestNodeInterface(t *testing.T) {
	var hn sg.Node = &TouchTestNode{}
	if _, ok := hn.(sg.Hoverable); !ok {
		t.Fatalf("TouchTestNode does not implement sg.Hoverable")
	}
	if _, ok := hn.(sg.Moveable); !ok {
		t.Fatalf("TouchTestNode does not implement sg.Moveable")
	}
}

type touchState struct {
	touchPoint sg.Vec2
	buttonDown bool
	buttonUp   bool
}

type touchDeliveryTest struct {
	touchStates  []touchState
	itemGeometry [][4]float32
	enterPoints  [][]sg.Vec2
	movePoints   [][]sg.Vec2
	leavePoints  [][]sg.Vec2
}

// Should not get any events: mouse position stays out of bounds the whole time.
func TestNoEnterLeave(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.Vec2{X: -1, Y: -1}}, // top left
			{touchPoint: sg.Vec2{X: 5, Y: -1}},  // top center
			{touchPoint: sg.Vec2{X: 11, Y: -1}}, // top right
			{touchPoint: sg.Vec2{X: -1, Y: 11}}, // bottom left
			{touchPoint: sg.Vec2{X: 5, Y: 11}},  // bottom center
			{touchPoint: sg.Vec2{X: 11, Y: 11}}, // bottom right
			{touchPoint: sg.Vec2{X: -1, Y: 5}},  // left center
			{touchPoint: sg.Vec2{X: 5, Y: 55}},  // right center
		},
		itemGeometry: [][4]float32{
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
		},
		movePoints:  [][]sg.Vec2{},
		enterPoints: [][]sg.Vec2{},
		leavePoints: [][]sg.Vec2{},
	}
	touchTestHelper(t, &testData)
}

// Should get a single enter event, mouse enters the position and stays there.
func TestSingleEnterWhenCursorMoves(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.Vec2{X: -1, Y: -1}},
			{touchPoint: sg.Vec2{X: 1, Y: 1}},
			{touchPoint: sg.Vec2{X: 1, Y: 1}},
		},
		itemGeometry: [][4]float32{
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
		},
		enterPoints: [][]sg.Vec2{
			[]sg.Vec2{},                    // initial touch outside: no enter
			[]sg.Vec2{sg.Vec2{X: 1, Y: 1}}, // touch inside: enter
			[]sg.Vec2{},
		},
		movePoints:  [][]sg.Vec2{},
		leavePoints: [][]sg.Vec2{},
	}
	touchTestHelper(t, &testData)
}

// Should get a single leave event, mouse enters the position and leaves it.
func TestSingleLeaveWhenCursorMoves(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.Vec2{X: -1, Y: -1}},
			{touchPoint: sg.Vec2{X: 1, Y: 1}},
			{touchPoint: sg.Vec2{X: 1, Y: 1}},
			{touchPoint: sg.Vec2{X: -1, Y: -1}},
			{touchPoint: sg.Vec2{X: -1, Y: -1}},
		},
		itemGeometry: [][4]float32{
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
		},
		enterPoints: [][]sg.Vec2{
			[]sg.Vec2{},                    // initial touch outside: no enter
			[]sg.Vec2{sg.Vec2{X: 1, Y: 1}}, // touch inside: enter
			[]sg.Vec2{},                    // stationary
			[]sg.Vec2{},
			[]sg.Vec2{},
		},
		movePoints: [][]sg.Vec2{},
		leavePoints: [][]sg.Vec2{
			[]sg.Vec2{},
			[]sg.Vec2{},
			[]sg.Vec2{},
			[]sg.Vec2{sg.Vec2{X: -1, Y: -1}}, // point leaves
			[]sg.Vec2{},                      // point already left; no second leave
		},
	}
	touchTestHelper(t, &testData)
}

// If the item size changes to move under the pointer, we should get an enter.
func TestEnterWhenItemSizeChanges(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.Vec2{X: 15, Y: 1}},
			{touchPoint: sg.Vec2{X: 15, Y: 1}},
			{touchPoint: sg.Vec2{X: 15, Y: 1}},
		},
		itemGeometry: [][4]float32{
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 15, 10},
			[4]float32{0, 0, 15, 10},
		},
		enterPoints: [][]sg.Vec2{
			[]sg.Vec2{},                     // initial touch outside: no enter
			[]sg.Vec2{sg.Vec2{X: 15, Y: 1}}, // touch inside: enter
			[]sg.Vec2{},                     // no additional enter
		},
		movePoints:  [][]sg.Vec2{},
		leavePoints: [][]sg.Vec2{},
	}
	touchTestHelper(t, &testData)
}

// If the item size changes to move out from under the pointer, we should get a leave.
func TestLeaveWhenItemSizeChanges(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.Vec2{X: 15, Y: 1}},
			{touchPoint: sg.Vec2{X: 15, Y: 1}},
			{touchPoint: sg.Vec2{X: 15, Y: 1}},
			{touchPoint: sg.Vec2{X: 15, Y: 1}},
		},
		itemGeometry: [][4]float32{
			[4]float32{0, 0, 20, 10},
			[4]float32{0, 0, 20, 10},
			[4]float32{0, 0, 5, 10},
			[4]float32{0, 0, 5, 10},
		},
		enterPoints: [][]sg.Vec2{
			[]sg.Vec2{sg.Vec2{X: 15, Y: 1}}, // touch inside: enter
			[]sg.Vec2{},                     // no further enters
			[]sg.Vec2{},
			[]sg.Vec2{},
		},
		movePoints: [][]sg.Vec2{},
		leavePoints: [][]sg.Vec2{
			[]sg.Vec2{},
			[]sg.Vec2{},
			[]sg.Vec2{sg.Vec2{X: 15, Y: 1}},
			[]sg.Vec2{}, // no further leaves
		},
	}
	touchTestHelper(t, &testData)
}

// Make sure that item dimensions don't affect point contains testing
// That is, that a point is always inside, no matter which dimension is larger.
func TestTallerThanWider(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			// taller than wider
			{touchPoint: sg.Vec2{X: 1, Y: 25}},
			{touchPoint: sg.Vec2{X: 1, Y: 15}},

			// wider than taller
			{touchPoint: sg.Vec2{X: 25, Y: 1}},
			{touchPoint: sg.Vec2{X: 15, Y: 1}},
		},
		itemGeometry: [][4]float32{
			// taller than wider
			[4]float32{0, 0, 10, 20},
			[4]float32{0, 0, 10, 20},

			// wider than taller
			[4]float32{0, 0, 20, 10},
			[4]float32{0, 0, 20, 10},
		},
		enterPoints: [][]sg.Vec2{
			// taller than wider
			[]sg.Vec2{},                     // start outside
			[]sg.Vec2{sg.Vec2{X: 1, Y: 15}}, // move inside

			// wider than taller
			[]sg.Vec2{},                     // start outside
			[]sg.Vec2{sg.Vec2{X: 15, Y: 1}}, // move inside
		},
		movePoints: [][]sg.Vec2{},
		leavePoints: [][]sg.Vec2{
			[]sg.Vec2{},
			[]sg.Vec2{},
			[]sg.Vec2{sg.Vec2{X: 25, Y: 1}},
			[]sg.Vec2{},
		},
	}
	touchTestHelper(t, &testData)
}

// Items not at the origin should get points delivered in item coordinates, not
// scene coordinates.
func TestItemNotAtOrigin(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.Vec2{X: 15, Y: 25}},
			{touchPoint: sg.Vec2{X: 25, Y: 35}},
		},
		itemGeometry: [][4]float32{
			[4]float32{5, 5, 20, 20},
			[4]float32{5, 5, 20, 20},
		},
		enterPoints: [][]sg.Vec2{
			[]sg.Vec2{sg.Vec2{X: 10, Y: 20}},
			[]sg.Vec2{},
		},
		movePoints: [][]sg.Vec2{},
		leavePoints: [][]sg.Vec2{
			[]sg.Vec2{},
			[]sg.Vec2{sg.Vec2{X: 20, Y: 30}},
		},
	}
	touchTestHelper(t, &testData)
}

func touchTestHelper(t *testing.T, testData *touchDeliveryTest) {
	hn := &TouchTestNode{}
	ih := NewInputHelper()

	if (len(testData.enterPoints) != 0 && len(testData.touchStates) != len(testData.enterPoints)) ||
		(len(testData.leavePoints) != 0 && len(testData.touchStates) != len(testData.leavePoints)) ||
		(len(testData.movePoints) != 0 && len(testData.touchStates) != len(testData.movePoints)) ||
		(len(testData.touchStates) != len(testData.itemGeometry)) { // these are mandatory
		t.Fatalf("Invalid form of test data. Input sizes must match output sizes.")
	}

	for idx, _ := range testData.touchStates {
		ih.MousePos = testData.touchStates[idx].touchPoint
		ih.ButtonDown = testData.touchStates[idx].buttonDown
		ih.ButtonUp = testData.touchStates[idx].buttonUp

		geo := testData.itemGeometry[idx]
		ih.ProcessPointerEvents(sg.Vec2{geo[0], geo[1]}, geo[2], geo[3], hn)
		ih.ResetFrameState()

		expectedEnters := []sg.Vec2{}
		expectedLeaves := []sg.Vec2{}
		expectedMoves := []sg.Vec2{}
		if len(testData.enterPoints) != 0 {
			expectedEnters = testData.enterPoints[idx]
		}
		if len(testData.leavePoints) != 0 {
			expectedLeaves = testData.leavePoints[idx]
		}
		if len(testData.movePoints) != 0 {
			expectedMoves = testData.movePoints[idx]
		}
		if len(hn.Enters) != len(expectedEnters) {
			t.Fatalf("Got unexpected enter count: %d, wanted %d", len(hn.Enters), len(expectedEnters))
		}
		if len(hn.Leaves) != len(expectedLeaves) {
			t.Fatalf("Got unexpected leave count: %d, wanted %d", len(hn.Leaves), len(expectedLeaves))
		}
		if len(hn.Moves) != len(expectedMoves) {
			t.Fatalf("Got unexpected move count: %d, wanted %d", len(hn.Moves), len(expectedMoves))
		}

		for idx, _ := range hn.Enters {
			if hn.Enters[idx] != expectedEnters[idx] {
				t.Fatalf("Expected enter %s at index %d, got %s instead", expectedEnters[idx], idx, hn.Enters[idx])
			}
		}
		for idx, _ := range hn.Leaves {
			if hn.Leaves[idx] != expectedLeaves[idx] {
				t.Fatalf("Expected leave %s at index %d, got %s instead", expectedLeaves[idx], idx, hn.Leaves[idx])
			}
		}
		for idx, _ := range hn.Moves {
			if hn.Moves[idx] != expectedMoves[idx] {
				t.Fatalf("Expected move %s at index %d, got %s instead", expectedMoves[idx], idx, hn.Moves[idx])
			}
		}

		hn.Enters = []sg.Vec2{}
		hn.Leaves = []sg.Vec2{}
		hn.Moves = []sg.Vec2{}
	}
}
