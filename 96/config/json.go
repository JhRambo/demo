package config

var JsonData = `{
	"configId": 1,
	"spaceId": 1,
	"eid": 1000,
	"data": {
		"node1": {
			"baseInfo": {
				"type": "模型",
				"owner": 1,
				"level": 1,
				"name": "节点名称字符串",
				"description": "描述信息字符串"
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
				"owner": 1,
				"level": 1,
				"name": "节点名称字符串",
				"description": "描述信息字符串"
			},
			"slices":[{
				"a":1,
				"b":2,
				"c":3
			},{
				"a":"string",
				"b":"string",
				"c":"string"
			},{
				"a":[30001,3002],
				"b":"string",
				"c":"string"
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
