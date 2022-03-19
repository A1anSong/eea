<template>
  <!-- This example requires Tailwind CSS v2.0+ -->
<div class="px-4 sm:px-6 lg:px-8">
  <div class="sm:flex sm:items-center">
    <div class="sm:flex-auto">
      <h1 class="text-xl font-semibold text-gray-900">Users</h1>
      <p class="mt-2 text-sm text-gray-700">A list of all the users in your account including their name, title, email and role.</p>
    </div>
  </div>

  <el-table :data="datas" style="width: 100%" class="bg-gray-50">
    <el-table-column prop="ID" label="ID" width="40" />
    <el-table-column prop="User.email" label="Email" width="200" />
    <el-table-column prop="Currency" label="Currency" width="120" />
    <el-table-column label="Balance" width="120">
      <template #default="scope">
            <label>{{ scope.row.Balance/10000 }}</label>
      </template>
    </el-table-column>
      <el-table-column label="Operations">
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
          >Edit</el-button
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
      <label>{{currentRow.User.email}}</label>
    </el-form-item>
    <el-form-item label="Currency">
        <el-input v-model="currentRow.Currency"></el-input>
    </el-form-item>
    <el-form-item label="Balance">
        <el-input-number v-model="currentRow.Balance"></el-input-number>
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
</div>
</template>

<script>
import { ref } from 'vue'
import Api from '../../api/api'
import {ElMessage} from 'element-plus'

export default {
  name: "AdminBalance",
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
      this.api.getBalanceList(params).then(function(resp){
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
    updateRow(){
      let t = this
      this.api.updateBalance(this.currentRow.UserID, this.currentRow).then(function(resp){
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
