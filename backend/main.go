package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages []Message `json:"messages"`
	Image    string    `json:"image,omitempty"`
}

type ChatResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID      string   `json:"id"`
		Choices []struct {
			FinishReason string `json:"finish_reason"`
			Message      Message `json:"message"`
		} `json:"choices"`
		Usage struct {
			PromptTokens     int `json:"prompt_tokens"`
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		} `json:"usage"`
	} `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var apiKey = "de9a57b958df4fb8b71834f6d08f7478.p9awjnrRb5ZA9AS8"
var apiBaseURL = "https://open.bigmodel.cn/api/paas/v4/chat/completions"

func chatHandler(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + apiKey,
	}

	messages := req.Messages

	if req.Image != "" {
		imageDescription := describeImage(req.Image)
		messages = append([]Message{{
			Role:    "user",
			Content: imageDescription,
		}}, messages...)
	}

	requestBody := map[string]interface{}{
		"model":    "glm-4v",
		"messages": messages,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	reqBody := bytes.NewBuffer(jsonData)
	httpReq, err := http.NewRequest("POST", apiBaseURL, reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	for key, value := range headers {
		httpReq.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, chatResp)
}

func describeImage(imageBase64 string) string {
	return fmt.Sprintf(`请仔细分析以下图片并详细描述其中的内容：

图片数据（Base64）：%s

请描述：
1. 图片中的主要对象和场景
2. 图片的颜色和构图
3. 任何文字或标识
4. 图片的整体氛围和主题

然后基于图片内容继续之前的对话。`, imageBase64[:100]+"...")
}

func uploadImageHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    400,
			Message: "请上传图片文件",
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	base64String := base64.StdEncoding.EncodeToString(fileBytes)

	fileName := uuid.New().String() + "_" + file.Filename
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"fileName": fileName,
			"base64":   "data:" + file.Header.Get("Content-Type") + ";base64," + base64String,
		},
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.POST("/api/chat", chatHandler)
	r.POST("/api/upload", uploadImageHandler)
	r.Static("/static", "../frontend/dist")
	r.StaticFile("/", "../frontend/dist/index.html")

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}