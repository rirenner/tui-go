package tui

import (
	"testing"

	"github.com/kr/pretty"
)

func TestList_Draw(t *testing.T) {
	surface := newTestSurface(10, 5)
	painter := NewPainter(surface, NewTheme())

	l := NewList()
	l.AddItems("foo", "bar")
	l.Resize(surface.size)
	l.Draw(painter)

	want := `
foo       
bar       
..........
..........
..........
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}
}

func TestList_RemoveItem(t *testing.T) {
	surface := newTestSurface(5, 3)
	painter := NewPainter(surface, NewTheme())

	l := NewList()
	l.AddItems("one", "two", "three", "four", "five")
	l.SetSelected(1)
	l.Resize(surface.size)
	l.Draw(painter)

	want := `
one  
two  
three
`

	// Make sure okay before removing any items.
	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
		return
	}

	// Remove a visible item.
	l.RemoveItem(2)
	l.Draw(painter)

	want = `
one  
two  
four 
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
		return
	}

	// Remove an item not visible.
	l.RemoveItem(3)
	l.Draw(painter)

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}

	// Selected item should not have changed.
	if l.Selected() != 1 {
		t.Error(pretty.Diff(l.Selected, 1))
	}
}
