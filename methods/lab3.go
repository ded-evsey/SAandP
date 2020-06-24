package main

import (
	"encoding/json"
	"io/ioutil"
)
// Задать структуру данных для хранения электронных адресов включая
//		ФИО,
//		email,
//		url страницы в одной из социальных сетей,
//		возраст владельца.
// Реализовать методы:
//		добавления записей об электронных адресах
//		удаления записей об электронных адресах,
//		поиска по социальным сетям,
//		поиска по фамилии,
//		загрузки данных в файл,
//		выгрузки данных в файл.
type Account struct {
	FirstName, MiddleName, LastName, Email, Url string
	Age int
}
type Accounts []Account
func (accounts Accounts) SearchMiddleName(middleName string) int  {
	for id, curAcc := range accounts{
		if curAcc.MiddleName == middleName {
			return id
		}
	}
	return -1
}

func (accounts Accounts) SearchUrl(url string) int {
	for id, curAcc := range accounts{
		if curAcc.Url == url {
			return id
		}
	}
	return -1
}
func (accounts Accounts)  GetId(str string) int{
	if 0 < accounts.SearchUrl(str){
		return accounts.SearchUrl(str)
	} else {
		return accounts.SearchMiddleName(str)
	}
}
func (accounts Accounts) Delete(str string) Accounts {
	id := accounts.GetId(str)
	if  id < 0{
		return accounts
	}
	return append(accounts[:id], accounts[id:] ...)
}
func (accounts Accounts) AddNew(firstName, middleName, lastName, email, url string, age int) Accounts {
	newAccount := Account{
		FirstName:   firstName,
		MiddleName: middleName,
		LastName: lastName,
		Email: email,
		Url:   url,
		Age:   age,
	}
	return append(accounts, newAccount)
}

func (accounts Accounts) WriteFile(filename string)  {
	jsonData, _ := json.Marshal(accounts)
	_ = ioutil.WriteFile(filename, jsonData, 0777)
}
func (accounts Accounts) ReadFile(filename string)  {
	data, _ := ioutil.ReadFile(filename)
	var f Accounts
	_ =json.Unmarshal(data, &f)
	for _, data := range f{
		accounts.AddNew(
			data.FirstName,
			data.MiddleName,
			data.LastName,
			data.Email,
			data.Url,
			data.Age)
	}
}
