<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.title"
            placeholder="输入文章名查找"
            enter-button
            allowClear
            @search="getArtList"
          />
        </a-col>
      </a-row>

      <!-- <a-table
        rowKey="ID"
        :columns="columns"
        :pagination="pagination"
        :dataSource="Artlist"
        v-if="Artlist.author == username"
        bordered
        @change="handleTableChange"
      > -->
      <a-table
        rowKey="ID"
        :columns="columns"
        :pagination="pagination"
        :dataSource="Artlist"
        bordered
        @change="handleTableChange"
      >
        <span class="ArtImg" slot="img" slot-scope="img">
          <img :src="img" />
        </span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              size="small"
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="$router.push(`/addart/${data.ID}`)"
              >编辑</a-button
            >
            <a-button
              size="small"
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deleteArt(data.ID)"
              >删除</a-button
            >
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
import day from 'dayjs'

const columns = [
  {
    title: '文章标题',
    dataIndex: 'title',
    width: '15%',
    key: 'title',
    align: 'center'
  },
  {
    title: '文章描述',
    dataIndex: 'desc',
    width: '20%',
    key: 'desc',
    align: 'center'
  },
  {
    title: '更新日期',
    dataIndex: 'UpdatedAt',
    width: '10%',
    key: 'UpdatedAt',
    align: 'center',
    customRender: val => {
      return val ? day(val).format('YYYY年MM月DD日 HH:mm') : '暂无'
    }
  },
  {
    title: '缩略图',
    dataIndex: 'img',
    width: '15%',
    key: 'img',
    align: 'center',
    scopedSlots: { customRender: 'img' }
  },
  {
    title: '操作',
    width: '13%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  created() {
    this.username = window.sessionStorage.getItem('username')
    this.getArtList()
    this.getArtListByAuthor()
  },
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: total => `共${total}条`
      },
      Artlist: [],
      Catelist: [],
      columns,
      queryParam: {
        title: '',
        pagesize: 5,
        pagenum: 1
      },
      username: ''
    }
  },
  methods: {
    // 根据用户名获取文章列表
    async getArtListByAuthor() {
      const { data: res } = await this.$http.get('articlefrontbyauthor', {
        params: {
          author: this.$route.params.author,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        //eslint-disable-next-line
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/')
        }
        this.$message.error(res.message)
      }

      this.Artlist = res.data
      this.pagination.total = res.total
    },
    // 根据标题搜索文章
    async getArtList() {
      const { data: res } = await this.$http.get('articlefront', {
        params: {
          title: this.queryParam.title,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        //eslint-disable-next-line
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/')
        }
        this.$message.error(res.message)
      }

      this.Artlist = res.data
      this.pagination.total = res.total

    },
    // 更改分页
    async handleTableChange(pagination) {
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
      this.getArtListByAuthor()
    },
    // 删除文章
    deleteArt(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该文章吗？一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`articlefront/${id}`)
          if (res.status != 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getArtListByAuthor()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        }
      })
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
.ArtImg {
  height: 100%;
  width: 100%;
}

.ArtImg img {
  width: 100px;
  height: 80px;
}
</style>
