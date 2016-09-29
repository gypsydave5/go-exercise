package main

import "fmt"

type person struct {
	name    string
	friends *list
}

type list struct {
	friend *person
	next   *list
}

func newListItem(p *person) *list {
	return &list{
		friend: p,
	}
}

func newPerson(n string) *person {
	return &person{
		name: n,
	}
}

func addToList(l *list, p *person) *list {
	n := newListItem(p)

	if l != nil {
		n.next = l
	}

	return n
}

func listFriends(p *person) {
	applyToList(p.friends, func(p *person) { fmt.Print(p.name, " ") })
	fmt.Print("\n")
}

func applyToList(l *list, f func(p *person)) {
	for l != nil {
		f(l.friend)
		l = l.next
	}
}

func makeFriends(p1 *person, p2 *person) {
	p1.friends = addToList(p1.friends, p2)
	p2.friends = addToList(p2.friends, p1)
}

func main() {
	a := newPerson("Aaron")
	b := newPerson("Basil")
	c := newPerson("Charlie")
	d := newPerson("David")
	e := newPerson("Edmund")

	makeFriends(a, b)
	makeFriends(a, c)
	makeFriends(a, d)
	makeFriends(a, e)

	listFriends(a)
}
