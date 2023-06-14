package config

var JsonData = `{
	"configId": 1,
	"spaceId": 1,
	"eid": 1000,
	"nodeList": {
		"node1": {
			"baseInfo": {
				"type": "模型",
				"level": 1
			},
			"transformInfo": {
				"scale": {
					"x": 1.1,
					"y": 1.1,
					"z": 1.1
				},
				"position": {
					"x": 1.1,
					"y": 1.1,
					"z": 1.1
				}
			},
			"edit": true,
			"moveTotal": true
		},
		"node2": {
			"baseInfo": {
				"type": "导航台",
				"level": 1
			},
			"slices":[{
				"a":1,
				"b":2
			},{
				"a":"string",
				"b":"string"
			},{
				"a":[30001,30002],
				"b":"string"
			}]
		}
	},
	"basedata": {
		"light": {
			"yaw": 36.1,
			"pitch": -90.0
		}
	}
}`
