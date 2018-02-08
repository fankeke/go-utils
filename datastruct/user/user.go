package main

import (
        "strconv"
       )

type User struct {
    name    string
    age     int
    sex     string
    phone   string
}

func (self *User)SetName(name string) {
   self.name = name
}

func (self *User)GetName()string {
    return self.name
}

func (self *User)SetAge(age int) {
    self.age = age
}

func (self *User)GetAge()int{
    return self.age
}

func (self *User)SetSex(sex string) {
    self.sex = sex
}

func (self *User)GetSex()string {
    return self.sex
}

func (self *User) SetPhone(phone string) {
    self.phone = phone
}

func (self *User) GetPhone() string {
    return self.phone
}

func(self *User) GenInfo() string{
    return "name is " + self.name + ", age is " + strconv.Itoa(self.age) + 
                ", sex=" + self.sex + ", phone=" + self.phone
}



