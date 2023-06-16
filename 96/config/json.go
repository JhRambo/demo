package config

var JsonData = `{
	"configId": 1,
	"spaceId": 1,
	"eid": 1000,
	"nodeList": [{
			"id": "uuid1",
			"type": 1,
			"level": 1,
			"baseInfo": {
				"name": "基础信息1",
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
			}
		},
		{
			"id": "uuid2",
			"type": 2,
			"level": 2,
			"baseInfo": {
				"name": "基础信息2",
				"description": "详细描述信息2"
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
			"slices": [{
				"a": 111,
				"b": 111,
				"c": 111
			}, {
				"a": 222,
				"b": 222,
				"c": 222
			}, {
				"a": 333,
				"b": 333,
				"c": 333
			}]
		}
	],
	"basedata": {
		"light": {
			"yaw": 36.1,
			"pitch": -90.0
		},
		"globalMusic": {
			"playType": 1,
			"fileInfoList": []
		},
		"guide": {
			"fileInfoList": []
		},
		"others": {
			"x": 1,
			"y": 2,
			"z": 3
		}
	}
}`
