<template>
  <el-container>
    <el-header>
      <el-row class="h-full">
        <el-col :span="16" class="flex justify-between mx-auto items-center">
          <div>
            <a href="https://essay.zag13.com/" target="_blank"><strong class="title">TopHub</strong></a>
          </div>
          <UseDark v-slot="{ toggleDark }">
            <el-switch
              v-model="isBlack"
              @click="toggleDark()"
              :active-icon="Sunny"
              :inactive-icon="Moon"
              style="--el-switch-on-color: var(--el-bg-color-page);--el-switch-border-color: var(--el-color-info-light-7);"
              inline-prompt
            />
          </UseDark>
        </el-col>
      </el-row>
    </el-header>
    <el-main style="background-color: var(--el-bg-color-page)">
      <el-row class="justify-center">
        <el-col :span="16" class="" style="">
          <el-menu
            mode="horizontal"
            default-active="zhihu"
            @select="handleSelect"
            style="background-color: var(--el-bg-color); border: 1px var(--el-border-color) solid; border-radius: 0.375rem; margin-bottom: 10px;"
          >
            <el-menu-item v-for="site in sites" :index="site.tab">
              <img :src="site.icon" alt="" class="w-1 h-1" style="margin-right: 0.2rem">
              {{ site.name }}
            </el-menu-item>
          </el-menu>
          <zhi-hu/>
        </el-col>
      </el-row>
    </el-main>
    <el-footer class="text-center">Footer</el-footer>
  </el-container>
</template>

<script lang="ts">
import {UseDark} from "@vueuse/components";
import {Moon, Sunny} from '@element-plus/icons-vue'
import ZhiHu from "@/components/sites/ZhiHu.vue";

export default {
  components: {
    UseDark,
    ZhiHu
  },
  data() {
    return {
      sites: sites,
      isBlack: true,
      Sunny: Sunny,
      Moon: Moon,
    }
  },
  methods: {
    handleSelect(key: string, keyPath: string[]) {
      console.log(key, keyPath)
    }
  }
}

const sites = [
  {tab: "zhihu", name: "知乎", icon: "https://static.zhihu.com/heifetz/favicon.ico"},
  {tab: "weibo", name: "微博", icon: "https://weibo.com/favicon.ico"},
]
</script>
