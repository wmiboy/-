// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "易轩软件 http://www.99v66.com",
            "url": "http://www.99v66.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/apply/add": {
            "post": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "应用管理"
                ],
                "summary": "添加应用",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "收费模式 1=扣时 2=扣点",
                        "name": "mold",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "绑定模式 1=不绑定 2=绑定机器码 3=绑定IP",
                        "name": "bin",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "登录模式 1=免费 2=单点登录 3=多点登录",
                        "name": "login",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "远程数据",
                        "name": "msg",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"添加成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/apply/edit": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "应用管理"
                ],
                "summary": "编辑应用信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "aid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用名称",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "收费模式 1=扣时 2=扣点",
                        "name": "mold",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "绑定模式 1=不绑定 2=绑定机器码 3=绑定IP",
                        "name": "bin",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "登录模式 1=免费 2=单点登录 3=多点登录",
                        "name": "login",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "远程数据",
                        "name": "msg",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "1=可用 2=不可用",
                        "name": "state",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/apply/get": {
            "post": {
                "description": "数据已加密处理",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "应用管理"
                ],
                "summary": "获取应用信息[数据已加密]",
                "parameters": [
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "aid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"查询成功\",\"data\":{\"aid\":1002,\"name\":\"QQ注册1\",\"mold\":2,\"login\":1,\"bin\":3,\"msg\":\"关键参数2333\",\"state\":2,\"time\":1625000360}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/apply/getAll": {
            "get": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "应用管理"
                ],
                "summary": "获取应用列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "获取的数量",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "aid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用状态1=可用 2=不可用",
                        "name": "state",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用名称",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"查询成功\",\"count\":2,\"data\":[{\"aid\":1001,\"name\":\"JD注册\",\"mold\":1,\"login\":3,\"bin\":2,\"msg\":\"关键参数\",\"state\":1,\"time\":1625000292}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/key/cread": {
            "post": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "卡密管理"
                ],
                "summary": "生成卡密",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "生成数量",
                        "name": "num",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "收费模式 1=扣时 2=扣点",
                        "name": "mold",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "点数",
                        "name": "point",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "小时数",
                        "name": "day",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "aid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"生成成功\",\"data\":[{\"kid\":2001,\"cdk\":\"acff954d-9160-4cc3-96c3-e9284543a68b\",\"mold\":1,\"point\":100,\"day\":24,\"uid\":0,\"aid\":1002,\"cread_time\":1625006496,\"use_time\":0,\"state\":1}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/key/edit": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "卡密管理"
                ],
                "summary": "编辑卡密信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "卡密ID",
                        "name": "kid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "1=可用 2=不可用",
                        "name": "state",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"修改成功\"}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/key/getAll": {
            "get": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "卡密管理"
                ],
                "summary": "获取卡密列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "获取的数量",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "aid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "应用状态1=可用 2=不可用",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "卡密",
                        "name": "cdk",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"查询成功\",\"count\":1004,\"data\":[{\"kid\":1005,\"cdk\":\"3c262067-384b-4ba5-80ab-14ca31cbadf5\",\"mold\":1,\"point\":100,\"day\":24,\"uid\":0,\"aid\":1002,\"cread_time\":1625006471,\"use_time\":0,\"state\":1}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/power/charge": {
            "post": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "充值",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "卡密",
                        "name": "cdk",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "aid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"充值成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/power/check": {
            "post": {
                "description": "数据已加密",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "验证相关"
                ],
                "summary": "心跳[数据已加密]",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "机器码",
                        "name": "dev",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"效验成功\",\"data\":{\"a\":{\"aid\":3,\"name\":\"JD注册\",\"mold\":2,\"login\":3,\"bin\":1,\"msg\":\"关键参数222\",\"state\":1,\"time\":1625069421},\"p\":{\"pid\":3,\"uid\":1,\"aid\":3,\"kid\":0,\"day\":1625183400,\"point\":90,\"ip\":\"192.168.1.7\",\"code\":\"ca5df0fdae677510f937c256a7f5\",\"state\":1,\"time\":1625084599}}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/power/edit": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限管理"
                ],
                "summary": "编辑权限信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "aid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "1=可用 2=不可用",
                        "name": "state",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "到期时间的十位时间戳",
                        "name": "day",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "剩余点数",
                        "name": "point",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "绑定IP",
                        "name": "ip",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "绑定机器",
                        "name": "code",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/power/getAll": {
            "get": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限管理"
                ],
                "summary": "获取权限列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "获取的数量",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "aid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "状态1=可用 2=不可用",
                        "name": "state",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"查询成功\",\"count\":1,\"data\":[{\"uid\":1001,\"aid\":1002,\"kid\":0,\"day\":1625191303,\"point\":200,\"ip\":\"\",\"code\":\"\",\"state\":1,\"time\":1625018503}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/power/login": {
            "post": {
                "description": "数据已加密",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "验证相关"
                ],
                "summary": "用户登录[数据已加密]",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "aid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "机器码",
                        "name": "dev",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"登陆成功\",\"data\":{\"token\":\"111...\"}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/power/sub": {
            "post": {
                "description": "数据已加密",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "验证相关"
                ],
                "summary": "扣点[数据已加密]",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "扣除点数不可为负",
                        "name": "point",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"扣除成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/edit": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "编辑用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "用户状态1=可用 2=不可用",
                        "name": "state",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getAll": {
            "get": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "获取的数量",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户状态1=可用 2=不可用",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "user",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"查询成功\",\"count\":2,\"data\":[{\"uid\":1001,\"user\":\"99v66\",\"state\":2,\"time\":1624982434},{\"uid\":1002,\"user\":\"529887136\",\"state\":1,\"time\":1624983043}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getInfo": {
            "get": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "获取个人信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"更新成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"登录成功\",\"data\":\"eyJ1aWQiOjEwMDEsImFpZCI6MCwidGltZSI6MTYyNDk4MjQ5OSwicG93ZXIiOjIsInRpbWVfdXBkYXRhIjowLCJ0aW1lX0xvZ2luIjowLCJzaWduIjoiNWQ3OGYyYTE4OGY1Y2NhNjQ2NDdhYzI4N2MyMmVlMzIifQ==\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/newPwd": {
            "post": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "旧密码",
                        "name": "oldPwd",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "newPwd",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"修改成功,请重新登录\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/out": {
            "get": {
                "description": "将造成所有软件token失效,暂停接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "退出",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"修改成功,请重新登录\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/reg": {
            "post": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"注册成功,请前往登录\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/upInfo": {
            "post": {
                "description": "详细描述",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "更新个人信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录返回的token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "更新内容",
                        "name": "msg",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"更新成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:52777",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "验证系统 API",
	Description: "统一编码格式-UTF8\n",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
