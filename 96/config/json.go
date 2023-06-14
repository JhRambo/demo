package config

var JsonData = `{
	"configId": 1,
	"spaceId": 1,
	"eid": 1000,
	"nodeList": [{
			"id": "uuid1",
			"baseInfo": {
				"name": "节点名称1",
				"description": "详细描述信息1"
			},
			"transformInfo": {
				"scale": {
					"x": 1.1,
					"y": 1.1,
					"z": 1.1
				},
				"position": {
					"x": 2.2,
					"y": 2.2,
					"z": 2.2
				},
				"rotation": {
					"x": 3.3,
					"y": 3.3,
					"z": 3.3
				}
			},
			"type": "模型",
			"level": 1
		},
		{
			"id": "uuid2",
			"baseInfo": {
				"name": "节点名称2",
				"description": "详细描述信息2"
			},
			"slices": [{
				"a": 1,
				"b": 2
			}, {
				"a": [30001, 30002],
				"b": "xyz"
			}],
			"type": "造型",
			"level": 2
		}
	],
	"basedata": {
		"light": {
			"yaw": 36.1,
			"pitch": -90.0
		}
	}
}`
