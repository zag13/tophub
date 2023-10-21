<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { Feed, feed } from "@/api/channels.ts";

const route = useRoute();

const channels = ref<
  {
    name: string;
    url: string;
    site: string;
  }[]
>([
  {
    name: "知乎",
    url: "/c/zhihu",
    site: "zhihu",
  },
  {
    name: "微博",
    url: "/c/weibo",
    site: "weibo",
  },
  {
    name: "Hacker News",
    url: "/c/hacker-news",
    site: "hacker-news",
  },
]);
const currentChannel = computed(() => {
  return channels.value.find((channel) => channel.url === route.path);
});

const tops = ref<Feed[]>([]);

onMounted(async () => {
  feed({ site: currentChannel.value?.site })
    .then(({ data }) => {
      tops.value = data.list;
    })
    .catch((err) => {
      console.log(err);
    });
});
</script>

<template>
  <el-affix>
    <el-row
      justify="center"
      class="shadow-sm bg-white h-14 items-center text-lg border-b border-blue-50"
    >
      <el-col :xs="22" :sm="20" :lg="18" :xl="16" class="">
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
  <el-row justify="center" class="custom-channel-info bg-white my-10">
    <el-col :xs="22" :sm="20" :lg="18" :xl="16">
      <el-card class="py-3">
        <section
          v-for="(v, i) in tops"
          :key="i"
          class="flex py-3 h-17 px-5 text-gray-800"
          style="font-size: 1.05rem;font-weight: 400;"
        >
          <div class="w-full">
            <a
              :href="v.url"
              :title="v.title"
              target="_blank"
              class="visited:text-gray-300"
            >
              <h2 class="line-clamp-2">
                {{ v.rank }} &nbsp;&nbsp; {{ v.title }}
              </h2>
            </a>
          </div>
        </section>
      </el-card>
    </el-col>
  </el-row>
</template>

<style>
.custom-channel-info  .el-card__body {
    padding: 0;
}
</style>
```