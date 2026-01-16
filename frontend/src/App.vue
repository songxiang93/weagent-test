<template>
  <div id="app">
    <div class="chat-container">
      <div class="chat-header">
        <h1>智能对话系统</h1>
        <p>支持文字和图片的多模态对话</p>
      </div>

      <div class="chat-messages" ref="messagesContainer">
        <div v-for="(message, index) in messages" :key="index"
             :class="['message', message.role]">
          <div class="message-avatar">{{ message.role === 'user' ? '你' : 'AI' }}</div>
          <div class="message-content">
            <div v-if="message.image" class="message-image-container">
              <img :src="message.image" alt="用户图片" class="message-image">
            </div>
            <div>{{ message.content }}</div>
          </div>
        </div>
        <div v-if="isSending" class="message assistant">
          <div class="message-avatar">AI</div>
          <div class="message-content loading">
            <div class="loading-dot"></div>
            <div class="loading-dot"></div>
            <div class="loading-dot"></div>
            <span>正在思考...</span>
          </div>
        </div>
      </div>

      <div class="chat-input">
        <div class="input-container">
          <div class="input-group">
            <textarea
              v-model="messageText"
              @keydown.enter.prevent="sendMessage"
              class="message-input"
              placeholder="输入消息..."
              rows="1"
            ></textarea>
            <label class="image-input">
              <input type="file" @change="handleImageUpload" accept="image/*">
              📷
            </label>
          </div>
          <button @click="sendMessage" :disabled="isSending || !messageText.trim()" class="send-button">
            ➤
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',
  data() {
    return {
      messages: [
        {
          role: 'assistant',
          content: '你好！我是智能对话助手。我可以理解文字和图片，有什么可以帮助你的吗？'
        }
      ],
      messageText: '',
      isSending: false,
      uploadedImage: null
    }
  },
  mounted() {
    this.autoResizeTextarea();
    this.scrollToBottom();
  },
  watch: {
    messages() {
      this.$nextTick(() => {
        this.scrollToBottom();
      });
    }
  },
  methods: {
    autoResizeTextarea() {
      const textarea = this.$el.querySelector('.message-input');
      textarea.addEventListener('input', () => {
        textarea.style.height = 'auto';
        textarea.style.height = Math.min(textarea.scrollHeight, 120) + 'px';
      });
    },

    scrollToBottom() {
      const container = this.$refs.messagesContainer;
      if (container) {
        container.scrollTop = container.scrollHeight;
      }
    },

    async handleImageUpload(event) {
      const file = event.target.files[0];
      if (!file) return;

      const formData = new FormData();
      formData.append('image', file);

      try {
        const response = await axios.post('/api/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        });

        if (response.data.success) {
          this.uploadedImage = response.data.data.base64;
          this.$emit('image-uploaded', this.uploadedImage);
        }
      } catch (error) {
        console.error('图片上传失败:', error);
        alert('图片上传失败，请重试');
      }
    },

    async sendMessage() {
      if (!this.messageText.trim() || this.isSending) return;

      const userMessage = {
        role: 'user',
        content: this.messageText.trim(),
        image: this.uploadedImage
      };

      this.messages.push(userMessage);
      this.isSending = true;
      this.messageText = '';
      this.uploadedImage = null;

      try {
        const chatMessages = this.messages.map(msg => ({
          role: msg.role,
          content: msg.content
        }));

        const requestBody = {
          messages: chatMessages
        };

        if (userMessage.image) {
          requestBody.image = userMessage.image;
        }

        const response = await axios.post('/api/chat', requestBody);

        if (response.data.code === 200) {
          const assistantMessage = {
            role: 'assistant',
            content: response.data.data.choices[0].message.content
          };
          this.messages.push(assistantMessage);
        } else {
          throw new Error(response.data.message);
        }
      } catch (error) {
        console.error('发送消息失败:', error);
        const errorMessage = {
          role: 'assistant',
          content: '抱歉，发送消息时出现了错误。请稍后重试。'
        };
        this.messages.push(errorMessage);
      } finally {
        this.isSending = false;
      }
    }
  }
}
</script>