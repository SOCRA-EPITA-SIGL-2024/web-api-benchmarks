package main

type position struct {
	lat	float64	`json:"lat"`
	lng	float64 `json:"lng"`
}

type product struct {
	categoryId	string 	`json:"categoryId"`
	name		string 	`json:"name"`
	price		float64	`json:"price"`
}

type garden struct {
	title		string 		`json:"title"`
	id 			string 		`json:"id"`
	position	position 	`json:"position"`
	products	[]product	`json:"products"`
}

var gardens = []garden {
	{
	"title": "Potager 1",
	"id": "u1",
	"position": {
		"lat": 49.252719,
		"lng": 3.720863,
	},
	"products": []product{
		{
			"categoryId": "vegetable",
			"name": "Courgette",
			"price": 0.74,
		},
		{
			"categoryId": "vegetable",
			"name": "Chou",
			"price": 1.95
		},
		{
			"categoryId": "vegetable",
			"name": "Coing",
			"price": 1.74
		},
		{
			"categoryId": "vegetable",
			"name": "Mandarine",
			"price": 0.46
		},
		{
			"categoryId": "vegetable",
			"name": "Cassis",
			"price": 2.97
		},
		{
			"categoryId": "vegetable",
			"name": "Orange",
			"price": 0.13
		},
		{
			"categoryId": "vegetable",
			"name": "Poireau",
			"price": 1.29
		},
		{
			"categoryId": "vegetable",
			"name": "Fenouil",
			"price": 0.33
		},
		{
			"categoryId": "vegetable",
			"name": "Poivron",
			"price": 1.02
		},
		{
			"categoryId": "vegetable",
			"name": "Figue",
			"price": 0.07
		},
		{
			"categoryId": "vegetable",
			"name": "Salsifis",
			"price": 0.39
		},
		{
			"categoryId": "vegetable",
			"name": "Navet",
			"price": 1.94
		},
		{
			"categoryId": "vegetable",
			"name": "Ch\u00e2taigne",
			"price": 1.23
		},
		{
			"categoryId": "vegetable",
			"name": "\u00c9chalotte",
			"price": 0.27
		},
		{
			"categoryId": "vegetable",
			"name": "Choux",
			"price": 2.12
		},
		{
			"categoryId": "vegetable",
			"name": "Aubergine",
			"price": 2.47
		},
		{
			"categoryId": "vegetable",
			"name": "Panais",
			"price": 2.87
		},
		{
			"categoryId": "vegetable",
			"name": "Cerise",
			"price": 2.07
		},
		{
			"categoryId": "vegetable",
			"name": "Brocoli",
			"price": 0.61
		},
		{
			"categoryId": "vegetable",
			"name": "Tomate",
			"price": 1.9
		},
		{
			"categoryId": "vegetable",
			"name": "Kiwi",
			"price": 1.16
		},
		{
			"categoryId": "vegetable",
			"name": "Past\u00e8que",
			"price": 0.5
		},
		{
			"categoryId": "vegetable",
			"name": "Asperge",
			"price": 1.59
		},
		{
			"categoryId": "vegetable",
			"name": "C\u00e9leri",
			"price": 0.93
		},
		{
			"categoryId": "vegetable",
			"name": "Oignon",
			"price": 2.59
		},
		{
			"categoryId": "vegetable",
			"name": "Prune",
			"price": 0.25
		},
		{
			"categoryId": "vegetable",
			"name": "Raisin",
			"price": 1.16
		},
		{
			"categoryId": "vegetable",
			"name": "Carotte",
			"price": 0.93
		},
		{
			"categoryId": "vegetable",
			"name": "Cl\u00e9mentine",
			"price": 0.87
		},
		{
			"categoryId": "vegetable",
			"name": "Champignons de Paris",
			"price": 2.55
		},
		{
			"categoryId": "vegetable",
			"name": "Salade",
			"price": 1.66
		},
		{
			"categoryId": "vegetable",
			"name": "Framboise",
			"price": 1.58
		},
		{
			"categoryId": "vegetable",
			"name": "Cresson",
			"price": 2.51
		},
		{
			"categoryId": "vegetable",
			"name": "Concombre",
			"price": 1.01
		},
		{
			"categoryId": "vegetable",
			"name": "P\u00eache",
			"price": 0.61
		},
		{
			"categoryId": "vegetable",
			"name": "Chou-fleur",
			"price": 1.56
		},
		{
			"categoryId": "vegetable",
			"name": "Topinambour",
			"price": 2.36
		},
		{
			"categoryId": "vegetable",
			"name": "Blette",
			"price": 1.41
		},
		{
			"categoryId": "vegetable",
			"name": "Pamplemousse",
			"price": 0.77
		},
		{
			"categoryId": "vegetable",
			"name": "Radis",
			"price": 1.38
		},
		{
			"categoryId": "vegetable",
			"name": "Mirabelle",
			"price": 1.98
		},
		{
			"categoryId": "vegetable",
			"name": "Champignon de Paris",
			"price": 2.22
		},
		{
			"categoryId": "vegetable",
			"name": "carotte",
			"price": 1.5
		},
		{
			"categoryId": "vegetable",
			"name": "Myrtille",
			"price": 0.42
		},
		{
			"categoryId": "vegetable",
			"name": "Nectarine",
			"price": 0.0
		},
		{
			"categoryId": "vegetable",
			"name": "Choux de Bruxelles",
			"price": 0.34
		},
		{
			"categoryId": "vegetable",
			"name": "Potiron",
			"price": 2.75
		},
		{
			"categoryId": "vegetable",
			"name": "Artichaut",
			"price": 2.25
		},
		{
			"categoryId": "vegetable",
			"name": "M\u00fbre",
			"price": 2.22
		},
		{
			"categoryId": "vegetable",
			"name": "Petit pois",
			"price": 2.03
		},
		{
			"categoryId": "vegetable",
			"name": "Rhubarbe",
			"price": 2.8
		},
		{
			"categoryId": "vegetable",
			"name": "Noix",
			"price": 0.97
		},
		{
			"categoryId": "vegetable",
			"name": "Melon",
			"price": 0.14
		},
		{
			"categoryId": "vegetable",
			"name": "Endive",
			"price": 1.95
		},
		{
			"categoryId": "vegetable",
			"name": "Ail",
			"price": 2.45
		},
		{
			"categoryId": "vegetable",
			"name": "\u00c9pinard",
			"price": 2.04
		},
		{
			"categoryId": "vegetable",
			"name": "Kaki",
			"price": 2.19
		},
		{
			"categoryId": "vegetable",
			"name": "Haricot vert",
			"price": 2.47
		},
		{
			"categoryId": "vegetable",
			"name": "Abricot",
			"price": 1.61
		},
		{
			"categoryId": "vegetable",
			"name": "Fraise",
			"price": 2.49
		},
		{
			"categoryId": "vegetable",
			"name": "Courge",
			"price": 1.77
		},
		{
			"categoryId": "vegetable",
			"name": "Betterave",
			"price": 2.37
		},
		{
			"categoryId": "vegetable",
			"name": "Pomme",
			"price": 1.48
		},
		{
			"categoryId": "vegetable",
			"name": "M\u00e2che",
			"price": 2.37
		},
		{
			"categoryId": "vegetable",
			"name": "Citron",
			"price": 2.7
		},
		{
			"categoryId": "vegetable",
			"name": "\u00c9chalote",
			"price": 0.7
		},
		{
			"categoryId": "vegetable",
			"name": "Poire",
			"price": 1.88
		}
	]
}
}