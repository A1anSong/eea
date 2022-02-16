<template>
  <el-header>Header
    <el-row justify="space-between">
      <el-col :span="6">contact</el-col>
      <el-col :span="6">social</el-col>
    </el-row>
    <el-affix>
      <el-row>
        <el-col :span="12">
          <div>
            <el-image src="/src/assets/logo.png" fit="contain" class="block"></el-image>
          </div>
        </el-col>
        <el-col :span="12">nav</el-col>
      </el-row>
    </el-affix>
  </el-header>
</template>

<script>
export default {
  name: "Header"
}
</script>

<style scoped>
.block {
  /*padding: 30px 0;*/
  /*text-align: center;*/
  /*border-right: solid 1px var(--el-border-color-base);*/
  /*display: inline-block;*/
  height: 100%;
  box-sizing: border-box;
  vertical-align: top;
}
</style>
