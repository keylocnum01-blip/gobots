package config

func RemoveSticker(items []*Stickers, item *Stickers) []*Stickers {
	newitems := []*Stickers{}
	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}
	return newitems
}
