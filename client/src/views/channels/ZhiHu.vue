<script setup lang="ts">
import { onMounted, ref } from "vue";
import { feed, Feed } from "@/api/channels";

interface Top {
  title: string;
  url: string;
  extra: {
    description: string;
    image: string;
  };
}

onMounted(async () => {
  feed({ site: "zhihu" })
    .then(({ data }) => {
      listTopData.value = data.list.map((item: Feed) => {
        return {
          title: item.title,
          url: item.url,
          extra: {
            description: "item.extra.description",
            image: "item.extra.image",
          },
        };
      });
    })
    .catch((err) => {
      console.log(err);
    });
});

const listTopData = ref<Top[]>([]);
</script>

<template>
  <el-row justify="center" class="bg-white mt-10">
    <el-col :span="20" class="rtl">
      <el-card>
        <section v-for="(v, i) in listTopData" :key="i" class="HotItem">
          <div class="HotItem-index">
            <div class="HotItem-rank" :class="i < 3 ? 'HotItem-hot' : ''">
              {{ i + 1 }}
            </div>
          </div>
          <div class="HotItem-content">
            <a :href="v.url" :title="v.title" target="_blank">
              <h2 class="HotItem-title">
                {{ v.title }}
              </h2>
              <p class="HotItem-excerpt">
                {{ v.extra.description }}
              </p>
            </a>
          </div>
          <a class="HotItem-img" :href="v.url" :title="v.title" target="_blank">
            <img loading="lazy" :src="v.extra.image" alt="v.title" />
          </a>
        </section>
      </el-card>
    </el-col>
  </el-row>
</template>

<style>
.el-tab-pane {
  height: 1000px;
}
</style>

<style scoped>
.HotItem {
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;
  padding: 1rem 1rem 1rem 0;
  position: relative;
  outline: none;
}

.HotItem:not(:first-child) {
  border: solid var(--el-border-color-light);
  border-width: 0.5px 0 0;
}

.HotItem-index {
  text-align: center;
  width: 3.5rem;
}

.HotItem-hot {
  color: #ff9607 !important;
}

.HotItem-rank {
  font-size: 1.125rem;
  line-height: 1.75rem;
  font-weight: 700;
  color: #999;
}

.HotItem-content {
  -webkit-box-flex: 1;
  -ms-flex: 1 1;
  flex: 1 1;
  overflow: hidden;
}

.HotItem-title {
  font-size: 1.125rem;
  line-height: 1.75rem;
  max-height: 3.5rem;
  display: -webkit-box;
  text-overflow: ellipsis;
  overflow: hidden;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  font-weight: 600;
  font-synthesis: style;
}

.HotItem-excerpt {
  line-height: 1.5rem;
  font-size: 1rem;
  margin-top: 0.2rem;
  min-height: 2rem;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.HotItem-img {
  display: block;
  height: 105px;
  margin-left: 16px;
  position: relative;
  z-index: 1;
}

.HotItem-img img {
  display: block;
  -o-object-fit: cover;
  object-fit: cover;
}

.HotItem-img:after,
.HotItem-img img {
  border-radius: 4px;
  height: 105px;
  width: 190px;
}
</style>
