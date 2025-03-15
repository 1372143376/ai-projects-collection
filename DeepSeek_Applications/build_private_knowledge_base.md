# 使用 DeepSeek 搭建私有知识库教程

## 1. 环境准备

### 1.1 系统要求
- Linux 或 macOS 系统
- Python 3.8+
- Docker 20.10+

### 1.2 安装依赖
```bash
# 安装 Python 依赖
pip install deepseek-sdk flask pymongo

# 安装 Docker
# Ubuntu
sudo apt-get update
sudo apt-get install docker.io

# macOS
brew install docker
```

## 2. 配置 MongoDB 数据库

```bash
# 启动 MongoDB 容器
docker run -d --name deepseek-mongo -p 27017:27017 mongo:latest

# 创建数据库和用户
docker exec -it deepseek-mongo mongo admin --eval '
db.createUser({
  user: "deepseek",
  pwd: "your_password",
  roles: [{ role: "readWrite", db: "knowledge_base" }]
})'
```

## 3. 配置 DeepSeek 服务

### 3.1 创建配置文件
```bash
mkdir config
cat > config/settings.yaml <<EOF
database:
  host: localhost
  port: 27017
  name: knowledge_base
  username: deepseek
  password: your_password

server:
  host: 0.0.0.0
  port: 5000
EOF
```

### 3.2 启动服务
```bash
deepseek-server --config config/settings.yaml
```

## 4. 添加知识数据

### 4.1 通过 API 添加数据
```python
import requests

url = "http://localhost:5000/api/v1/knowledge"
headers = {"Content-Type": "application/json"}
data = {
    "title": "示例知识",
    "content": "这是示例知识内容",
    "tags": ["示例", "教程"]
}

response = requests.post(url, json=data, headers=headers)
print(response.json())
```

### 4.2 批量导入数据
```bash
deepseek-import --file data.json --config config/settings.yaml
```

## 5. 查询知识库

```python
import requests

url = "http://localhost:5000/api/v1/knowledge/search"
params = {"query": "示例"}

response = requests.get(url, params=params)
print(response.json())
```

## 6. 部署到生产环境

### 6.1 使用 Docker Compose
```yaml
version: '3'
services:
  mongo:
    image: mongo:latest
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_ROOT_USERNAME: deepseek
      MONGO_INITDB_ROOT_PASSWORD: your_password

  deepseek:
    build: .
    ports:
      - "5000:5000"
    depends_on:
      - mongo
```

### 6.2 启动服务
```bash
docker-compose up -d
```

## 7. 维护与更新

- 定期备份数据库
```bash
docker exec deepseek-mongo mongodump --out /backup
```

- 更新 DeepSeek 版本
```bash
docker-compose pull
docker-compose up -d
```

## 8. 常见问题

Q: 如何重置管理员密码？
A: 通过 MongoDB 命令行修改用户密码
```bash
docker exec -it deepseek-mongo mongo admin --eval '
db.changeUserPassword("deepseek", "new_password")'
```

Q: 如何扩展存储容量？
A: 修改 Docker 卷配置或使用外部 MongoDB 集群
