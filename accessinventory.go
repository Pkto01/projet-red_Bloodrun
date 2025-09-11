package projet_red_Bloodrun

import "fmt"

func accessInventory(inventory []string) {
	if len(inventory) == 0 {
		fmt.Println("L'inventaire est vide !")
		return
	}

	for i, item := range inventory {
		fmt.Printf("%d: %s\n", i+1, item)
	}
}
