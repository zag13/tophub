<script setup lang="ts">
import { ref } from "vue";

const toBe = ref(false);
</script>

<template>
  <el-header class="bg-slate-100 w-full flex items-center">
    <el-row class="w-full ">
      <el-col class="flex p-1 justify-between">
        <div>
          <a href="https://essay.zag13.com/" target="_blank">
            <img class="h-8" src="/favicon.ico" alt="essay.zag13.com"
          /></a>
        </div>
        <button @click="() => (toBe = !toBe)">
          <span>{{ toBe ? "To Be" : "Not To Be" }}</span>
        </button>
      </el-col>
    </el-row>
  </el-header>
</template>
