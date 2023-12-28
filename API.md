# 写在前面

中间件：

```go
func cors() // 跨域，所有API都带，要不然前前端无法带上后端设置的cookie
func auth() // 鉴权，主要是评论模块和查询个人信息相关模块，会在跨域中间件后调用
```

后端接口定义：

```json
{
    "Code": 200,  										// 状态码，因为返回非200，浏览器会报错，所以要单开Code表示状态
    "Message": "Register success.",   // 消息
    "Data": [													// payload，是一个列表，如果返回一个数组，会再嵌套一层object array
        null
    ]
}
```

# /ping

用来检测后端服务器是否存活，GET方法，直接用浏览器访问http://localhost:8080/ping，能够看到

```json
{
    "Code": 200,
    "Message": "Pong!",
    "Data": [
        null
    ]
}
```

即可，说明您已经正确启动了服务器

# /user

这个和用户登录注册等有关

## /register

注册，三个参数非空

```json
// method
[POST]

// backend routing
http://localhost:8080/user/register

// middleware
cors()

// input
// 都必须有，不能缺参数
{
    "name": "test",
    "email": "1221@q.com",
    "password": "test"
}

// output
{
    "Code": 200,
    "Message": "Register success.",
    "Data": [
        null
    ]
}
```

## /login

登录，用户名和邮箱二选一，密码必须有。注意密码是哈希过的，注册时要记住密码，不能直接拿数据库的登录

```json
// method
[POST]

// backend routing
http://localhost:8080/user/login

// middleware
cors()

// input
{
    "name": "test",
    // "email": "1221@q.com", 二选一，注释上面那个也行
    "password": "test"
}

// output
{
    "Code": 200,
    "Message": "Login success.",
    "Data": [
        null
    ]
}
```

## /logout

注销，实际上是让cookie立马过期

```json
// method
[GET]

// backend routing
http://localhost:8080/user/logout

// middleware
cors()

// input
{}

// output，没有cookie就注销会提示
{
    "Code": 200,
    "Message": "Logout success",
    "Data": [
        null
    ]
}
```

# /site

处理景点相关参数，没加鉴权

## /search

用来获取景点信息，没加鉴权

```json
// method
[POST]

// backend routing
http://localhost:8080/site/search

// middleware
cors()

// input，四个随便来一个都行。后端数据库保证景点名不重复
{
    "continent": "亚洲",
    "country": "中国",
    "city": "黄山市",
    "siteName": "黄山"
}

// output，注意嵌套关系
{
    "Code": 200,
    "Message": "Sites are: ",
    "Data": [
        [
            {
                "continent": "亚洲",
                "country": "中国",
                "city": "黄山市",
                "siteName": "黄山",
                "description": "黄山是世界文化与自然遗产、世界地质公园、世界生物圈保护区，是国家级风景名胜区、全国文明风景旅游区、国家5A级旅游景区，与长江、长城、黄河同为中华壮丽山河和灿烂文化的杰出代表，被世人誉为“人间仙境”、“天下第一奇山”，素以奇松、怪石、云海、温泉、冬雪“五绝”著称于世。境内群峰竞秀，怪石林立，有千米以上高峰88座，“莲花”、“光明顶”、“天都”三大主峰，海拔均逾1800米。明代大旅行家徐霞客曾两次登临黄山，赞叹道：“薄海内外无如徽之黄山，登黄山天下无山，观止矣！”后人据此概括为“五岳归来不看山，黄山归来不看岳”。",
                "grade": 0,
                "gradeCount": 0,
                "latitude": 118.2014318,
                "longitude": 30.305406,
                "openTime": "7：00-17:00",
                "season": "夏季",
                "ticket": 932756,
                "free": false,
                "link": "https://www.bilibili.com/video/BV1no4y1w7Dc/?spm_id_from=333.337.search-card.all.click&vd_source=df79d8622b20f745687871ec7a054788"
            }
        ]
    ]
}
```

## /get_remark

获取某景点的所有评论，这个没加鉴权

```json
// method
[POST]

// backend routing
http://localhost:8080/site/get_remark

// middleware
cors()

// input，一个景点名就行
{
    "siteName":"黄山"
}

// output，注意嵌套关系
{
    "Code": 200,
    "Message": "Remarks are: ",
    "Data": [
        {
            "siteName": "黄山",
            "average": 6.335,
            "totalCnt": 2,
            "remarks": [
                {
                    "userName": "zxq",
                    "content": "垃圾",
                    "mark": 2.77
                },
                {
                    "userName": "xyz",
                    "content": "不错",
                    "mark": 9.9
                }
            ]
        }
    ]
}
```

# /remark

处理评论相关，必须鉴权

## /add

```json
// method
[POST]

// backend routing
http://localhost:8080/remark/add

// middleware
cors()
auth()

// input
{
    "userName": "test",
    "siteName": "黄山",
    "content": "你好",
    "mark": 9.83
}

// output，注意嵌套关系
{
    "Code": 200,
    "Message": "Remark success.",
    "Data": [
        null
    ]
}

// 已经评论过了，不能添加，只能修改
{
    "Code": 412,
    "Message": "You have remarked before.",
    "Data": [
        null
    ]
}
```

## /delete

```json
// method
[POST]

// backend routing
http://localhost:8080/remark/delete

// middleware
cors()
auth()

// input
{
    "userName": "test",
    "siteName": "黄山"
}

// output，注意嵌套关系
{
    "Code": 200,
    "Message": "Delete success.",
    "Data": [
        null
    ]
}

// 其他不合法输出
...
```

## /modify

```json
// method
[POST]

// backend routing
http://localhost:8080/remark/modify

// middleware
cors()
auth()

// input
{
    "userName": "test",
    "siteName": "黄山",
    "content": "牛魔",
    "mark": 0.88
}
// output，注意嵌套关系
{
    "Code": 200,
    "Message": "Modify success.",
    "Data": [
        null
    ]
}

// 其他不合法输出
...
```

## /user

获取用户曾经所有评论

```json
// method
[GET]

// backend routing
http://localhost:8080/remark/user

// middleware
cors()
auth()

// input
{}

// output，注意嵌套关系
{
    "Code": 200,
    "Message": "Your remarks are: ",
    "Data": [
        [
            {
                "siteName": "黄山",
                "content": "不错",
                "mark": 9.9
            }// ,...
        ]
    ]
}

// 其他不合法输出
...
```

