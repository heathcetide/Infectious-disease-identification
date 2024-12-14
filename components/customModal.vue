<template>
  <view class="modal-container" v-if="visible">
    <view class="modal-dialog">
      <view class="modal-header">
        <text class="modal-title">{{ title }}</text>
      </view>
      <view class="modal-body">
        <text>{{ content }}</text>
      </view>
      <view class="modal-footer">
        <button @click="cancel" class="modal-cancel-button">{{ cancelText }}</button>
        <button @click="confirm" class="modal-confirm-button">{{ confirmText }}</button>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: '提示'
    },
    content: {
      type: String,
      default: '您确定要提交吗？'
    },
    confirmText: {
      type: String,
      default: '确定'
    },
    cancelText: {
      type: String,
      default: '取消'
    }
  },
  emits: ['confirm', 'cancel'],
  methods: {
    confirm() {
      this.$emit('confirm');
    },
    cancel() {
      this.$emit('cancel');
    }
  }
};
</script>

<style scoped>
.modal-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-dialog {
  background-color: #fff;
  padding: 20px;
  border-radius: 10px;
  width: 80%;
  max-width: 400px;
}

.modal-header {
  margin-bottom: 10px;
}

.modal-title {
  font-size: 18px;
  font-weight: bold;
}

.modal-body {
  margin-bottom: 10px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
}

.modal-cancel-button,
.modal-confirm-button {
  padding: 10px 20px;
  margin-left: 10px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.modal-cancel-button {
  background-color: #ccc;
}

.modal-confirm-button {
  background-color: #4CAF50;
  color: white;
}
</style>
