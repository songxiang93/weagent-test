# Smart Conversation System

一个支持文字和图片多模态对话的智能系统，使用Go后端和Vue前端，集成了智谱AI的GLM-4.6v模型。

## 功能特性

- 🎯 多模态对话：支持文字和图片输入
- 🤖 GLM-4.6v模型：强大的图像识别和对话能力
- 📱 响应式设计：支持桌面和移动端
- 🚀 实时交互：流畅的聊天体验
- 🖼️ 图片上传：支持多种图片格式
- 🔒 安全部署：Docker容器化部署

## 技术栈

### 后端
- Go 1.21
- Gin Web Framework
- GLM-4.6v API集成

### 前端
- Vue 3
- Vite
- Axios

### 部署
- Docker
- Docker Compose
- Nginx

## 快速开始

### 环境要求
- Go 1.21+
- Node.js 18+
- Docker (可选)

### 开发环境

1. 克隆项目
```bash
git clone <repository-url>
cd smart-conversation-system
```

2. 启动后端
```bash
cd backend
go mod tidy
go run main.go
```

3. 启动前端
```bash
cd ../frontend
npm install
npm run dev
```

4. 访问应用
- 前端：http://localhost:3000
- 后端API：http://localhost:8080

### Docker部署

```bash
# 构建并启动所有服务
docker-compose up --build

# 后台运行
docker-compose up -d

# 停止服务
docker-compose down
```

## API文档

### 聊天接口
- **URL**: `POST /api/chat`
- **Content-Type**: `application/json`
- **Request**:
  ```json
  {
    "messages": [
      {
        "role": "user",
        "content": "你好"
      }
    ],
    "image": "data:image/jpeg;base64,..."
  }
  ```

### 图片上传接口
- **URL**: `POST /api/upload`
- **Content-Type**: `multipart/form-data`
- **Request**: Form data with image file

## 配置

### 后端配置
编辑 `backend/main.go` 中的API配置：
```go
var apiKey = "your-api-key"
var apiBaseURL = "https://open.bigmodel.cn/api/paas/v4/chat/completions"
```

### 前端配置
编辑 `frontend/vite.config.js` 中的代理配置：
```javascript
proxy: {
  '/api': {
    target: 'http://localhost:8080',
    changeOrigin: true
  }
}
```

## 项目结构

```
smart-conversation-system/
├── backend/
│   ├── main.go              # Go后端主程序
│   ├── go.mod               # Go模块依赖
│   ├── Dockerfile           # Docker构建文件
│   └── README.md            # 后端说明
├── frontend/
│   ├── src/
│   │   ├── App.vue          # 主组件
│   │   ├── main.js          # 应用入口
│   │   └── index.html       # HTML模板
│   ├── package.json         # Node.js依赖
│   ├── vite.config.js       # Vite配置
│   ├── Dockerfile           # Docker构建文件
│   └── README.md            # 前端说明
├── docker-compose.yml       # Docker编排配置
├── nginx.conf               # Nginx配置
├── .gitignore              # Git忽略文件
└── README.md              # 项目说明
```

## 使用说明

1. **发送文字消息**
   - 在输入框中输入文字
   - 点击发送按钮或按Enter键

2. **发送图片**
   - 点击相机图标选择图片
   - 输入文字描述（可选）
   - 点击发送按钮

3. **查看对话历史**
   - 所有对话消息会显示在聊天区域
   - 支持滚动查看历史消息

## 开发指南

### 添加新功能
1. 后端：修改 `backend/main.go` 添加新的API端点
2. 前端：修改 `frontend/src/App.vue` 添加新的UI组件
3. 重新构建并测试

### 部署优化
- 使用环境变量管理敏感配置
- 配置HTTPS和域名
- 设置负载均衡和缓存

## 故障排除

### 常见问题

1. **后端启动失败**
   - 检查Go版本是否为1.21+
   - 确认网络连接正常

2. **前端无法连接后端**
   - 检查后端服务是否启动
   - 确认代理配置正确

3. **图片上传失败**
   - 检查图片格式是否支持
   - 确认图片大小限制

### 日志查看

```bash
# 查看后端日志
docker-compose logs backend

# 查看前端日志
docker-compose logs frontend

# 查看Nginx日志
docker-compose logs nginx
```

## 贡献

欢迎提交Issue和Pull Request来改进这个项目。

## 许可证

MIT License