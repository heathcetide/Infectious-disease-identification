<template>
  <div
    class="chat-page"
    :class="{ dragging: isDragging }"
    @dragenter.prevent="onDragEnter"
    @dragover.prevent="onDragOver"
    @dragleave.prevent="onDragLeave"
    @drop.prevent="onDrop"
  >
    <header class="chat-head">
      <div class="head-left">
        <img class="head-logo" src="/logo.png" alt="" />
        <div>
          <h3>诊断专家</h3>
          <p class="sub">
            {{ info.model || '—' }} · {{ info.protocol || '—' }}
            <el-tag size="small" :type="info.enabled ? 'success' : 'danger'" style="margin-left: 6px">
              {{ info.enabled ? '已启用' : '未启用' }}
            </el-tag>
          </p>
        </div>
      </div>
      <div class="head-actions">
        <el-button text @click="showSettings = !showSettings">设置</el-button>
        <el-button text @click="clearChat">清空</el-button>
      </div>
    </header>

    <el-collapse-transition>
      <div v-show="showSettings" class="settings-panel">
        <div class="label">System Prompt</div>
        <el-input
          v-model="system"
          type="textarea"
          :rows="3"
          placeholder="可选系统提示；留空使用默认助手人设"
        />
      </div>
    </el-collapse-transition>

    <div ref="scroller" class="messages">
      <div v-if="!messages.length && !loading" class="empty-state">
        <img class="empty-icon" :src="agentIcon" alt="" />
        <p class="empty-title">诊断专家已就绪</p>
        <p class="empty-desc">输入症状描述，或把图片拖进对话区</p>
      </div>

      <div
        v-for="(m, i) in messages"
        :key="i"
        class="row"
        :class="m.role"
      >
        <img
          v-if="m.role === 'assistant'"
          class="avatar agent"
          :src="agentIcon"
          alt="agent"
        />
        <div class="bubble" :class="m.role">
          <div v-if="m.thinking" class="think-block">
            <button type="button" class="think-toggle" @click="m.thinkOpen = !m.thinkOpen">
              {{ m.thinkOpen ? '收起思考过程' : '查看思考过程' }}
            </button>
            <div v-show="m.thinkOpen" class="think-body">{{ m.thinking }}</div>
          </div>
          <div
            v-if="m.role === 'assistant'"
            class="md"
            v-html="renderMarkdown(m.displayContent || m.content)"
          />
          <div v-else class="content">{{ m.content }}</div>
          <div v-if="m.previews?.length" class="previews">
            <img v-for="(src, j) in m.previews" :key="j" :src="src" alt="" />
          </div>
          <div v-if="m.meta" class="meta">{{ m.meta }}</div>
        </div>
        <img
          v-if="m.role === 'user'"
          class="avatar me"
          src="/logo.png"
          alt="me"
        />
      </div>

      <div v-if="loading" class="row assistant">
        <img class="avatar agent" :src="agentIcon" alt="agent" />
        <div class="bubble assistant">
          <div class="typing">
            <span /><span /><span />
            <em>思考中</em>
          </div>
        </div>
      </div>
    </div>

    <div class="composer" @paste="onPaste">
      <div v-if="pendingPreviews.length" class="pending">
        <div v-for="(p, i) in pendingPreviews" :key="i" class="pending-item">
          <img :src="p.preview" alt="" />
          <button type="button" class="rm" @click="removePending(i)">×</button>
        </div>
      </div>
      <div class="composer-box">
        <el-input
          v-model="draft"
          type="textarea"
          :rows="3"
          resize="none"
          placeholder="输入问题；可粘贴或拖入图片。Enter 发送，Shift+Enter 换行"
          @keydown.enter.exact.prevent="send"
        />
        <el-button
          class="send-btn"
          type="primary"
          size="small"
          :loading="loading"
          :disabled="!info.enabled"
          @click="send"
        >
          发送
        </el-button>
      </div>
    </div>

    <div v-if="isDragging" class="drop-mask">松开以添加图片</div>
  </div>
</template>

<script>
import { marked } from 'marked'
import { getLLMInfo, postLLMChat } from '@/api/system/llm.js'
import agentIcon from '@/assets/menu_icons/diagnosis.svg'

marked.setOptions({ breaks: true, gfm: true })

const ALLOWED = new Set(['image/jpeg', 'image/jpg', 'image/png', 'image/webp', 'image/gif'])

function splitThink(raw) {
  let text = String(raw || '')
  const blocks = []
  const patterns = [
    /<(?:think|thinking|redacted_thinking)>([\s\S]*?)<\/(?:think|thinking|redacted_thinking)>/gi,
    /```(?:thinking|think)\s*([\s\S]*?)```/gi,
  ]
  for (const re of patterns) {
    text = text.replace(re, (_, body) => {
      if (body?.trim()) blocks.push(body.trim())
      return ''
    })
  }
  return {
    thinking: blocks.join('\n\n'),
    content: text.trim(),
  }
}

function readFileAsDataURL(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.onload = () => resolve(String(reader.result || ''))
    reader.onerror = reject
    reader.readAsDataURL(file)
  })
}

/** Convert SVG (unsupported by MiniMax) to PNG via canvas. */
async function svgToPngFile(file) {
  const text = await file.text()
  const blob = new Blob([text], { type: 'image/svg+xml' })
  const url = URL.createObjectURL(blob)
  try {
    const img = await new Promise((resolve, reject) => {
      const el = new Image()
      el.onload = () => resolve(el)
      el.onerror = reject
      el.src = url
    })
    const w = Math.max(img.naturalWidth || 512, 64)
    const h = Math.max(img.naturalHeight || 512, 64)
    const canvas = document.createElement('canvas')
    canvas.width = w
    canvas.height = h
    const ctx = canvas.getContext('2d')
    ctx.fillStyle = '#fff'
    ctx.fillRect(0, 0, w, h)
    ctx.drawImage(img, 0, 0, w, h)
    const dataUrl = canvas.toDataURL('image/png')
    const bin = atob(dataUrl.split(',')[1])
    const arr = new Uint8Array(bin.length)
    for (let i = 0; i < bin.length; i++) arr[i] = bin.charCodeAt(i)
    return new File([arr], (file.name || 'image').replace(/\.svg$/i, '.png'), { type: 'image/png' })
  } finally {
    URL.revokeObjectURL(url)
  }
}

async function normalizeImageFile(file) {
  const mime = (file.type || '').toLowerCase()
  if (mime === 'image/svg+xml' || /\.svg$/i.test(file.name || '')) {
    try {
      return await svgToPngFile(file)
    } catch (e) {
      throw new Error('SVG 无法转换，请改用 JPG/PNG')
    }
  }
  if (mime && !ALLOWED.has(mime) && !mime.startsWith('image/')) {
    throw new Error('仅支持图片')
  }
  if (mime && !ALLOWED.has(mime) && mime !== 'image/gif') {
    // try keep; backend may still accept
  }
  return file
}

async function fileToAttachment(file) {
  const normalized = await normalizeImageFile(file)
  const result = await readFileAsDataURL(normalized)
  const m = result.match(/^data:([^;]+);base64,(.+)$/)
  if (!m) throw new Error('read failed')
  return {
    mime: m[1],
    base64: m[2],
    fileName: normalized.name || file.name || 'image.png',
    preview: result,
  }
}

export default {
  name: 'LLMChatHarness',
  data() {
    return {
      agentIcon,
      info: { enabled: false, model: '', protocol: '' },
      system: '',
      showSettings: false,
      draft: '',
      messages: [],
      pendingPreviews: [],
      loading: false,
      isDragging: false,
      dragDepth: 0,
      typeTimer: null,
    }
  },
  async mounted() {
    try {
      const res = await getLLMInfo()
      if (res.data?.code === 1) this.info = res.data.data || {}
    } catch (e) {
      console.error(e)
    }
  },
  beforeUnmount() {
    if (this.typeTimer) clearInterval(this.typeTimer)
  },
  methods: {
    renderMarkdown(text) {
      try {
        return marked.parse(String(text || ''))
      } catch {
        return String(text || '')
      }
    },
    clearChat() {
      if (this.typeTimer) clearInterval(this.typeTimer)
      this.messages = []
      this.pendingPreviews = []
    },
    removePending(i) {
      this.pendingPreviews.splice(i, 1)
    },
    async addFiles(fileList) {
      const files = Array.from(fileList || []).filter((f) => f && f.type?.startsWith('image/'))
      for (const f of files) {
        try {
          const a = await fileToAttachment(f)
          this.pendingPreviews.push(a)
        } catch (e) {
          this.$message.warning(e.message || '图片处理失败')
        }
      }
    },
    onPaste(e) {
      const items = e.clipboardData?.items
      if (!items) return
      const files = []
      for (const it of items) {
        if (it.type?.startsWith('image/')) {
          const f = it.getAsFile()
          if (f) files.push(f)
        }
      }
      if (files.length) {
        e.preventDefault()
        this.addFiles(files)
      }
    },
    onDragEnter() {
      this.dragDepth += 1
      this.isDragging = true
    },
    onDragOver() {
      this.isDragging = true
    },
    onDragLeave() {
      this.dragDepth = Math.max(0, this.dragDepth - 1)
      if (this.dragDepth === 0) this.isDragging = false
    },
    onDrop(e) {
      this.dragDepth = 0
      this.isDragging = false
      const files = e.dataTransfer?.files
      if (files?.length) this.addFiles(files)
    },
    scrollBottom() {
      const el = this.$refs.scroller
      if (el) el.scrollTop = el.scrollHeight
    },
    typewriter(msgIndex, fullText) {
      return new Promise((resolve) => {
        if (this.typeTimer) clearInterval(this.typeTimer)
        let i = 0
        const step = Math.max(1, Math.ceil(fullText.length / 80))
        this.typeTimer = setInterval(() => {
          i = Math.min(fullText.length, i + step)
          const m = this.messages[msgIndex]
          if (m) m.displayContent = fullText.slice(0, i)
          this.scrollBottom()
          if (i >= fullText.length) {
            clearInterval(this.typeTimer)
            this.typeTimer = null
            resolve()
          }
        }, 18)
      })
    },
    async send() {
      if (this.loading) return
      const text = this.draft.trim()
      if (!text && !this.pendingPreviews.length) return

      const attachments = this.pendingPreviews.map((p) => ({
        mime: p.mime,
        base64: p.base64,
        fileName: p.fileName,
      }))
      const previews = this.pendingPreviews.map((p) => p.preview)

      this.messages.push({
        role: 'user',
        content: text || '（见图片）',
        previews,
      })
      this.draft = ''
      this.pendingPreviews = []
      this.loading = true
      this.$nextTick(() => this.scrollBottom())

      try {
        const history = this.messages
          .filter((m) => m.role === 'user' || m.role === 'assistant')
          .slice(0, -1)
          .slice(-8)
          .map((m) => ({ role: m.role, content: m.content }))
        const res = await postLLMChat({
          system: this.system,
          message: text,
          history,
          attachments,
        })
        if (res.data?.code !== 1) throw new Error(res.data?.msg || 'chat failed')
        const data = res.data.data || {}
        const parsed = splitThink(data.reply || '')
        this.loading = false
        this.messages.push({
          role: 'assistant',
          content: parsed.content,
          displayContent: '',
          thinking: parsed.thinking,
          thinkOpen: false,
          meta: [data.model, data.usedVision ? 'vision' : 'text'].filter(Boolean).join(' · '),
        })
        const idx = this.messages.length - 1
        await this.typewriter(idx, parsed.content)
      } catch (e) {
        this.loading = false
        this.messages.push({
          role: 'assistant',
          content: '错误：' + (e.message || e),
          displayContent: '错误：' + (e.message || e),
        })
      } finally {
        this.loading = false
        this.$nextTick(() => this.scrollBottom())
      }
    },
  },
}
</script>

<style scoped>
.chat-page {
  position: relative;
  height: calc(100vh - 140px);
  display: flex;
  flex-direction: column;
  background: var(--el-bg-color, #fff);
  border: 1px solid var(--el-border-color-lighter, #ebeef5);
  border-radius: 12px;
  overflow: hidden;
}
.chat-page.dragging {
  outline: 2px dashed #285ac8;
}
.chat-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid var(--el-border-color-lighter, #ebeef5);
  background: var(--admin-card-bg, #fafafa);
}
.head-left {
  display: flex;
  gap: 10px;
  align-items: center;
}
.head-logo {
  width: 36px;
  height: 36px;
  object-fit: contain;
  border-radius: 8px;
  background: #fff;
  padding: 2px;
}
.chat-head h3 {
  margin: 0;
  font-size: 16px;
  color: var(--el-text-color-primary);
}
.sub {
  margin: 2px 0 0;
  font-size: 12px;
  color: var(--el-text-color-secondary);
}
.head-actions {
  display: flex;
  gap: 4px;
}
.settings-panel {
  padding: 10px 16px;
  border-bottom: 1px solid var(--el-border-color-lighter, #eee);
  background: var(--el-fill-color-blank, #fff);
}
.settings-panel .label {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  margin-bottom: 6px;
}
.messages {
  flex: 1;
  min-height: 0;
  overflow: auto;
  padding: 16px 18px 8px;
  background: linear-gradient(180deg, #f5f7fb 0%, var(--el-bg-color, #fff) 120px);
}
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 42vh;
  padding: 32px 16px;
  text-align: center;
}
.empty-icon {
  width: 56px;
  height: 56px;
  padding: 12px;
  border-radius: 16px;
  background: #eef3ff;
  border: 1px solid #dbe4ff;
  object-fit: contain;
  margin-bottom: 14px;
  box-sizing: border-box;
}
.empty-title {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: var(--el-text-color-primary);
}
.empty-desc {
  margin: 8px 0 0;
  font-size: 13px;
  color: var(--el-text-color-secondary);
}
.row {
  display: flex;
  gap: 10px;
  margin-bottom: 16px;
  align-items: flex-start;
}
.row.user {
  justify-content: flex-end;
}
.avatar {
  width: 36px;
  height: 36px;
  border-radius: 10px;
  flex-shrink: 0;
  object-fit: contain;
  background: #fff;
  border: 1px solid var(--el-border-color-lighter, #e8e8e8);
  padding: 4px;
  box-sizing: border-box;
}
.avatar.agent {
  padding: 6px;
  background: #eef3ff;
}
.bubble {
  max-width: min(720px, 78%);
  padding: 10px 14px;
  border-radius: 14px;
  font-size: 14px;
  line-height: 1.6;
  color: var(--el-text-color-primary);
  box-shadow: 0 1px 2px rgba(15, 23, 42, 0.04);
}
.bubble.assistant {
  background: #fff;
  border: 1px solid var(--el-border-color-lighter, #ebeef5);
  border-top-left-radius: 4px;
}
.bubble.user {
  background: rgba(40, 90, 200, 0.12);
  border: 1px solid rgba(40, 90, 200, 0.18);
  border-top-right-radius: 4px;
}
.content {
  white-space: pre-wrap;
  word-break: break-word;
}
.md :deep(p) {
  margin: 0 0 0.6em;
}
.md :deep(p:last-child) {
  margin-bottom: 0;
}
.md :deep(ul),
.md :deep(ol) {
  margin: 0.4em 0;
  padding-left: 1.2em;
}
.md :deep(code) {
  background: rgba(15, 23, 42, 0.06);
  padding: 1px 5px;
  border-radius: 4px;
  font-size: 0.92em;
}
.md :deep(pre) {
  background: #0f172a;
  color: #e2e8f0;
  padding: 10px 12px;
  border-radius: 8px;
  overflow: auto;
  font-size: 12px;
}
.md :deep(pre code) {
  background: transparent;
  padding: 0;
  color: inherit;
}
.think-block {
  margin-bottom: 8px;
  border-radius: 8px;
  background: #f8fafc;
  border: 1px dashed #cbd5e1;
  overflow: hidden;
}
.think-toggle {
  width: 100%;
  border: 0;
  background: transparent;
  text-align: left;
  padding: 8px 10px;
  font-size: 12px;
  color: #64748b;
  cursor: pointer;
}
.think-body {
  padding: 0 10px 10px;
  font-size: 12px;
  color: #64748b;
  white-space: pre-wrap;
  line-height: 1.5;
  max-height: 220px;
  overflow: auto;
}
.previews {
  display: flex;
  gap: 8px;
  margin-top: 8px;
  flex-wrap: wrap;
}
.previews img {
  width: 112px;
  height: 112px;
  object-fit: cover;
  border-radius: 8px;
  border: 1px solid rgba(0, 0, 0, 0.06);
}
.meta {
  margin-top: 6px;
  font-size: 11px;
  color: var(--el-text-color-secondary);
}
.typing {
  display: flex;
  align-items: center;
  gap: 8px;
  min-height: 22px;
}
.typing span {
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: #285ac8;
  opacity: 0.35;
  animation: bounce 1.1s infinite ease-in-out;
}
.typing span:nth-child(2) {
  animation-delay: 0.15s;
}
.typing span:nth-child(3) {
  animation-delay: 0.3s;
}
.typing em {
  font-style: normal;
  font-size: 12px;
  color: #64748b;
}
@keyframes bounce {
  0%,
  80%,
  100% {
    transform: translateY(0);
    opacity: 0.3;
  }
  40% {
    transform: translateY(-5px);
    opacity: 1;
  }
}
.composer {
  border-top: 1px solid var(--el-border-color-lighter, #ebeef5);
  padding: 10px 14px 12px;
  background: var(--el-bg-color, #fff);
}
.composer-box {
  position: relative;
}
.composer-box :deep(.el-textarea__inner) {
  padding-bottom: 44px;
  min-height: 96px;
  border-radius: 12px;
}
.send-btn {
  position: absolute;
  right: 10px;
  bottom: 10px;
  z-index: 2;
}
.pending {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  margin-bottom: 8px;
}
.pending-item {
  position: relative;
  width: 72px;
  height: 72px;
}
.pending-item img {
  width: 72px;
  height: 72px;
  object-fit: cover;
  border-radius: 8px;
  border: 1px solid #e2e8f0;
}
.pending-item .rm {
  position: absolute;
  top: -6px;
  right: -6px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  border: 0;
  background: #0f172a;
  color: #fff;
  cursor: pointer;
  line-height: 20px;
  font-size: 14px;
}
.drop-mask {
  position: absolute;
  inset: 0;
  background: rgba(40, 90, 200, 0.12);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 600;
  color: #285ac8;
  pointer-events: none;
  z-index: 5;
}
</style>
