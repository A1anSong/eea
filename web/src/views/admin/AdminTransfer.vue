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
    <el-table-column prop="Type" label="Type" width="120" />
    <el-table-column prop="Bank" label="Bank" width="120" />
    <el-table-column prop="BankAccount" label="BankAccount" width="120" />
    <el-table-column prop="Currency" label="Currency" width="120" />
    <el-table-column prop="Status" label="Status" width="120" />
    <el-table-column label="Amount" width="120">
      <template #default="scope">
            <label>{{ scope.row.Amount/10000 }}</label>
      </template>
    </el-table-column>
      <el-table-column label="Operations">
      <template #default="scope">
        <div v-if="scope.row.Status == 'init'">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
          >Confim</el-button
        >
        </div>
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
    <el-form-item label="Type">
      <label>{{currentRow.Type}}</label>
    </el-form-item>
    <el-form-item label="Bank">
      <label>{{currentRow.Bank}}</label>
    </el-form-item>
    <el-form-item label="BankAccount">
      <label>{{currentRow.BankAccount}}</label>
    </el-form-item>
    <el-form-item label="Currency">
        <label>{{currentRow.Currency}}</label>
    </el-form-item>
    <el-form-item label="Amount">
        <label>{{currentRow.Amount/10000}}</label>
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
  name: "AdminTransfer",
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
      delUserID: 0,
      statusOptions: [{label: "进行中", value: "init"},{label: "成功", value: "success"}, {label: "失败", value: "failed"}]
    }
  },
  mounted() {
    this.fetchData()
  },
  methods:{
    fetchData(){
      let t = this
      let params = {page: this.page, perPage: this.perPage}
      this.api.getTransferList(params).then(function(resp){
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
      this.api.transferConfim(this.currentRow.ID).then(function(resp){
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
