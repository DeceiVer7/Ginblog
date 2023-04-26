<template>
  <div>
    <p style="font-size: 18px; margin-left: 10px; margin-top: 7px" v-for="item in logs.slice().reverse()" :key="item.id">
     {{item.created_time +'，'+ item.content }}
    </p>

    <a-button type="danger" style="margin-left: 40%; margin-top: 15px" @click="deletelogs()">清空日志</a-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      logs: [],
    }
  },
  created() {
    this.getlogs()
  },
  methods: {
    async getlogs() {
      const { data: res } = await this.$http.get('getlogs')
      this.logs = res.data
    },
    deletelogs(){
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要清空所有日志吗？一旦清空，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`deletelogs`)
          if (res.status != 200) return this.$message.error(res.message)
          this.$message.success('清空成功')
          this.getlogs()
        },
        onCancel: () => {
          this.$message.info('取消清空')
        }
      })
    }
  },
}
</script>
