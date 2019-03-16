---
layout: default
title: README

---

# README
创建时间: 2016/03/17 16:13:00  修改时间: 2016/03/18 17:24:46 作者:lijiao

----

## 摘要

Keystone API的使用，这里的操作在Operation目录下完成。

## 说明

Operation中的脚本使用方式如下：

	$ ./ACTION.sh  APIVERSION/RESOURCE  [DATAFILE|DATA]   [TOKENID]

在keystone.conf配置的超级TOKEN时"ADMIN"，在keystone中没有任何用户的时候，用超级TOKEN进行操作。

## 授权策略

授权策略在keystone.conf中配置，这里使用的是keystone中提供的policy.v3cloudsample.json。

在policy.v3cloudsample.json中规划的策略，简单概括就是：“两个角色，一个管理域”

角色admin:

	admin可以对不属于自己的资源进行管理。

角色service:

	service拥有且只拥有验证任意Token的权利。

管理域:

	在管理域中拥有admin角色的用户，具有对云进行管理对权限。

管理域是在policy.v3cloudsample.json文件中指定的，需要用超级TOKEN创建一个domain，将domain的id填写到policy文件中，重启keystone。

>授权策略完全可以按照自己的需要设置，包括并且不限于增加新的角色、为新的角色设置不同的权限等。

## Domain

创建名为cloud_admin的domain:

	$./post.sh v3/domains ./json/domain_cloud_admin.json ADMIN
	{
	    "domain": {
	        "description": "Domain: cloud_admin",
	            "enabled": true,
	            "id": "c20b7508d23342148a7cf1b64112706f",
	            "links": {
	                "self": "http://127.0.0.1:35357/v3/domains/c20b7508d23342148a7cf1b64112706f"
	            },
	            "name": "cloud_admin"
	    }
	}

查看系统中的domains:

	$./get.sh v3/domains ADMIN
	{
	    "domains": [
	    {
	        "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
	            "enabled": true,
	            "id": "default",
	            "links": {
	                "self": "http://127.0.0.1:35357/v3/domains/default"
	            },
	            "name": "Default"
	    },
	    {
	        "description": "Domain: cloud_admin",
	        "enabled": true,
	        "id": "c20b7508d23342148a7cf1b64112706f",
	        "links": {
	            "self": "http://127.0.0.1:35357/v3/domains/c20b7508d23342148a7cf1b64112706f"
	        },
	        "name": "cloud_admin"
	    }
	    ],
	        "links": {
	            "next": null,
	            "previous": null,
	            "self": "http://127.0.0.1:35357/v3/domains"
	        }
	}

## User

必须指定在哪个Domain中创建用户，编辑 ./json/user_admin.json，修改其中的domain_id

	"domain_id": "c20b7508d23342148a7cf1b64112706f",

然后执行:

	$./post.sh  v3/users ./json/user_admin.json  ADMIN
	{
	    "user": {
	        "description": "admin@jd.com",
	        "domain_id": "c20b7508d23342148a7cf1b64112706f",
	        "email": "admin@mail.com",
	        "enabled": true,
	        "id": "e4247e41f1a74d7d9d33c26c33e80a29",
	        "links": {
	            "self": "http://127.0.0.1:35357/v3/users/e4247e41f1a74d7d9d33c26c33e80a29"
	        },
	        "name": "admin@mail.com"
	    }
	}

一个Domain中可以有多个用户加入, Domain的管理员可以查看Domain中的用户的列表:

	$./get.sh v3/users?domain_id=c20b7508d23342148a7cf1b64112706f  68856a6b4d474602b477c0805090b00c
	{
	    "links": {
	        "next": null,
	            "previous": null,
	            "self": "http://127.0.0.1:35357/v3/users?domain_id=c20b7508d23342148a7cf1b64112706f"
	    },
	        "users": [
	        {
	            "description": "admin@jd.com",
	            "domain_id": "c20b7508d23342148a7cf1b64112706f",
	            "email": "admin@mail.com",
	            "enabled": true,
	            "id": "e4247e41f1a74d7d9d33c26c33e80a29",
	            "links": {
	                "self": "http://127.0.0.1:35357/v3/users/e4247e41f1a74d7d9d33c26c33e80a29"
	            },
	            "name": "admin@mail.com"
	        },
	        {
	            "description": "userA@mail.com",
	            "domain_id": "c20b7508d23342148a7cf1b64112706f",
	            "email": "userA@mail.com",
	            "enabled": true,
	            "id": "aef6a0c06bcf4dad98c817a8554e245e",
	            "links": {
	                "self": "http://127.0.0.1:35357/v3/users/aef6a0c06bcf4dad98c817a8554e245e"
	            },
	            "name": "userA@mail.com"
	        }
	    ]
	}

>注意这里用的Token(688...)必须是作用于domain(c20b...)的，并且创建这个Token的用户是domain的管理员。

>Token的作用范围，见Auth一节。

## Auth

用户通过密码或者其它方式通过认证以后，会得到一个token，每个token都有一个ID号，存放在返回的http头的X-Subject-Token字段中。

Token分为有作用域的(scoped)和无作用域的(unscoped)，这里的作用域指的是Token适用于哪一个project或者domain。对project或者domain进行操作时，Token必须是有作用域的，并且适用于要操作的目标。

用户要使用Token的时候，只需要在http头的X-Auth-Token字段中指定使用的Token的ID。

对刚创建的用户admin@mail.com进行认证:

	$ ./post.sh v3/auth/tokens ./json/auth_user_admin_unscope_name.json
	> POST /v3/auth/tokens HTTP/1.1
	> User-Agent: curl/7.35.0
	> Host: 127.0.0.1:35357
	> Accept: */*
	> Content-Type: application/json
	> X-Auth-Token: ADMIN
	> Content-Length: 221
	>
	} [data not shown]
	* upload completely sent off: 221 out of 221 bytes
	< HTTP/1.1 201 Created
	< X-Subject-Token: 4a457e2b765a46c49135eddc66bb1488
	< Vary: X-Auth-Token
	< Content-Type: application/json
	< Content-Length: 335
	< X-Openstack-Request-Id: req-a0809f8f-18fd-457f-997f-962347133146
	< Date: Thu, 17 Mar 2016 09:31:31 GMT
	<
	{ [data not shown]
	
	{
	    "token": {
	        "audit_ids": [
	            "v0i4FpunS_CgmQS2D5rMWw"
	        ],
	        "expires_at": "2016-03-17T10:31:31.554213Z",
	        "extras": {},
	        "issued_at": "2016-03-17T09:31:31.554246Z",
	        "methods": [
	            "password"
	        ],
	        "user": {
	            "domain": {
	                "id": "c20b7508d23342148a7cf1b64112706f",
	                "name": "cloud_admin"
	            },
	            "id": "e4247e41f1a74d7d9d33c26c33e80a29",
	            "name": "admin@mail.com"
	        }
	    }
	}

验证刚获得的TOKEN(000a4cd444cd9c7029eff91f57f4):

	$./verify.sh v3/auth/tokens 4a457e2b765a46c49135eddc66bb1488 ADMIN
	* Hostname was NOT found in DNS cache
	*   Trying 127.0.0.1...
	* Connected to 127.0.0.1 (127.0.0.1) port 35357 (#0)
	> GET /v3/auth/tokens HTTP/1.1
	> User-Agent: curl/7.35.0
	> Host: 127.0.0.1:35357
	> Accept: */*
	> X-Auth-Token: ADMIN
	> X-Subject-Token: 4a457e2b765a46c49135eddc66bb1488
	>
	< HTTP/1.1 200 OK
	< X-Subject-Token: 4a457e2b765a46c49135eddc66bb1488
	< Vary: X-Auth-Token
	< Content-Type: application/json
	< Content-Length: 335
	< X-Openstack-Request-Id: req-4d4c893c-a3f6-4bb6-9a30-800059fe89c9
	< Date: Fri, 18 Mar 2016 03:47:39 GMT
	<
	{ [data not shown]
	* Connection #0 to host 127.0.0.1 left intact
	{
		"token": {
			"audit_ids": [
				"0LXKZVImRP2TNuKZamxJ2g"
			],
			"expires_at": "2016-03-18T04:46:47.082707Z",
			"extras": {},
			"issued_at": "2016-03-18T03:46:47.082742Z",
			"methods": [
				"password"
			],
			"user": {
				"domain": {
					"id": "c20b7508d23342148a7cf1b64112706f",
					"name": "cloud_admin"
				},
				"id": "e4247e41f1a74d7d9d33c26c33e80a29",
				"name": "admin@mail.com"
			}
		}
	}

## Role

创建角色:

	$ ./post.sh v3/roles  json/role_admin.json ADMIN
	{
		"role": {
			"id": "e7947105f050479cab0917bfea9422eb",
			"links": {
				"self": "http://127.0.0.1:35357/v3/roles/e7947105f050479cab0917bfea9422eb"
			},
			"name": "admin"
		}
	}

用户的角色是相对于domain或者project变化的，同一个用户可以在不同的domain或者project中拥有不同的角色。

使用超级TOKEN，在cloud_admin域中赋予用户(admin@mail.com)admin角色:

	$ ./put.sh v3/domains/c20b7508d23342148a7cf1b64112706f/users/e4247e41f1a74d7d9d33c26c33e80a29/roles/e7947105f050479cab0917bfea9422eb

查看用户(admin@mail.com)在cloud_admin中的角色:

	$ ./get.sh v3/domains/c20b7508d23342148a7cf1b64112706f/users/e4247e41f1a74d7d9d33c26c33e80a29/roles ADMIN
	{
		"links": {
			"next": null,
			"previous": null,
			"self": "http://127.0.0.1:35357/v3/domains/c20b7508d23342148a7cf1b64112706f/users/e4247e41f1a74d7d9d33c26c33e80a29/roles"
		},
		"roles": [
			{
				"id": "e7947105f050479cab0917bfea9422eb",
				"links": {
					"self": "http://127.0.0.1:35357/v3/roles/e7947105f050479cab0917bfea9422eb"
				},
				"name": "admin"
			}
		]
	}

角色除了可以直接授予用户，还可以授予用户组，用户组拥有的角色，其中的用户都拥有。

## 文献

1. http://developer.openstack.org/api-ref-identity-v3.html
2. http://developer.openstack.org/api-ref-identity-v3-ext.html
3. http://developer.openstack.org/api-ref-identity-v2.html 
