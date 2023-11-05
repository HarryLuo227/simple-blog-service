# Blog-Service API Testing

---

## 第一步 - 權限驗證

資料庫已存在以下使用者資訊，請登入取得 Token ，以利後續 API 使用。

> 使用者:tester
> 
> 密碼:123456

```Bash
curl -X POST -H 'user: tester' -H 'password: 123456' http://127.0.0.1:8000/auth
```

## 第二步 - 標籤相關 API 測試

*如有人想進行測試，因底下的 Token 有設定過期時間，可以下載檔案到測試者本機然後取代原有的過期 Token 並自行更新以進行測試。*

### TAG

#### 新增一個標籤

##### 新增一個資料庫不存在的標籤

```Bash
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags -F 'name=Go' -F 'created_by=tester'
```

##### 新增一個資料庫已存在的標籤 

```Bash
# (同建立者)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags -F 'name=Go' -F 'created_by=tester'

# (不同建立者)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags -F 'name=Go' -F 'created_by=boss'
```

#### 取得所有標籤

```Bash
curl -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags
```

#### 取得特定標籤

##### 取得一個資料庫已存在的標籤

```Bash
curl -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags/{id}
```

##### 取得一個資料庫不存在的標籤

```Bash
curl -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags/{id}
```

#### 更新特定標籤

##### 更新一個資料庫已存在的標籤

```Bash
curl -X PUT -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags/{id} -F 'name=C' -F 'modified_by=tester'
```

##### 更新一個資料庫不存在的標籤

```Bash
curl -X PUT -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags/{id} -F 'name=C' -F 'modified_by=tester'
```

#### 刪除特定標籤

##### 刪除一個資料庫已存在的標籤

```Bash
curl -X DELETE -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags/{id}
```

##### 刪除一個資料庫不存在的標籤

```Bash
curl -X DELETE -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/tags/{id}
```

### Article

#### 新增一篇文章

##### 新增一個資料庫不存在的文章

```Bash
# (單個標籤，標籤存在)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=1' -F 'title=go-simple-blog-service-test-1' -F 'desc=testing tag exists' -F 'content=Test add article which not exists, with single exist tag.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'
# (單個標籤，標籤不存在)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=2' -F 'title=go-simple-blog-service-test-2' -F 'desc=testing tag not exists' -F 'content=Test add article which not exists, single tag is not exists.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'

# (多個標籤，標籤存在)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=1' -F 'tag_id=5' -F 'title=article-with-multiple-tags-test-1' -F 'desc=testing each tags in list exists' -F 'content=Test add article which not exists, with all exist tag.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'
# (多個標籤，標籤部份不存在)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=1' -F 'tag_id=2' -F 'title=article-with-multiple-tags-test-2' -F 'desc=testing each tags in list partially exists' -F 'content=Test add article which not exists, with not all exist tag.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'
```

##### 新增一個資料庫已存在的文章

```Bash
# (單個標籤，標籤存在且相同)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=1' -F 'title=go-simple-blog-service-test-1' -F 'desc=testing tag exists' -F 'content=Test add article which exists, with single exist tag and is same tag.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'
# (單個標籤，標籤存在但不相同)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=5' -F 'title=go-simple-blog-service-test-1' -F 'desc=testing tag exists' -F 'content=Test add article which exists, with single exist tag and is not the same tag.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'
# (單個標籤，標籤不存在)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=2' -F 'title=go-simple-blog-service-test-1' -F 'desc=testing tag not exists' -F 'content=Test add article which not exists, single tag is not exists.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'

# (多個標籤，標籤存在且相同)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=1' -F 'tag_id=5' -F 'title=article-with-multiple-tags-test-1' -F 'desc=testing each tags in list exists' -F 'content=Test add article which not exists, with all exist tag.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'
# (多個標籤，標籤存在但不相同)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=1' -F 'tag_id=6' -F 'title=article-with-multiple-tags-test-1' -F 'desc=testing each tags in list exists' -F 'content=Test add article which exists, with all exist tag.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'
# (多個標籤，標籤部份不存在)
curl -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=1' -F 'tag_id=2' -F 'title=article-with-multiple-tags-test-1' -F 'desc=testing each tags in list partially exists' -F 'content=Test add article which not exists, with not all exist tag.' -F 'cover_image_url=https://en.wikipedia.org/wiki/Go_%28programming_language%29#/media/File:Go_Logo_Blue.svg' -F 'created_by=tester'
```

#### 取得所有文章

```Bash
curl -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles -F 'tag_id=6'
```

#### 取得特定文章

##### 取得一個資料庫已存在的文章

```Bash
curl -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id}
```

##### 取得一個資料庫不存在的文章

```Bash
curl -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id}
```

#### 更新特定文章

##### 更新一個資料庫已存在的文章

```Bash
# [無帶標籤]
# (更新後標題與創作者並無重複)
curl -X PUT -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id} -F 'title=C-blog-implement' -F 'desc=Test update article, the article after updates is not unique.' -F 'modified_by=tester'
# (更新後標題與創作者重複)
curl -X PUT -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id} -F 'title=C-blog-implement' -F 'modified_by=tester'

# [有帶單一標籤]
# (更新後標題與創作者並無重複)
curl -X PUT -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id} -F 'tag_id=6' -F 'title=go-simple-blog' -F 'desc=Test update article, the article after updates is not unique.' -F 'modified_by=tester'
# (更新後標題與創作者重複)
curl -X PUT -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id} -F 'tag_id=6' -F 'title=go-simple-blog' -F 'modified_by=tester'

# [有帶多個標籤]
# (更新後標題與創作者並無重複)
curl -X PUT -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id} -F 'tag_id=5' -F 'tag_id=6' -F 'title=test-update-article-with-multiple-tags' -F 'desc=Test update article with multiple tag, the article after updates is not unique.' -F 'modified_by=tester'
# (更新後標題與創作者重複)
curl -X PUT -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id} -F 'tag_id=1' -F 'tag_id=6' -F 'title=test-update-article-with-multiple-tags' -F 'modified_by=tester'
```

##### 更新一個資料庫不存在的文章

```Bash
curl -X PUT -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id} -F 'name=C' -F 'modified_by=tester'
```

#### 刪除特定文章

```Bash
curl -X DELETE -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoiZjVkMTI3OGU4MTA5ZWRkOTRlMWU0MTk3ZTA0ODczYjkiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiZXhwIjoxNjk5MTc2NTA1LCJpc3MiOiJibG9nLXNlcnZpY2UifQ.7d0QL6D692LB3-sBZbfmCUAI865ZLwB9Mnx87TSEDWE' http://127.0.0.1:8000/api/v1/articles/{id}
```
