
# mustache 1.0

Mustache templating

---
- #### Categories: misc
- #### Image: gcr.io/direktiv/functions/mustache 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/mustache/issues
- #### URL: https://github.com/direktiv-apps/mustache
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About mustache

This function requires a [Mustache](https://mustache.github.io/) template and data. The function returns the rendered template with the data provided. It runs Mustache version 1.1.1.

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: mustache
  image: gcr.io/direktiv/functions/mustache:1.0
  type: knative-workflow
```
   #### Direct template
```yaml
- id: mustache 
  type: action
  action:
    function: mustache
    input: 
      files:
      - name: template.tmpl
        data: <h1>{{header}}</h1>
      template: template.tmpl
      data:
        header: This is the header
```
   #### Template from Direktiv variable
```yaml
- id: mustache 
  type: action
  action:
    function: mustache
    files:
    - key: templ.tpl
      scope: workflow
    input: 
      template: templ.tpl
      data:
        header: This is the header
```

   ### Secrets


*No secrets required*







### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": "\u003ch1\u003eTHIS IS THE HEADER\u003c/h1\u003e",
    "success": true
  },
  {
    "result": "My Template Output",
    "success": true
  }
]
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| mustache | [PostOKBodyMustache](#post-o-k-body-mustache)| `PostOKBodyMustache` |  | |  |  |


#### <span id="post-o-k-body-mustache"></span> postOKBodyMustache

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | string| `string` | ✓ | | Template output |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [interface{}](#interface)| `interface{}` |  | |  |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |
| template | string| `string` |  | | Mustache template to use |  |

 
