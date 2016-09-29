package main

import "testing"

func TestNewPerson(t *testing.T) {
	name := "boris"
	p := newPerson(name)
	if p.name != name {
		t.Error("Expected name to be ", name, " but got ", p.name)
	}
}

func TestNewListItem(t *testing.T) {
	p := newPerson("boris")
	i := newListItem(p)

	if i.friend.name != p.name {
		t.Error("Expected name to be", p.name, "but got", i.friend.name)
	}
}

func TestAddToList(t *testing.T) {
	p1 := newPerson("boris")
	i := newListItem(p1)
	p2 := newPerson("spider")

	l := addToList(i, p2)

	if l.friend.name != p2.name {
		t.Fail()
	}

	if l.next.friend.name != p1.name {
		t.Fail()
	}

}

func TestAddToListNil(t *testing.T) {
	i := &list{}
	p := newPerson("boris")

	l := addToList(i, p)

	if l.friend.name != p.name {
		t.Fail()
	}
}

func TestMakeFriends(t *testing.T) {
	p1 := newPerson("boris")
	p2 := newPerson("spider")

	makeFriends(p1, p2)

	if p1.friends.friend.name != p2.name {
		t.Fail()
	}

	if p2.friends.friend.name != p1.name {
		t.Fail()
	}
}

func TestApplyToList(t *testing.T) {
	var pFromF *person
	f := func(p *person) { pFromF = p }
	l := &list{friend: newPerson("harald")}

	applyToList(l, f)

	if pFromF.name != l.friend.name {
		t.Fail()
	}
}
