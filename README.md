

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
