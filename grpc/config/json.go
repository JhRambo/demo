package config

var JsonData = `{
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
			},
			"fileInfo": {

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
			"fileInfo": {

			}
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
		}
	}
}`
