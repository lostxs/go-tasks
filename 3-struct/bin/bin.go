package bin

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewBin(id, name string, private bool) *Bin {
	return &Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}

func (bin *Bin) Print() {
	fmt.Printf("ID: %s\nName: %s\n", bin.ID, bin.Name)
}
