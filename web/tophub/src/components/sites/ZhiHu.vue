<template>
  <div class="flex" style="background-color: var(--el-bg-color);">
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="User" name="first">User</el-tab-pane>
      <el-tab-pane label="Config" name="second">Config</el-tab-pane>
      <el-tab-pane label="Role" name="third">Role</el-tab-pane>
      <el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
    </el-tabs>
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