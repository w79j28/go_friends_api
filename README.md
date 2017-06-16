

## 开发框架说明
* Go、[xorm](http://www.xorm.io)、[SwaggerUI](https://w79j28.github.io/go_friends_api)
* PostgreSql


## 项目背景
For any application with a need to build its own social network, "Friends Management" is a
common requirement which ussually starts off simple but can grow in complexity depending
on the application's use case.
Usually, applications would start with features like "Friend", "Unfriend", "Block", "Receive
Updates" etc.

## API说明
#### 1 /usr/friends    [`POST`] 

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

#### 2 /usr/friends/list    [`POST`] 

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


