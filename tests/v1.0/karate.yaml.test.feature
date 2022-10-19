
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:


Scenario: get request

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"files": [
			{
				"name": "template.tmpl",
				"data": "<h1>{{header}}</h1>"
			}
		],
		"template": "template.tmpl",
		"data": {
			"header": "THIS IS THE HEADER"
		}
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	{
	"mustache": 
	{
		"result": "<h1>THIS IS THE HEADER</h1>",
		"success": true
	}
	}
	"""
	