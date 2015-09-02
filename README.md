# mstf
Microservices test framework

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
      "params":[{
        "name":"XXXX",
        "position":"header",
        "value":"XXXX",
        "from":{
          "case_name":"XXXX"
          "param_name:"XXXX"
        }
      },{...},{...}]
    },
    "response":{
      "params":[{
        "name":"XXXXX"
        "position":"body"
        "value":"XXXX"
      },{...},{...}]
    }
  },{...},{...}]
}
