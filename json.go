package main

import (
	"encoding/json"
	"fmt"
)

const (
	jsonBlob = string(`{
		"_id": "61410596ba0c39a1fae2f47f",
		"index": 0,
		"guid": "fe1102ba-b5f9-4783-a62c-13adb3c2c293",
		"isActive": false,
		"balance": "$1,500.31",
		"picture": "http://placehold.it/32x32",
		"age": 33,
		"eyeColor": "blue",
		"name": "Aguirre Aguilar",
		"gender": "male",
		"company": "COMTRAK",
		"email": "aguirreaguilar@comtrak.com",
		"phone": "+1 (828) 595-3377",
		"address": "425 Cheever Place, Norfolk, New York, 6272",
		"about": "Veniam do incididunt dolor commodo  labore ea sint dolore . Nulla amet  non Lorem aliquip proident laboris ut. Minim sint ullamco eu magna voluptate. Ut esse cupidatat minim est qui irure ex ullamco qui aliquip.\r\n",
		"registered": "2019-05-12T10:09:42 +07:00",
		"latitude": 34.248956,
		"longitude": 124.1861,
		"tags": [
			"adipisicing",
			"cillum",
			"mollit",
			"consectetur",
			"adipisicing",
			"duis",
			"adipisicing"
		],
		"friends": [{
				"id": 0,
				"name": "Susanne Anderson"
			},
			{
				"id": 1,
				"name": "Kara Moon"
			},
			{
				"id": 2,
				"name": "Petty Holden"
			}
		],
		"greeting": "Hello, Aguirre Aguilar! You have 4 unread messages.",
		"favoriteFruit": "strawberry"
	}`)
)

// this will unmarshall into m
func deserialize(js string) {
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonBlob), &m)
	if err != nil {
		panic(err)
	}
	for k, v := range m {
		vType := fmt.Sprintf("%T", v)
		fmt.Println("key: ", k, "value: ", v, "valuetype: ", vType)
	}
}

/*
var v interface{}
if err := json.Unmarshal(data, &v); err != nil {
    // handle error
}
switch v := v.(type) {
case []interface{}:
    // it's an array
case map[string]interface{}:
    // it's an object
default:
    // it's something else
}
*/
