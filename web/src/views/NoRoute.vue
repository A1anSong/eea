<template>
  <el-result title="404" sub-title="Sorry, request error" icon="error">
    <template #extra>
      <el-button type="primary" @click="$router.push('/')">Back</el-button>
    </template>
  </el-result>
</template>

<script>
export default {
  name: "NoRoute"
}
</script>

<style scoped>

</style>
