<script setup lang="ts">
import { ref, computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

interface Channel {
  name: string;
  url: string;
}
const channels = ref<Channel[]>([
  {
    name: "知乎",
    url: "/c/zhihu",
  },
]);
const currentChannel = computed(() => {
  return channels.value.find((channel) => channel.url === route.path);
});
</script>

<template>
  <el-affix>
    <el-row
      justify="center"
      class="shadow-sm bg-white h-14 items-center text-lg border-b border-blue-50"
    >
      <el-col :span="20" class="rtl">
        <template v-for="channel in channels" :key="channel.name">
          <router-link
            :to="channel.url"
            :class="[
              'me-8',
              'py-4',
              'hover:border-b-2',
              currentChannel?.name == channel.name
                ? 'border-b-2 border-b-violet-500'
                : '',
            ]"
          >
            {{ channel.name }}
          </router-link>
        </template>
      </el-col>
    </el-row>
  </el-affix>
  <router-view v-slot="{ Component, route }">
    <transition mode="out-in">
      <component :is="Component" :key="route.path" />
    </transition>
  </router-view>
</template>
