package before

import "fmt"

type Gudang struct {
	items []string
}

func (g *Gudang) AddItem(item string) {
	g.items = append(g.items, item)
}

func (g *Gudang) RemoveItem(item string) {
	for i, it := range g.items {
		if it == item {
			g.items = append(g.items[:i], g.items[i+1:]...)
			return
		}
	}
}

func (g *Gudang) ListItems() {
	fmt.Println("Barang Di Gudang:")
	for _, item := range g.items{
		fmt.Println(item)
	}
}

type Logger struct{}

func (l *Logger) Log(pesan string) {
	fmt.Println("Pesan: " , pesan)
}

// func main() {
// 	gudangmain := gudang{}
// 	gudangmain.AddItem("Barang 1")
// 	gudangmain.AddItem("Barang 2")
// 	gudangmain.AddItem("Barang 3")
// 	gudangmain.ListItems()
// }