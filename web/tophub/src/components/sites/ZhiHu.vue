<template>
  <div style="background-color: var(--el-bg-color);">
    <section v-for="(v, i) in 10" :key="i" class="HotItem">
      <div class="HotItem-index">
        <div class="HotItem-rank" :class=" i < 3 ? 'HotItem-hot' : '' ">{{ i }}</div>
      </div>
      <div class="HotItem-content">
        <a href="" target="_blank">
          <h2 class="HotItem-title">

          </h2>
          <p class="HotItem-excerpt">

          </p>
        </a>
      </div>
      <a href=""></a>
    </section>
  </div>
</template>

<script lang="ts">
import type {TabsPaneContext} from 'element-plus'

export default {
  methods: {
    handleClick(tab: TabsPaneContext, event: Event) {
      console.log(tab, event)
    },
  }
}
</script>