package private

import (
	"testing"

	"github.com/CrimsonAS/goggle/sg"
)

type TouchTestNode struct {
	X, Y, W, H float32
	Enters     []sg.TouchPoint
	Leaves     []sg.TouchPoint
	Moves      []sg.TouchPoint
}

func (this *TouchTestNode) PointerEnter(tp sg.TouchPoint) {
	this.Enters = append(this.Enters, tp)
}
func (this *TouchTestNode) PointerLeave(tp sg.TouchPoint) {
	this.Leaves = append(this.Enters, tp)
}
func (this *TouchTestNode) PointerMoved(tp sg.TouchPoint) {
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
	touchPoint sg.TouchPoint
	buttonDown bool
	buttonUp   bool
}

type touchDeliveryTest struct {
	touchStates  []touchState
	itemGeometry [][4]float32
	enterPoints  [][]sg.TouchPoint
	movePoints   [][]sg.TouchPoint
	leavePoints  [][]sg.TouchPoint
}

// Should not get any events: mouse position stays out of bounds the whole time.
func TestNoEnterLeave(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.TouchPoint{X: -1, Y: -1}}, // top left
			{touchPoint: sg.TouchPoint{X: 5, Y: -1}},  // top center
			{touchPoint: sg.TouchPoint{X: 11, Y: -1}}, // top right
			{touchPoint: sg.TouchPoint{X: -1, Y: 11}}, // bottom left
			{touchPoint: sg.TouchPoint{X: 5, Y: 11}},  // bottom center
			{touchPoint: sg.TouchPoint{X: 11, Y: 11}}, // bottom right
			{touchPoint: sg.TouchPoint{X: -1, Y: 5}},  // left center
			{touchPoint: sg.TouchPoint{X: 5, Y: 55}},  // right center
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
		movePoints:  [][]sg.TouchPoint{},
		enterPoints: [][]sg.TouchPoint{},
		leavePoints: [][]sg.TouchPoint{},
	}
	touchTestHelper(t, &testData)
}

// Should get a single enter event, mouse enters the position and stays there.
func TestSingleEnterWhenCursorMoves(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.TouchPoint{X: -1, Y: -1}},
			{touchPoint: sg.TouchPoint{X: 1, Y: 1}},
			{touchPoint: sg.TouchPoint{X: 1, Y: 1}},
		},
		itemGeometry: [][4]float32{
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
		},
		enterPoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},                          // initial touch outside: no enter
			[]sg.TouchPoint{sg.TouchPoint{X: 1, Y: 1}}, // touch inside: enter
			[]sg.TouchPoint{},
		},
		movePoints:  [][]sg.TouchPoint{},
		leavePoints: [][]sg.TouchPoint{},
	}
	touchTestHelper(t, &testData)
}

// Should get a single leave event, mouse enters the position and leaves it.
func TestSingleLeaveWhenCursorMoves(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.TouchPoint{X: -1, Y: -1}},
			{touchPoint: sg.TouchPoint{X: 1, Y: 1}},
			{touchPoint: sg.TouchPoint{X: 1, Y: 1}},
			{touchPoint: sg.TouchPoint{X: -1, Y: -1}},
			{touchPoint: sg.TouchPoint{X: -1, Y: -1}},
		},
		itemGeometry: [][4]float32{
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 10, 10},
		},
		enterPoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},                          // initial touch outside: no enter
			[]sg.TouchPoint{sg.TouchPoint{X: 1, Y: 1}}, // touch inside: enter
			[]sg.TouchPoint{},                          // stationary
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
		movePoints: [][]sg.TouchPoint{},
		leavePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{sg.TouchPoint{X: -1, Y: -1}}, // point leaves
			[]sg.TouchPoint{},                            // point already left; no second leave
		},
	}
	touchTestHelper(t, &testData)
}

// If the item size changes to move under the pointer, we should get an enter.
func TestEnterWhenItemSizeChanges(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.TouchPoint{X: 15, Y: 1}},
			{touchPoint: sg.TouchPoint{X: 15, Y: 1}},
			{touchPoint: sg.TouchPoint{X: 15, Y: 1}},
		},
		itemGeometry: [][4]float32{
			[4]float32{0, 0, 10, 10},
			[4]float32{0, 0, 15, 10},
			[4]float32{0, 0, 15, 10},
		},
		enterPoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},                           // initial touch outside: no enter
			[]sg.TouchPoint{sg.TouchPoint{X: 15, Y: 1}}, // touch inside: enter
			[]sg.TouchPoint{},                           // no additional enter
		},
		movePoints:  [][]sg.TouchPoint{},
		leavePoints: [][]sg.TouchPoint{},
	}
	touchTestHelper(t, &testData)
}

// If the item size changes to move out from under the pointer, we should get a leave.
func TestLeaveWhenItemSizeChanges(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.TouchPoint{X: 15, Y: 1}},
			{touchPoint: sg.TouchPoint{X: 15, Y: 1}},
			{touchPoint: sg.TouchPoint{X: 15, Y: 1}},
			{touchPoint: sg.TouchPoint{X: 15, Y: 1}},
		},
		itemGeometry: [][4]float32{
			[4]float32{0, 0, 20, 10},
			[4]float32{0, 0, 20, 10},
			[4]float32{0, 0, 5, 10},
			[4]float32{0, 0, 5, 10},
		},
		enterPoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{sg.TouchPoint{X: 15, Y: 1}}, // touch inside: enter
			[]sg.TouchPoint{},                           // no further enters
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
		movePoints: [][]sg.TouchPoint{},
		leavePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{sg.TouchPoint{X: 15, Y: 1}},
			[]sg.TouchPoint{}, // no further leaves
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
			{touchPoint: sg.TouchPoint{X: 1, Y: 25}},
			{touchPoint: sg.TouchPoint{X: 1, Y: 15}},

			// wider than taller
			{touchPoint: sg.TouchPoint{X: 25, Y: 1}},
			{touchPoint: sg.TouchPoint{X: 15, Y: 1}},
		},
		itemGeometry: [][4]float32{
			// taller than wider
			[4]float32{0, 0, 10, 20},
			[4]float32{0, 0, 10, 20},

			// wider than taller
			[4]float32{0, 0, 20, 10},
			[4]float32{0, 0, 20, 10},
		},
		enterPoints: [][]sg.TouchPoint{
			// taller than wider
			[]sg.TouchPoint{},                           // start outside
			[]sg.TouchPoint{sg.TouchPoint{X: 1, Y: 15}}, // move inside

			// wider than taller
			[]sg.TouchPoint{},                           // start outside
			[]sg.TouchPoint{sg.TouchPoint{X: 15, Y: 1}}, // move inside
		},
		movePoints: [][]sg.TouchPoint{},
		leavePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{sg.TouchPoint{X: 25, Y: 1}},
			[]sg.TouchPoint{},
		},
	}
	touchTestHelper(t, &testData)
}

// Items not at the origin should get points delivered in item coordinates, not
// scene coordinates.
func TestItemNotAtOrigin(t *testing.T) {
	testData := touchDeliveryTest{
		touchStates: []touchState{
			{touchPoint: sg.TouchPoint{X: 15, Y: 25}},
			{touchPoint: sg.TouchPoint{X: 25, Y: 35}},
		},
		itemGeometry: [][4]float32{
			[4]float32{5, 5, 20, 20},
			[4]float32{5, 5, 20, 20},
		},
		enterPoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{sg.TouchPoint{X: 10, Y: 20}},
			[]sg.TouchPoint{},
		},
		movePoints: [][]sg.TouchPoint{},
		leavePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{sg.TouchPoint{X: 20, Y: 30}},
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

		expectedEnters := []sg.TouchPoint{}
		expectedLeaves := []sg.TouchPoint{}
		expectedMoves := []sg.TouchPoint{}
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

		hn.Enters = []sg.TouchPoint{}
		hn.Leaves = []sg.TouchPoint{}
		hn.Moves = []sg.TouchPoint{}
	}
}
