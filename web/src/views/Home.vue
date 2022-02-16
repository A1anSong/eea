<template>
  <el-container>
    <Header/>
    <el-main>
      Main
      <el-button type="primary">Home</el-button>
      <el-button type="primary" @click="$router.push('/login')">Login</el-button>
      <div v-for="i in 100">
        {{ i }}
      </div>
    </el-main>
    <el-footer>Footer</el-footer>
  </el-container>
</template>

<script>
import Header from "../components/Header.vue";

export default {
  name: "Home",
  components: {Header}
}
</script>

<style scoped>

</style>
