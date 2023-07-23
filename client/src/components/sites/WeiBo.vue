<template>
  <div style="background-color: var(--el-bg-color);">
    <section v-for="(v, i) in listTopData" :key="i" class="HotItem">
      <div class="HotItem-index">
        <div class="HotItem-rank" :class=" i < 3 ? 'HotItem-hot' : '' ">{{ i + 1 }}</div>
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
        <img loading="lazy" :src="v.extra.image" alt="v.title">
      </a>
    </section>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {listData} from '@/api/data'

export default defineComponent({
  data() {
    return {
      listTopData: Array
    }
  },
  created() {
    listData({tab: "zhihu"})
      .then((res: any) => {
        for (let i = 0; i < res.list.length; i++) {
          res.list[i]['extra'] = JSON.parse(res.list[i]['extra'])
        }
        this.listTopData = res.list
      }).catch(err => {
      console.log(err)
    })
  },
})
</script>