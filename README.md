​# mstf
Microservices test framework
mstf only support json in PUT, POST and DELETE request and only parse json in all response.

Config file

format is json format.

{

  "protocols":"http",
  
  "url":"127.0.0.1",
  
  "case":[{
  
    "name":"ms1",
    "request:{
      "url":"/ms1",
      "method":"GET",
      "header":[{
        "name":"XXXX",
        "value":"XXXX",
        "from":{
          "case_name":"XXXX"
          "param_name:"XXXX"
        }
      },{...},{...}
      }],
      "params":[{
        "name":"XXXX",
        "value":"XXXX",
        "from":{
          "case_name":"XXXX"
          "param_name:"XXXX"
        }
      },{...},{...}]
    },
    "response":{
      "status":200,
      "body":[{
        "name":"XXXXX"
        "position":"body"
        "value":"XXXX"
      },{...},{...}]
    }
  },{...},{...}]
}

