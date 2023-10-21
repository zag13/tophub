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
  <el-row justify="center" class="bg-white my-10">
    <el-col :xs="22" :sm="20" :lg="18" :xl="16">
      <el-card>
        <section
          v-for="(v, i) in tops"
          :key="i"
          class="flex py-3"
          style="height: 7.5rem"
        >
          <div>
            <p class="w-7">{{ v.rank }}</p>
          </div>
          <div class="w-full pr-7">
            <a
              :href="v.url"
              :title="v.title"
              target="_blank"
              class="visited:text-gray-300"
            >
              <h2 class="line-clamp-2" style="font-size: 1.125rem;font-weight: 700;">
                {{ v.title }}
              </h2>
              <p class="line-clamp-2 mt-2 break-al">
                {{ v.description }}
              </p>
            </a>
          </div>
        </section>
      </el-card>
    </el-col>
  </el-row>
</template>
