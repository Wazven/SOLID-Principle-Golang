package after

type Penyimpanan interface {
	AddItem(item string)
	RemoveItem(item string)
}

type Listable interface {
	ListItems()
}

type Inputable interface {
	InputItem(item string)
}