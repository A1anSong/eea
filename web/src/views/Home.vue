<template>
  <Header/>
  <Content/>
  <Footer/>
</template>

<script>
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import Content from "../components/Content.vue";

export default {
  name: "Home",
  components: {Content, Footer, Header},
}
</script>

<style scoped>

</style>
