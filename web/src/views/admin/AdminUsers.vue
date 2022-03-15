<template>
  <!-- This example requires Tailwind CSS v2.0+ -->
<div class="px-4 sm:px-6 lg:px-8">
  <div class="sm:flex sm:items-center">
    <div class="sm:flex-auto">
      <h1 class="text-xl font-semibold text-gray-900">Users</h1>
      <p class="mt-2 text-sm text-gray-700">A list of all the users in your account including their name, title, email and role.</p>
    </div>
    <div class="mt-4 sm:mt-0 sm:ml-16 sm:flex-none">
      <button type="button" class="inline-flex items-center justify-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:w-auto">Add user</button>
    </div>
  </div>

  <el-table :data="datas" style="width: 100%" class="bg-gray-50">
    <el-table-column prop="ID" label="ID" width="40" />
    <el-table-column prop="email" label="email" width="200" />
    <el-table-column prop="firstName" label="FirstName" width="120" />
    <el-table-column prop="lastName" label="LastName" width="120" />
    <el-table-column prop="role" label="Role" width="120" />
    <el-table-column prop="status" label="Status" width="120" />
    <el-table-column prop="authLevel" label="AuthLevel" width="120" />
    <el-table-column fixed="lastLogin" label="LastLogin" width="120"/>
      <el-table-column label="Operations">
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
          >Edit</el-button
        >
        <el-button
          size="small"
          type="danger"
          class="delete"
          @click="handleDelete(scope.$index, scope.row)"
          >Delete</el-button
        >
      </template>
    </el-table-column>
  </el-table>

  <div class="pagination-block">
    <el-pagination layout="prev, pager, next" :total="total" v-model:currentPage="page"></el-pagination>
  </div>

<el-dialog v-model="openDialog" title="Warning" width="30%" center>
   <el-form :model="currentRow" label-width="120px">
    <el-form-item label="Email">
      <el-input v-model="currentRow.email"></el-input>
    </el-form-item>
    <el-form-item label="FirstName">
        <el-input v-model="currentRow.firstName"></el-input>
    </el-form-item>
    <el-form-item label="lastName">
        <el-input v-model="currentRow.lastName"></el-input>
    </el-form-item>
    <el-form-item label="Role">
      <el-select v-model="currentRow.role" placeholder="please select user role">
        <el-option label="Admin" value="admin"></el-option>
        <el-option label="User" value="user"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="Status">
      <el-select v-model="currentRow.status" placeholder="please select user status">
        <el-option label="Active" value="active"></el-option>
        <el-option label="InActive" value="inactive"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="AuthLevel">
      <el-select v-model="currentRow.authLevel" placeholder="please select user authLevel">
        <el-option
      v-for="item in authOptions"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    >
    </el-option>

      </el-select>
    </el-form-item>
   </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="openDialog = false">Cancel</el-button>
        <el-button type="primary" @click="updateRow" class="ebutton"
          >Confirm</el-button
        >
      </span>
    </template>
</el-dialog>
  <el-dialog v-model="deleteDialogOpen" title="Tips" width="30%" draggable>
    <span>Confim delete user {{currentRow.email}}</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="deleteDialogOpen = false">Cancel</el-button>
        <el-button type="primary" @click="confimDelete" class="delete"
          >Confirm</el-button
        >
      </span>
    </template>
  </el-dialog>
</div>
</template>

<script>
import { ref } from 'vue'
import Api from '../../api/api'
import {ElMessage} from 'element-plus'

export default {
  name: "AdminUsers",
  data() {
    return {
      api: new Api(),
      page: ref(1),
      total: ref(20),
      datas: [],
      openDialog: false,
      deleteDialogOpen: false,
      currentRow: {},
      authOptions: [{label: "None", value: 0},{label: "KYC", value: 1}],
      delUserID: 0
    }
  },
  mounted() {
    this.fetchData()
  },
  methods:{
    fetchData(){
      let t = this
      let params = {page: this.page, perPage: this.perPage}
      this.api.getUserList(params).then(function(resp){
        t.datas = resp.data.data
        t.total = resp.data.total
      }, error => {
        ElMessage.error(error.response.data.msg)
      })
    },
    handleEdit(index, row){
      this.currentRow = row
      this.openDialog = true
    },
    handleDelete(index, row){
      this.currentRow = row
      this.deleteDialogOpen = true
    },
    confimDelete(){
      let t = this
      let id = this.currentRow.ID
      this.api.deleteUser(id).then(function(resp){
        t.deleteDialogOpen = false
        t.fetchData()
      }, error => {
        ElMessage.error(error.response.data.msg)
      })
    },
    updateRow(){
      let t = this
      this.api.updateUser(this.currentRow).then(function(resp){
        t.openDialog = false
      }, error => {
        ElMessage.error(error.response.data.msg)
      })
    }
  }
}
</script>

<style scoped>
.pagination-block {
  margin-top: 10px;
  margin-bottom: 16px;
}
.delete{
  background-color: red;
}
.ebutton{
  background-color: blue;
}
</style>
