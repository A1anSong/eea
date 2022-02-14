<template>
  <el-container>
    <el-header>header</el-header>
    <el-container>
      <el-aside>
        <el-menu>
          <el-sub-menu index="1">
            <template #title>
              <span>Navigator One</span>
            </template>
            <el-menu-item>
              aaa
            </el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-aside>
      <el-main>
        main
        <router-view/>
      </el-main>
    </el-container>
    <el-footer>footer</el-footer>
  </el-container>
</template>

<script>
export default {
  name: "Home"
}
</script>

<style scoped>
.el-container {
  height: 100%;
}
</style>
