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

func (this *TouchTestNode) Position() (x, y float32) {
	return this.X, this.Y
}
func (this *TouchTestNode) SetPosition(x, y float32) {
	this.X, this.Y = x, y
}
func (this *TouchTestNode) Size() (w, h float32) {
	return this.W, this.H
}
func (this *TouchTestNode) SetSize(w, h float32) {
	this.W, this.H = w, h
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

type enterLeaveDeliveryTest struct {
	touchPositions []sg.TouchPoint
	itemGeometry   [][4]float32
	enterPoints    [][]sg.TouchPoint
	movePoints     [][]sg.TouchPoint
	leavePoints    [][]sg.TouchPoint
}

// Should not get any events: mouse position stays out of bounds the whole time.
func TestNoEnterLeave(t *testing.T) {
	testData := enterLeaveDeliveryTest{
		touchPositions: []sg.TouchPoint{
			sg.TouchPoint{X: -1, Y: -1}, // top left
			sg.TouchPoint{X: 5, Y: -1},  // top center
			sg.TouchPoint{X: 11, Y: -1}, // top right
			sg.TouchPoint{X: -1, Y: 11}, // bottom left
			sg.TouchPoint{X: 5, Y: 11},  // bottom center
			sg.TouchPoint{X: 11, Y: 11}, // bottom right
			sg.TouchPoint{X: -1, Y: 5},  // left center
			sg.TouchPoint{X: 5, Y: 55},  // right center
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
		movePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
		enterPoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
		leavePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
	}
	touchTestHelper(t, &testData)
}

// Should get a single enter event, mouse enters the position and stays there.
func TestSingleEnterWhenCursorMoves(t *testing.T) {
	testData := enterLeaveDeliveryTest{
		touchPositions: []sg.TouchPoint{
			sg.TouchPoint{X: -1, Y: -1},
			sg.TouchPoint{X: 1, Y: 1},
			sg.TouchPoint{X: 1, Y: 1},
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
		movePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
		leavePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{}, // points stay inside. no leaves.
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
	}
	touchTestHelper(t, &testData)
}

// Should get a single leave event, mouse enters the position and leaves it.
func TestSingleLeaveWhenCursorMoves(t *testing.T) {
	testData := enterLeaveDeliveryTest{
		touchPositions: []sg.TouchPoint{
			sg.TouchPoint{X: -1, Y: -1},
			sg.TouchPoint{X: 1, Y: 1},
			sg.TouchPoint{X: 1, Y: 1},
			sg.TouchPoint{X: -1, Y: -1},
			sg.TouchPoint{X: -1, Y: -1},
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
		movePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
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
	testData := enterLeaveDeliveryTest{
		touchPositions: []sg.TouchPoint{
			sg.TouchPoint{X: 15, Y: 1},
			sg.TouchPoint{X: 15, Y: 1},
			sg.TouchPoint{X: 15, Y: 1},
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
		movePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
		leavePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{}, // point never leaves
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
	}
	touchTestHelper(t, &testData)
}

// If the item size changes to move out from under the pointer, we should get a leave.
func TestLeaveWhenItemSizeChanges(t *testing.T) {
	testData := enterLeaveDeliveryTest{
		touchPositions: []sg.TouchPoint{
			sg.TouchPoint{X: 15, Y: 1},
			sg.TouchPoint{X: 15, Y: 1},
			sg.TouchPoint{X: 15, Y: 1},
			sg.TouchPoint{X: 15, Y: 1},
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
		movePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
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
	testData := enterLeaveDeliveryTest{
		touchPositions: []sg.TouchPoint{
			// taller than wider
			sg.TouchPoint{X: 1, Y: 25},
			sg.TouchPoint{X: 1, Y: 15},

			// wider than taller
			sg.TouchPoint{X: 25, Y: 1},
			sg.TouchPoint{X: 15, Y: 1},
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
		movePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
		},
		leavePoints: [][]sg.TouchPoint{
			[]sg.TouchPoint{},
			[]sg.TouchPoint{},
			[]sg.TouchPoint{sg.TouchPoint{X: 25, Y: 1}},
			[]sg.TouchPoint{},
		},
	}
	touchTestHelper(t, &testData)
}

func touchTestHelper(t *testing.T, testData *enterLeaveDeliveryTest) {
	hn := &TouchTestNode{}
	ih := NewInputHelper()

	if len(testData.touchPositions) != len(testData.enterPoints) ||
		len(testData.touchPositions) != len(testData.leavePoints) ||
		len(testData.touchPositions) != len(testData.movePoints) ||
		len(testData.touchPositions) != len(testData.itemGeometry) {
		t.Fatalf("Invalid form of test data. Input sizes must match output sizes.")
	}

	for idx, _ := range testData.touchPositions {
		ih.MousePos = testData.touchPositions[idx]

		geo := testData.itemGeometry[idx]
		ih.ProcessPointerEvents(geo[0], geo[1], geo[2], geo[3], hn)
		ih.ResetFrameState()

		if len(hn.Enters) != len(testData.enterPoints[idx]) {
			t.Fatalf("Got unexpected enter count: %d, wanted %d", len(hn.Enters), len(testData.enterPoints[idx]))
		}
		if len(hn.Leaves) != len(testData.leavePoints[idx]) {
			t.Fatalf("Got unexpected leave count: %d, wanted %d", len(hn.Leaves), len(testData.leavePoints[idx]))
		}
		if len(hn.Moves) != len(testData.movePoints[idx]) {
			t.Fatalf("Got unexpected move count: %d, wanted %d", len(hn.Moves), len(testData.movePoints[idx]))
		}

		hn.Enters = []sg.TouchPoint{}
		hn.Leaves = []sg.TouchPoint{}
	}
}
