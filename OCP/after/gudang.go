package after

import "fmt"

type Gudang struct {
	Items []string
}

func (g *Gudang) AddItem(item string) {
	g.Items = append(g.Items, item)
}

func (g *Gudang) RemoveItem(item string) {
	for i, j := range g.Items {
		if j == item {
			g.Items = append(g.Items[:i], g.Items[i+1:]...)
			break
		}
	}
}

func (g *Gudang) ListItems() {
	fmt.Println("Barang di Gudang: ")
	for _, item := range g.Items{
		fmt.Println(item)
	}
}

func (g *Gudang) InputItem(Item string){
	g.AddItem(Item)
}