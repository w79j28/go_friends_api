
## 在线访问地址
  [Friends Management](https://wangjingzhu.run.aws-jp01-pr.ice.predix.io/swagger-ui/#!/user)

## 开发框架说明
* Go、[gin](https://github.com/gin-gonic/gin)、[xorm](http://www.xorm.io)、[SwaggerUI](https://wangjingzhu.run.aws-jp01-pr.ice.predix.io/swagger-ui/)
* PostgreSql


## 项目背景
For any application with a need to build its own social network, "Friends Management" is a
common requirement which ussually starts off simple but can grow in complexity depending
on the application's use case.
Usually, applications would start with features like "Friend", "Unfriend", "Block", "Receive
Updates" etc.

## 项目说明
* [Swagger 注释说明](https://github.com/yvasiyarov/swagger/wiki/Declarative-Comments-Format)
* swagger/main.go 生成Swagger文档
* 配置文件config/app_config.ini,首次运行时创建或自行创建
 
 `app_config.ini`
```java
  #数据库连接
  dburl            = postgres://user:password@ip:port/dbname?sslmode=disable
  #端口 predix时为cloud
  port             = port
  #Swagger Basepath
  swagger_basepath = http://ip:port
```   

* 运行后 Swagger UI 地址：http://ip:port/swagger-ui/

    如：[http://localhost:9090/swagger-ui/](http://localhost:9090/swagger-ui/)

## API说明

[https://wangjingzhu.run.aws-jp01-pr.ice.predix.io/swagger-ui/](https://wangjingzhu.run.aws-jp01-pr.ice.predix.io/swagger-ui/#!/user)

#### 1 /user/friends    [`POST`] 

*Create a friend connection between two email addresses.*

*JSON request:*
```json
  {
    "friends":[
        "andy@example.com",
        "john@example.com"
     ]
  }
```
*JSON response on success:*
```json
  {
    "success": true
  }
```  

#### 2 /user/friends/list    [`POST`] 

*Retrieve the friends list for an email address.*

*JSON request:*
```json
  {
    "email": "andy@example.com"
  }
```
*JSON response on success:*
```json
  {
     "success": true,
     "friends" : [
         "john@example.com"
     ],
     "count" : 1
  }
```  

#### 3 /user/friends/common    [`POST`] 

*Retrieve the common friends list between two email addresses.*

*JSON request:*
```json
  {
     "friends":[
         "andy@example.com",
         "john@example.com"
     ]
  }
```
*JSON response on success:*
```json
  {
     "success": true,
     "friends":[
         "common@example.com"
     ],
     "count" : 1
  }
```  

#### 4 /user/friend/subscribe    [`POST`] 

*Subscribe to updates from an email address.*

That "subscribing to updates" is NOT equivalent to "adding a friend connection".

*JSON request:*
```json
  {
      "requestor": "lisa@example.com",
      "target": "john@example.com"
  }
```
*JSON response on success:*
```json
  {
      "success": true
  }
```  

#### 5 /user/friend/block    [`POST`] 

*Block updates from an email address.*

*Suppose "andy@example.com" blocks "john@example.com":*
* if they are connected as friends, then "andy" will no longer receive notifications from "john"
* if they are not connected as friends, then no new friends connection can be added

*JSON request:*
```json
  {
      "requestor": "andy@example.com",
      "target": "john@example.com"
  }
```
*JSON response on success:*
```json
  {
      "success": true
  }
```  

#### 6 /user/friends/sender    [`POST`] 

*Retrieve all email addresses that can receive updates from an email address.*

*Eligibility for receiving updates from i.e. "john@example.com":*
* has not blocked updates from "john@example.com", and
* at least one of the following:
    * has a friend connection with "john@example.com"
    * has subscribed to updates from "john@example.com"
    * has been @mentioned in the update

*JSON request:*
```json
  {
     "sender": "john@example.com",
     "text": "Hello World! kate@example.com"
  }
```
*JSON response on success:*
```json
  {
     "success": true,
     "recipients" : [
        "lisa@example.com",
        "kate@example.com"
     ]
  }
```  
