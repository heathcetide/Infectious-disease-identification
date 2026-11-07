// src/stores/tags.js
import { defineStore } from 'pinia';

export const useTagsStore = defineStore('tags', {
  state: () => ({
    tagsList: [],
  }),

  actions: {
    addTag(tag) {
      // 检查是否已经存在该标签，避免重复添加
      const exists = this.tagsList.find(t => t.path === tag.path);
      if (!exists) {
        this.tagsList.push(tag);
      }
    },

    removeTag(path) {
      this.tagsList = this.tagsList.filter(tag => tag.path !== path);
    },

    getPreviousTag(currentIndex) {
      return currentIndex > 0 ? this.tagsList[currentIndex - 1] : null;
    },

    removeAllTag(){
      this.tagsList = [];
    },
  },
});
