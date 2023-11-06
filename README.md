# Simple-Blog-Service

這是一個簡易的、具備了 REST 架構風格的部落格後端專案，實作了 `文章` 與 `文章標籤` 的 CRUD 功能，並為之加上了權限驗證機制。

## 可執行的操作

### 標籤

- 列出所有標籤
- 取得特定標籤
- 新增標籤
- 修改特定標籤
- 刪除特定標籤

### 文章

- 列出有特定標籤的所有文章
- 取得特定文章
- 新增文章
- 修改特定文章
- 刪除特定文章

## RESTful API Document

專案本身有整合 Swagger 套件產生的 [API 文件](http://127.0.0.1:8000/swagger/index.html) 可以在啟動服務後查看，或者已整理好的相關 API 指令在 [APITest.md](APITest.md) 檔案。

![API-Document](/tmp/Simple-Blog-Service_API-DOC.png "API-Document")

## 其它

為了在缺乏資料儲存在資料庫的情況下，讓 API 測試不須按照先新增在修改等特定順序進行測試，不論在新增文章或標籤的功能裡，已自動設計成若無相同的資料則進行新增動作，有相同的資料則直接更新該相同資料。