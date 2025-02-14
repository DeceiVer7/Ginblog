<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="4">
          <a-button type="primary" @click="addCateVisible = true">新增标签</a-button>
        </a-col>
      </a-row>

      <a-table
        rowKey="id"
        :columns="columns"
        :pagination="pagination"
        :dataSource="Catelist"
        bordered
        @change="handleTableChange"
      >
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin-right: 15px" @click="editCate(data.id)">编辑</a-button>
            <a-button type="danger" icon="delete" style="margin-right: 15px" @click="deleteCate(data.id)"
              >删除</a-button
            >
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 新增标签区域 -->
    <a-modal
      closable
      title="新增标签"
      :visible="addCateVisible"
      width="60%"
      @ok="addCateOk"
      @cancel="addCateCancel"
      destroyOnClose
    >
      <a-form-model :model="newCate" :rules="addCateRules" ref="addCateRef">
        <a-form-model-item label="标签名称" prop="name">
          <a-input v-model="newCate.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑标签区域 -->
    <a-modal
      closable
      destroyOnClose
      title="编辑标签"
      :visible="editCateVisible"
      width="60%"
      @ok="editCateOk"
      @cancel="editCateCancel"
    >
      <a-form-model :model="CateInfo" :rules="CateRules" ref="addCateRef">
        <a-form-model-item label="标签名称" prop="name">
          <a-input v-model="CateInfo.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
import day from 'dayjs'
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '10%',
    key: 'id',
    align: 'center',
  },
  {
    title: '标签名',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center',
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' },
  },
]

export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      Catelist: [],
      CateInfo: {
        name: '',
        id: 0,
      },
      newCate: {
        name: '',
      },
      columns,
      queryParam: {
        pagesize: 5,
        pagenum: 1,
      },
      editVisible: false,
      CateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.CateInfo.name === '') {
                callback(new Error('请输入标签名'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
      addCateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.newCate.name === '') {
                callback(new Error('请输入标签名'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
      },
      editCateVisible: false,
      addCateVisible: false,
    }
  },
  created() {
    this.getCateList()
    this.username = window.sessionStorage.getItem('username')
  },
  methods: {
    // 删除标签日志
    async deleteCateLog(id) {
      const { data: res } = await this.$http.get(`category/${id}`)
          var nowDate = new Date();
          this.$http.post('log', {
          created_time: day(nowDate).format('YYYY年MM月DD日 HH:mm'),
          content: '管理员：' + this.username + '，删除了标签，' + '标签名称为：' + res.data.name,
        })
    },
    // 获取标签列表
    async getCateList() {
      const { data: res } = await this.$http.get('admin/category', {
        pagesize: this.queryParam.pagesize,
        pagenum: this.queryParam.pagenum,
      })

      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.Catelist = res.data
      this.pagination.total = res.total
    },

    // 更改分页
    handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getCateList()
    },
    // 删除标签
    deleteCate(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该标签吗？一旦删除，无法恢复',
        onOk: async () => {
          // 删除标签日志
          this.deleteCateLog(id)
          const { data: res } = await this.$http.delete(`category/${id}`)
          if (res.status != 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getCateList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },
    // 新增标签
    addCateOk() {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('category/add', {
          name: this.newCate.name,
        })
        // 新增标签日志
        var nowDate = new Date()
        await this.$http.post('log', {
          created_time: day(nowDate).format('YYYY年MM月DD日 HH:mm'),
          content: '管理员：' + this.username + '，新增了标签，' + '标签名称为：' + this.newCate.name,
        })
        if (res.status != 200) return this.$message.error(res.message)
        this.$refs.addCateRef.resetFields()
        this.addCateVisible = false
        this.$message.success('添加标签成功')
        this.getCateList()
      })
    },
    addCateCancel() {
      this.$refs.addCateRef.resetFields()
      this.addCateVisible = false
      this.$message.info('新增标签已取消')
    },
    // 编辑标签
    async editCate(id) {
      this.editCateVisible = true
      const { data: res } = await this.$http.get(`category/${id}`)
      this.CateInfo = res.data
      this.CateInfo.id = id
    },
    editCateOk() {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`category/${this.CateInfo.id}`, {
          name: this.CateInfo.name,
        })
        if (res.status != 200) return this.$message.error(res.message)
        this.editCateVisible = false
        this.$message.success('更新标签信息成功')
        this.getCateList()
        // 更新标签日志
        var nowDate = new Date()
        await this.$http.post('log', {
          created_time: day(nowDate).format('YYYY年MM月DD日 HH:mm'),
          content: '管理员：' + this.username + '，更新了标签，' + '标签名称为：' + this.CateInfo.name,
        })
      })
    },
    editCateCancel() {
      this.$refs.addCateRef.resetFields()
      this.editCateVisible = false
      this.$message.info('编辑已取消')
    },
  },
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
