package config

import (
	"../library/linetcr"
	"../utils"
)

func (self *Kickop) Ceko(pelaku int64) bool {
	if !utils.InArrayInt64(self.Opinv, pelaku) {
		self.Opinv = append(self.Opinv, pelaku)
		return true
	}
	return false
}

func (self *Kickop) Cek(pelaku string) bool {
	if !utils.InArrayString(self.Kick, pelaku) {
		self.Kick = append(self.Kick, pelaku)
		return true
	}
	return false
}

func (self *Kickop) Del(pelaku string) {
	self.Kick = utils.RemoveString(self.Kick, pelaku)
}

func (self *Kickop) Ceki(pelaku string) bool {
	defer linetcr.PanicOnly()
	if !utils.InArrayString(self.Inv, pelaku) {
		self.Inv = append(self.Inv, pelaku)
		return true
	}
	return false
}

func (self *Kickop) Deli(pelaku string) {
	self.Inv = utils.RemoveString(self.Inv, pelaku)
}

func (self *Kickop) Clear() {
	self.Inv = []string{}
	self.Kick = []string{}
	self.Opinv = []int64{}
}
