<template>
  <div class="min-h-full flex flex-col justify-center py-12 sm:px-6 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <img class="mx-auto h-12 w-auto" src="https://tailwindui.com/img/logos/workflow-mark-indigo-600.svg"
           alt="Workflow"/>
      <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">Sign in to your account</h2>
    </div>
    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <el-form class="space-y-6">
          <div>
            <el-form-item label="Email address" class="mt-0"/>
            <el-input v-model="email"/>
          </div>
          <div>
            <el-form-item label="Password"/>
            <el-input @keyup.enter="submit" class="mt-1" v-model="password" type="password" show-password/>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <el-checkbox v-model="remember" label="Remember me"/>
            </div>
            <div class="text-sm">
              <router-link to="/abcdefg" class="font-medium text-indigo-600 hover:text-indigo-500"> Forgot your
                password?
              </router-link>
            </div>
          </div>
          <div>
            <el-button @click="submit" :loading="loading">Sign in</el-button>
          </div>
        </el-form>
      </div>
    </div>
  </div>
</template>


<script>
import FormData from 'form-data'
import axios from 'axios'
import {ElMessage} from 'element-plus'
import router from "../router";

export default {
  name: "Login",
  data() {
    return {
      email: '',
      password: '',
      remember: true,
      loading: false,
    }
  },
  methods: {
    submit() {
      const form = new FormData()
      this.loading = true
      form.append('email', this.email)
      form.append('password', this.password)
      form.append('remember', this.remember)
      axios.post('/api/login', form)
          .then(function (response) {
            ElMessage.info(response.data.msg)
            window.location.reload()
          })
          .catch(error => {
            if (error.response.status === 400 || error.response.status === 401) {
              ElMessage.error(error.response.data.msg)
              this.loading = false
            }
          })
    },
  },
}
</script>

<style scoped>
:deep(.el-form-item) {
  @apply mb-0 !important;
}

:deep(.el-form-item__label) {
  @apply block text-sm font-medium text-gray-700 !important;
}

:deep(.el-input) {
  @apply mt-1 !important;
}

:deep(.el-input__inner) {
  @apply appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm !important;
}

:deep(.el-checkbox__original) {
  @apply h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded !important;
}

:deep(.el-checkbox__inner) {
  @apply border-indigo-600 hover:border-indigo-600 !important;
}

:deep(.el-checkbox__input.is-checked .el-checkbox__inner) {
  @apply bg-indigo-600 !important;
}

:deep(.el-checkbox__label) {
  @apply ml-2 block text-sm text-gray-900 !important;
}

:deep(.el-button) {
  @apply w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 !important;
}
</style>
