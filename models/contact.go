package models

import (
	u "ManagerContacts/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"`
}

func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Conatct should be on payload"), false
	}
	if contact.Phone == "" {
		return u.Message(false, "Contact number must be on payload"), false
	}
	if contact.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}
	return u.Message(true, "success"), true
}

func (contact *Contact) Create() map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}
	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("iser_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}
