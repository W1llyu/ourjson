package ourjson

import (
	"testing"
	"fmt"
)

func TestParseObject(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	jsonStr := `{
		"user": {
			"name": "aa",
			"age": 10,
			"phone": "12222222222",
			"emails": [
				"aa@164.com",
				"aa@165.com"
			],
			"address": [
				{
					"number": "101",
					"now_live": true
				},
				{
					"number": "102",
					"now_live": null
				}
			],
			"account": {
				"balance": 999.9
			}
		}
	}
	`
	jsonObject, err := ParseObject(jsonStr)
	fmt.Println(jsonObject, err)

	user := jsonObject.GetJsonObject("user")
	fmt.Println(user)

	name, err := user.GetString("name")
	fmt.Println(name, err)

	phone, err := user.GetInt64("phone")
	fmt.Println(phone, err)

	age, err := user.GetInt64("age")
	fmt.Println(age, err)

	account := user.GetJsonObject("account")
	fmt.Println(account)

	balance, err := account.GetFloat64("balance")
	fmt.Println(balance, err)

	email1, err := user.GetJsonArray("emails").GetString(0)
	fmt.Println(email1, err)

	address := user.GetJsonArray("address")
	fmt.Println(address)

	address1nowLive, err := user.GetJsonArray("address").GetJsonObject(0).GetBoolean("now_live")
	fmt.Println(address1nowLive, err)

	address2, err := address.Get(1)
	fmt.Println(address2, err)

	address2NowLive, err := address2.JsonObject().GetNullBoolean("now_live")
	fmt.Println(address2NowLive, err)
}

func TestParseArray(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	jsonStr := `[
		{
			"name": "Will"
		},
		{
			"name": "Uzi"
		}
	]
	`
	jsonArray, err := ParseArray(jsonStr)
	fmt.Println(jsonArray, err)

	name1, err := jsonArray.GetJsonObject(0).GetString("name")
	fmt.Println(name1, err)
}
